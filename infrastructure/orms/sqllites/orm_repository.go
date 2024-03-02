package sqllites

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/orms"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
)

type ormRepository struct {
	hashAdapter hash.Adapter
	builders    map[string]interface{}
	skeleton    skeletons.Skeleton
	dbPtr       *sql.DB
}

func createOrmRepository(
	hashAdapter hash.Adapter,
	builders map[string]interface{},
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) orms.Repository {
	out := ormRepository{
		hashAdapter: hashAdapter,
		builders:    builders,
		skeleton:    skeleton,
		dbPtr:       dbPtr,
	}

	return &out
}

// Retrieve retrieves an instance by path and hash
func (app *ormRepository) Retrieve(path []string, hash hash.Hash) (orms.Instance, error) {
	allResources := app.skeleton.Resources()
	resource, err := allResources.FetchByPath(path)
	if err != nil {
		return nil, err
	}

	tableName := strings.Join(path, resourceNameDelimiter)
	return app.retrieveByResourceAndHash(tableName, resource, hash, allResources)
}

func (app *ormRepository) retrieveByResourceAndHash(
	table string,
	resource resources.Resource,
	hash hash.Hash,
	allResources resources.Resources,
) (orms.Instance, error) {
	values, err := app.retrieveFieldValuesByHash(
		table,
		resource.Key(),
		resource.Fields(),
		hash,
		allResources,
	)

	if err != nil {
		return nil, err
	}

	if builderIns, ok := app.builders[table]; ok {
		errorStr := ""
		initialize := resource.Initialize()
		retValue, err := app.callMethodWithParamsOnInstanceReturnOneValue(
			builderIns,
			initialize,
			[]interface{}{},
			&errorStr,
		)

		if errorStr != "" {
			str := fmt.Sprintf("there was an error while executing the initialize method (name: %s) on the builder (name: %s)", initialize, table)
			return nil, errors.New(str)
		}

		if err != nil {
			return nil, err
		}

		fieldsList := resource.Fields().List()
		for idx, oneField := range fieldsList {
			if !oneField.HasBuilder() {
				continue
			}

			pValue := values[idx].(*interface{})
			kind := oneField.Kind()
			builder := oneField.Builder()
			retValue, err = app.callMethodWithParamAndKindOnInstanceReturnOneValue(
				retValue,
				builder,
				*pValue,
				kind,
				&errorStr,
			)

			if errorStr != "" {
				str := fmt.Sprintf("there was an error while executing the field method (name: %s) on the builder (name: %s): %s", builder, table, errorStr)
				return nil, errors.New(str)
			}

			if err != nil {
				return nil, err
			}

			if retValue == nil {
				str := fmt.Sprintf("the field (name: %s) returned nil when calling its builder method (name: %s) on table: %s", oneField.Name(), oneField.Builder(), table)
				return nil, errors.New(str)
			}
		}

		trigger := resource.Trigger()
		retValue, err = app.callMethodWithParamsOnInstanceReturnOneValue(
			retValue,
			trigger,
			[]interface{}{},
			&errorStr,
		)

		if errorStr != "" {
			str := fmt.Sprintf("there was an error while executing the trigger method (name: %s) on the builder (name: %s): %s", trigger, table, errorStr)
			return nil, errors.New(str)
		}

		if err != nil {
			return nil, err
		}

		if casted, ok := retValue.(orms.Instance); ok {
			return casted, nil
		}

		return nil, errors.New("the built instance could not be casted to an orms.Instance value")
	}

	str := fmt.Sprintf("there is no builder for the provided table (name: %s)", table)
	return nil, errors.New(str)
}

