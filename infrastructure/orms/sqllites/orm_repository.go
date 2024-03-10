package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons/resources"
	"github.com/steve-care-software/datastencil/domain/hash"
)

type ormRepository struct {
	hashAdapter              hash.Adapter
	buildInstances           map[string]buildInstanceFn
	elementsToListInstanceFn map[string]elementsToListInstanceFn
	skeleton                 skeletons.Skeleton
	dbPtr                    *sql.DB
}

func createOrmRepository(
	hashAdapter hash.Adapter,
	buildInstances map[string]buildInstanceFn,
	elementsToListInstanceFn map[string]elementsToListInstanceFn,
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) instances.RepositoryBackup {
	out := ormRepository{
		hashAdapter:              hashAdapter,
		buildInstances:           buildInstances,
		elementsToListInstanceFn: elementsToListInstanceFn,
		skeleton:                 skeleton,
		dbPtr:                    dbPtr,
	}

	return &out
}

// Retrieve retrieves an instance by path and hash
func (app *ormRepository) Retrieve(path []string, hash hash.Hash) (instances.Instance, error) {
	allResources := app.skeleton.Resources()
	resource, err := allResources.FetchByPath(path)
	if err != nil {
		return nil, err
	}

	allConnections := app.skeleton.Connections()
	tableName := app.createTableName(path)
	return app.retrieveByResourceAndHash(tableName, resource, hash, allResources, allConnections)
}

// List returns the list of hashes connected to the list
func (app *ormRepository) List(path []string) ([]hash.Hash, error) {
	return nil, nil
}

func (app *ormRepository) connectedList(fromPath []string, toPath []string, fromHash hash.Hash) ([]hash.Hash, error) {
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
) (instances.Instance, error) {
	valuesMap, err := app.retrieveFieldValuesByHash(
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

	if fnToBuild, ok := app.buildInstances[table]; ok {
		return fnToBuild(valuesMap)
	}

	str := fmt.Sprintf("there is no builderInstances fn for the provided table (name: %s)", table)
	return nil, errors.New(str)
}

func (app *ormRepository) retrieveFieldValuesByHash(
	table string,
	key resources.Field,
	fields resources.Fields,
	hash hash.Hash,
	allResources resources.Resources,
	allConnections connections.Connections,
) (map[string]interface{}, error) {
	values := []interface{}{}
	mapping := map[string]int{}
	allFieldsList := fields.List()
	fieldNames := app.fetchFieldsForSelect(fields)
	if len(fieldNames) > 0 {
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
	}

	valuesMap := map[string]interface{}{}
	for _, oneField := range allFieldsList {
		fieldName := oneField.Name()
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
			hashes, err := app.connectedList(from, toPath, hash)
			if err != nil {
				return nil, err
			}

			if len(hashes) <= 0 {
				continue
			}

			list := []interface{}{}
			for _, oneHash := range hashes {
				ins, err := app.Retrieve(toPath, oneHash)
				if err != nil {
					return nil, err
				}

				list = append(list, ins)
			}

			if listInstanceFn, ok := app.elementsToListInstanceFn[name]; ok {
				ins, err := listInstanceFn(list)
				if err != nil {
					return nil, err
				}

				valuesMap[fieldName] = ins
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
				if insValue == nil {
					continue
				}

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

					valuesMap[name] = &retInstance
					continue
				}

				return nil, errors.New("the reference type is invalid")
			}

			return nil, errors.New("the reference type was expected to contain bytes")
		}

		value := values[idx]
		if pValue, ok := values[idx].(*interface{}); ok {
			value = *pValue
		}

		valuesMap[name] = value
	}

	return valuesMap, nil
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
