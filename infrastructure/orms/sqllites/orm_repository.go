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
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
)

type ormRepository struct {
	hashAdapter  hash.Adapter
	builders     map[string]interface{}
	listInstances  map[string]toListInstance
	skeleton     skeletons.Skeleton
	dbPtr        *sql.DB
}

func createOrmRepository(
	hashAdapter hash.Adapter,
	builders map[string]interface{},
	listInstances  map[string]toListInstance,
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) orms.Repository {
	out := ormRepository{
		hashAdapter:  hashAdapter,
		builders:     builders,
		listInstances: listInstances,
		skeleton:     skeleton,
		dbPtr:        dbPtr,
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

	allConnections := app.skeleton.Connections()
	tableName := app.createTableName(path)
	return app.retrieveByResourceAndHash(tableName, resource, hash, allResources, allConnections)
}

// List retrieves a list of to hashes
func (app *ormRepository) List(fromPath []string, toPath []string, fromHash hash.Hash) ([]hash.Hash, error) {
	allConnections := app.skeleton.Connections()
	connection, err := allConnections.FetchByPaths(fromPath, toPath)
	if err != nil {
		return nil, err
	}

	tableName := createConnectionTableName(
		fromPath,
		toPath,
	)

	queryStr := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", connection.To().Name(), tableName, connection.From().Name())
	rows, err := app.dbPtr.Query(queryStr, fromHash.Bytes())
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	toHashes := []hash.Hash{}
	for {
		if !rows.Next() {
			break
		}

		bytes := []byte{}
		err = rows.Scan(&bytes)
		if err != nil {
			return nil, err
		}

		pHash, err := app.hashAdapter.FromBytes(bytes)
		if err != nil {
			return nil, err
		}

		toHashes = append(toHashes, *pHash)
	}

	return toHashes, nil
}

func (app *ormRepository) createTableName(path []string) string {
	return strings.Join(path, resourceNameDelimiter)
}

func (app *ormRepository) retrieveByResourceAndHash(
	table string,
	resource resources.Resource,
	hash hash.Hash,
	allResources resources.Resources,
	allConnections connections.Connections,
) (orms.Instance, error) {
	values, err := app.retrieveFieldValuesByHash(
		table,
		resource.Key(),
		resource.Fields(),
		hash,
		allResources,
		allConnections,
	)

	if err != nil {
		return nil, err
	}

	if builderIns, ok := app.builders[table]; ok {
		return app.buildInstance(
			resource,
			builderIns,
			values,
		)
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
	allConnections connections.Connections,
) ([]interface{}, error) {
	fieldNames := app.fetchFieldsForSelect(fields)
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

	cpt := 0
	mapping := map[string]int{}
	values := []interface{}{}
	allFieldsList := fields.List()
	for _, oneField := range allFieldsList {
		kind := oneField.Kind()
		if kind.IsConnection() {
			continue
		}

		retValue, err := app.generateValueFromKind(kind, allResources)
		if err != nil {
			return nil, err
		}

		name := oneField.Name()
		values = append(values, &retValue)
		mapping[name] = cpt
		cpt++
	}

	err = rows.Scan(values...)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	output := []interface{}{}
	for _, oneField := range allFieldsList {
		kind := oneField.Kind()
		if kind.IsConnection() {
			name := kind.Connection()
			if allConnections == nil {
				str := fmt.Sprintf("the table (name: %s) contains a connection field (name: %s) but the provided skeleton has no connections", table, oneField.Name())
				return nil, errors.New(str)
			}

			connection, err := allConnections.Fetch(name)
			if err != nil {
				return nil, err
			}

			from := connection.From().Path()
			to := connection.To()
			toPath := to.Path()
			hashes, err := app.List(from, toPath, hash)
			if err != nil {
				return nil, err
			}

			list := []interface{}{}
			for _, oneHash := range hashes {
				ins, err := app.Retrieve(toPath, oneHash)
				if err != nil {
					return nil, err
				}

				list = append(list, ins)
			}

			if listInstanceFn, ok := app.listInstances[name]; ok {
				ins, err := listInstanceFn(list)
				if err != nil {
					return nil, err
				}

				output = append(output, ins)
				continue
			}

			str := fmt.Sprintf("the field (name: %s) requires a missing list builder (name: %s)", oneField.Name(), name)
			return nil, errors.New(str)
		}

		name := oneField.Name()
		idx := mapping[name]
		if kind.IsReference() {
			if pIns, ok := values[idx].(*interface{}); ok {
				insValue := *pIns
				if bytes, ok := insValue.([]byte); ok {
					pHash, err := app.hashAdapter.FromBytes(bytes)
					if err != nil {
						return nil, err
					}

					reference := kind.Reference()
					retInstance, err := app.Retrieve(reference, *pHash)
					if err != nil {
						return nil, err
					}

					output = append(output, &retInstance)
					continue
				}

				return nil, errors.New("the reference type was expected to contain bytes")
			}

			return nil, errors.New("the reference type was expected to contain bytes")
		}

		output = append(output, values[idx])
	}

	return output, nil
}

func (app *ormRepository) fetchFieldsForSelect(
	fields resources.Fields,
) []string {
	out := []string{}
	list := fields.List()
	for _, oneField := range list {
		if oneField.Kind().IsConnection() {
			continue
		}

		out = append(out, oneField.Name())
	}

	return out
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
						if len(oneBytes) <= 0 {
							continue
						}

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

func (app *ormRepository) buildInstance(
	resource resources.Resource,
	builderIns interface{},
	values []interface{},
) (orms.Instance, error) {
	errorStr := ""
	initialize := resource.Initialize()
	retValue, err := app.callMethodWithParamsOnInstanceReturnOneValue(
		builderIns,
		initialize,
		[]interface{}{},
		&errorStr,
	)

	if errorStr != "" {
		str := fmt.Sprintf("there was an error while executing the initialize method (name: %s)", initialize)
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

		value := values[idx]
		if pInstance, ok := value.(*orms.Instance); ok {
			value = *pInstance
		}

		if pValue, ok := value.(*interface{}); ok {
			value = *pValue
		}

		if pValue, ok := value.(*[]interface{}); ok {
			value = *pValue
		}

		kind := oneField.Kind()
		builder := oneField.Builder()
		builderMethod := builder.Method()
		if !builder.ContainsParam() {

			if boolValue, ok := value.(int64); ok {
				// if the value is false, skip the method call
				if boolValue == 0 {
					continue
				}
			}

			retValue, err = app.callMethodWithParamsOnInstanceReturnOneValue(
				retValue,
				builderMethod,
				[]interface{}{},
				&errorStr,
			)

			if errorStr != "" {
				str := fmt.Sprintf("there was an error while executing the field method (name: %s) on the builder (name: %s): %s", oneField.Name(), oneField.Builder(), errorStr)
				return nil, errors.New(str)
			}

			if err != nil {
				return nil, err
			}

			if retValue == nil {
				str := fmt.Sprintf("the field (name: %s) returned nil when calling its builder method (name: %s)", oneField.Name(), oneField.Builder())
				return nil, errors.New(str)
			}
		} else {
			retValue, err = app.callMethodWithParamAndKindOnInstanceReturnOneValue(
				retValue,
				builderMethod,
				value,
				kind,
				&errorStr,
			)

			if errorStr != "" {
				str := fmt.Sprintf("there was an error while executing the field method (name: %s) on the builder (name: %s): %s", oneField.Name(), oneField.Builder(), errorStr)
				return nil, errors.New(str)
			}

			if err != nil {
				return nil, err
			}

			if retValue == nil {
				str := fmt.Sprintf("the field (name: %s) returned nil when calling its builder method (name: %s)", oneField.Name(), oneField.Builder())
				return nil, errors.New(str)
			}
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
		str := fmt.Sprintf("there was an error while executing the trigger method (name: %s): %s", trigger, errorStr)
		return nil, errors.New(str)
	}

	if err != nil {
		return nil, err
	}

	if casted, ok := retValue.(orms.Instance); ok {
		return casted, nil
	}

	fmt.Printf("\n%v\n", retValue)
	fmt.Printf("\n%s\n", resource.Name())

	return nil, errors.New("the built instance could not be casted to an orms.Instance value")
}