func (app *ormRepository) retrieveFieldValuesByHash(
	table string,
	key resources.Field,
	fields resources.Fields,
	hash hash.Hash,
	allResources resources.Resources,
) ([]interface{}, error) {
	fieldNames := fields.Names()
	fieldNamesStr := strings.Join(fieldNames, ",")
	queryStr := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", fieldNamesStr, table, key.Name())
	rows, err := app.dbPtr.Query(queryStr, hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given key (name: %s, value: %s) do NOT match a %s instance", key.Name(), hash.String(), table)
		return nil, errors.New(str)
	}

	values := []interface{}{}
	fieldList := fields.List()
	for _, oneField := range fieldList {
		kind := oneField.Kind()
		if kind.IsConnection() {
			continue
		}

		retValue, err := app.generateValueFromKind(kind, allResources)
		if err != nil {
			return nil, err
		}

		values = append(values, &retValue)
	}

	err = rows.Scan(values...)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (app *ormRepository) generateValueFromKind(
	kind resources.Kind,
	resourcesIns resources.Resources,
) (interface{}, error) {
	if kind.IsNative() {
		native := kind.Native()
		if native.IsSingle() {
			pValue := native.Single()
			return app.generateValueFromNative(*pValue), nil
		}

		if native.IsList() {
			return app.generateValueFromNative(resources.NativeBytes), nil
		}
	}

	reference := kind.Reference()
	return app.generateValueFromReference(reference, resourcesIns)
}

func (app *ormRepository) generateValueFromReference(
	path []string,
	resources resources.Resources,
) (interface{}, error) {
	retResource, err := resources.FetchByPath(path)
	if err != nil {
		return nil, err
	}

	kind := retResource.Key().Kind()
	if !kind.IsNative() {
		return nil, errors.New("the key was expected to contain a native key")
	}

	native := kind.Native()
	if native.IsSingle() {
		pValue := native.Single()
		return app.generateValueFromNative(*pValue), nil
	}

	panic(errors.New("finish the list in repository: generateValueFromReference"))
}

func (app *ormRepository) generateValueFromNative(kind uint8) interface{} {
	if kind == resources.NativeInteger {
		var value int
		return value
	}

	if kind == resources.NativeFloat {
		var value float64
		return value
	}

	if kind == resources.NativeString {
		var value string
		return value
	}

	var value []byte
	return value
}

func (app *ormRepository) callMethodWithParamAndKindOnInstanceReturnOneValue(
	ins interface{},
	method string,
	param interface{},
	kind resources.Kind,
	pErrorStr *string,
) (interface{}, error) {
	if param == nil {
		return ins, nil
	}

	if kind.IsNative() {
		native := kind.Native()
		if native.IsList() {
			if casted, ok := param.([]byte); ok {
				list := native.List()
				delimiter := list.Delimiter()
				listValue := list.Value()
				if listValue == resources.NativeString {
					output := []string{}
					bytesList := bytes.Split(casted, []byte(delimiter))
					for _, oneBytes := range bytesList {
						output = append(output, string(oneBytes))
					}

					return app.callMethodWithParamsOnInstanceReturnOneValue(
						ins,
						method,
						[]interface{}{
							output,
						},
						pErrorStr,
					)
				}

				if listValue == resources.NativeBytes {
					panic(errors.New("retrieveByResourceAndHash -> finish bytes in ormRepository"))
				}

				if listValue == resources.NativeInteger {
					panic(errors.New("retrieveByResourceAndHash -> finish integer in ormRepository"))
				}

				if listValue == resources.NativeFloat {
					panic(errors.New("retrieveByResourceAndHash -> finish float in ormRepository"))
				}
			}

			// error
			return nil, errors.New("invalid casting")
		}
	}

	return app.callMethodWithParamsOnInstanceReturnOneValue(
		ins,
		method,
		[]interface{}{
			param,
		},
		pErrorStr,
	)
}

func (app *ormRepository) callMethodWithParamsOnInstanceReturnOneValue(
	ins interface{},
	method string,
	params []interface{},
	pErrorStr *string,
) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			value := fmt.Sprint(r)
			*pErrorStr = value
		}
	}()

	value := reflect.ValueOf(ins)
	methodName := value.MethodByName(method)
	if !methodName.IsValid() {
		str := fmt.Sprintf("there is no method (name: %s) on the provided instance", method)
		return nil, errors.New(str)
	}

	retType := methodName.Type()
	if retType == nil {
		return nil, errors.New("the type was expected to be mandatory")
	}

	methodAmountArguments := retType.NumIn()
	if methodAmountArguments != len(params) {
		str := fmt.Sprintf("the methodName (%s) was expected to contain %d arguments, but it contains %d arguments in reality", method, len(params), methodAmountArguments)
		return nil, errors.New(str)
	}

	methodParams := []reflect.Value{}
	if params != nil && len(params) > 0 {
		for _, oneParam := range params {

			expectedType := retType.In(0)
			value := reflect.ValueOf(oneParam)
			currentType := value.Type()

			// if the types are different, try to conver it:
			if expectedType.Kind() != currentType.Kind() {
				if value.CanConvert(expectedType) {
					value = value.Convert(expectedType)
				}
			}

			methodParams = append(methodParams, value)
		}
	}

	retValues := value.MethodByName(method).Call(methodParams)
	if len(retValues) < 1 {
		str := fmt.Sprintf("%d  values were returned, at least %d were expected, when calling the method (name %s)", len(retValues), 1, method)
		return nil, errors.New(str)
	}

	return retValues[0].Interface(), nil
}
