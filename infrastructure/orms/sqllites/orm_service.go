package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/orms"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
)

type table struct {
	name     string
	key      field
	fields   []field
	children []table
}

type field struct {
	name     string
	kind     kind
	canBeNil bool
}

type kind struct {
	pSingle     *uint8
	pList       *list
	pForeignKey *foreignKey
	pConnection *connection
}

type list struct {
	value     uint8
	delimiter string
}

type foreignKey struct {
	localField       field
	foreignTableName string
	foreignField     field
}

type connection struct {
	middleTableName string
	from            foreignKey
	to              foreignKey
}

type ormService struct {
	callMethodsOnInstances            map[string]callMethodOnInstanceFn
	listInstanceToElementHashesListFn map[string]listInstanceToElementHashesListFn
	repository                        orms.Repository
	hashAdapter                       hash.Adapter
	skeleton                          skeletons.Skeleton
	dbPtr                             *sql.DB
	txPtr                             *sql.Tx
}

func createOrmService(
	callMethodsOnInstances map[string]callMethodOnInstanceFn,
	listInstanceToElementHashesListFn map[string]listInstanceToElementHashesListFn,
	repository orms.Repository,
	hashAdapter hash.Adapter,
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
	txPtr *sql.Tx,
) orms.Service {
	out := ormService{
		callMethodsOnInstances:            callMethodsOnInstances,
		listInstanceToElementHashesListFn: listInstanceToElementHashesListFn,
		repository:                        repository,
		hashAdapter:                       hashAdapter,
		skeleton:                          skeleton,
		dbPtr:                             dbPtr,
		txPtr:                             txPtr,
	}

	return &out
}

// Init initializes the service
func (app *ormService) Init() error {
	resources := app.skeleton.Resources()
	connections := app.skeleton.Connections()
	tables, err := app.generateTables(resources, resources, connections)
	if err != nil {
		return err
	}

	schema, err := app.writeSchema(tables, "")
	if err != nil {
		return err
	}

	_, err = app.dbPtr.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts an instance
func (app *ormService) Insert(ins orms.Instance, path []string) error {
	allResources := app.skeleton.Resources()
	resource, err := allResources.FetchByPath(path)
	if err != nil {
		return err
	}

	connections := app.skeleton.Connections()
	tableName := strings.Join(path, resourceNameDelimiter)
	return app.insertResource(tableName, ins, resource, allResources, connections)
}

func (app *ormService) insertResource(
	tableName string,
	ins orms.Instance,
	resource resources.Resource,
	allResources resources.Resources,
	allConnections connections.Connections,
) error {
	key := resource.Key()
	keyName, keyValue, err := app.fetchFieldValue(tableName, ins, key, allResources, allConnections)
	if err != nil {
		return err
	}

	fields := resource.Fields()
	fieldValues, connectionFieldValues, err := app.fetchFieldsValueList(tableName, ins, fields, allResources, allConnections)
	if err != nil {
		return err
	}

	allFieldNames := []string{
		keyName,
	}

	allFieldValues := []interface{}{
		keyValue,
	}

	for name, value := range fieldValues {
		allFieldNames = append(allFieldNames, name)
		allFieldValues = append(allFieldValues, value)
	}

	fieldValuePlaceHolders := []string{}
	for range allFieldNames {
		fieldValuePlaceHolders = append(fieldValuePlaceHolders, "?")
	}

	fieldNamesStr := strings.Join(allFieldNames, ",")
	fieldValuePlaceHoldersStr := strings.Join(fieldValuePlaceHolders, ",")
	queryStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, fieldNamesStr, fieldValuePlaceHoldersStr)
	_, err = app.txPtr.Exec(queryStr, allFieldValues...)
	if err != nil {
		return err
	}

	err = app.insertConnectionValues(
		tableName,
		ins,
		fields,
		allResources,
		allConnections,
		connectionFieldValues,
	)

	if err != nil {
		return err
	}

	return nil
}

func (app *ormService) insertConnectionValues(
	table string,
	ins orms.Instance,
	fields resources.Fields,
	allResources resources.Resources,
	allConnections connections.Connections,
	fieldValues map[string]interface{},
) error {
	fromBytes := ins.Hash().Bytes()
	list := fields.List()
	for _, oneField := range list {
		fieldName := oneField.Name()
		if value, ok := fieldValues[fieldName]; ok {
			if casted, ok := value.(orms.Instance); ok {
				kind := oneField.Kind()
				if !kind.IsConnection() {
					continue
				}

				connectionName := kind.Connection()
				if fnToCall, ok := app.listInstanceToElementHashesListFn[connectionName]; ok {
					currentConnection, err := allConnections.Fetch(connectionName)
					if err != nil {
						return err
					}

					from := currentConnection.From()
					to := currentConnection.To()
					tableName := fmt.Sprintf(
						"%s%s%s",
						strings.Join(from.Path(), resourceNameDelimiter),
						connectionNameDelimiter,
						strings.Join(to.Path(), resourceNameDelimiter),
					)

					queryStr := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)", tableName, from.Name(), to.Name())
					elements, err := fnToCall(casted)
					if err != nil {
						return err
					}

					for _, oneElement := range elements {
						toBytes := oneElement.Bytes()
						_, err = app.txPtr.Exec(queryStr, fromBytes, toBytes)
						if err != nil {
							return err
						}
					}

					continue
				}

				str := fmt.Sprintf("field: %s: there is no list fetcher for the connections (name: %s)", oneField.Name(), connectionName)
				return errors.New(str)
			}

			if value == nil {
				continue
			}

			fmt.Printf("\n%v\n", value)

			str := fmt.Sprintf("field: %s: the field was expected to contain an Instance instance", oneField.Name())
			return errors.New(str)
		}
	}

	return nil
}

func (app *ormService) fetchFieldsValueList(
	table string,
	ins orms.Instance,
	fields resources.Fields,
	allResources resources.Resources,
	allConnections connections.Connections,
) (map[string]interface{}, map[string]interface{}, error) {
	output := map[string]interface{}{}
	connectionsOutput := map[string]interface{}{}
	list := fields.List()
	for _, oneField := range list {
		retName, retValue, err := app.fetchFieldValue(table, ins, oneField, allResources, allConnections)
		if err != nil {
			return nil, nil, err
		}

		if oneField.Kind().IsConnection() {
			connectionsOutput[retName] = retValue
			continue
		}

		output[retName] = retValue
	}

	return output, connectionsOutput, nil
}

func (app *ormService) fetchFieldValue(
	tableName string,
	ins orms.Instance,
	field resources.Field,
	allResources resources.Resources,
	allConnections connections.Connections,
) (string, interface{}, error) {
	fieldName := field.Name()
	if toCallFn, ok := app.callMethodsOnInstances[tableName]; ok {
		isExecuted, value, err := toCallFn(ins, fieldName)
		if err != nil {
			return "", nil, err
		}

		if !isExecuted {
			return fieldName, nil, nil
		}

		return fieldName, value, nil
	}

	str := fmt.Sprintf("there is not Instance fetcher associated with the table: %s", tableName)
	return "", nil, errors.New(str)

}

func (app *ormService) writeSchema(
	tables []table,
	parentName string,
) (string, error) {
	output := []string{}
	for _, oneTable := range tables {
		retTables, err := app.writeSchemaTable(oneTable, parentName)
		if err != nil {
			return "", err
		}

		output = append(output, retTables...)
	}

	connectionTablesList, err := app.writeSchemaConnectionTables(tables, parentName)
	if err != nil {
		return "", err
	}

	output = append(output, connectionTablesList...)
	return strings.Join(output, endOfLine), nil
}

func (app *ormService) writeSchemaTable(
	table table,
	parentName string,
) ([]string, error) {
	keyFieldStr, err := app.writeSchemaTableField(true, table.key)
	if err != nil {
		return nil, err
	}

	fieldsStrList, err := app.writeSchemaTableFieldsList(table.fields)
	if err != nil {
		return nil, err
	}

	foreignKeysList := app.writeSchemaTableFieldsForeignKeysList(table.fields)

	allFieldsList := []string{
		keyFieldStr,
	}

	allFieldsList = append(allFieldsList, fieldsStrList...)

	if len(foreignKeysList) > 0 {
		allFieldsList = append(allFieldsList, foreignKeysList...)
	}

	fieldsStr := strings.Join(allFieldsList, fmt.Sprintf("%s%s", ",", endOfLine))

	tableName := table.name
	if parentName != "" {
		tableName = fmt.Sprintf("%s%s%s", parentName, resourceNameDelimiter, tableName)
	}

	dropTableStr := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	createTableStr := fmt.Sprintf("CREATE TABLE %s (%s%s%s);", tableName, endOfLine, fieldsStr, endOfLine)
	childrenSchemaStr := ""
	if len(table.children) > 0 {
		retChildrenSchema, err := app.writeSchema(table.children, tableName)
		if err != nil {
			return nil, err
		}

		childrenSchemaStr = retChildrenSchema
	}

	tableSchema := strings.Join(
		[]string{
			dropTableStr,
			createTableStr,
			endOfLine,
		},
		endOfLine,
	)

	output := []string{
		tableSchema,
	}

	output = append(output, childrenSchemaStr)
	return output, nil
}

func (app *ormService) writeSchemaTableFieldsList(
	fields []field,
) ([]string, error) {
	fieldsList := []string{}
	for _, oneField := range fields {
		fieldStr, err := app.writeSchemaTableField(false, oneField)
		if err != nil {
			return nil, err
		}

		if fieldStr == "" {
			continue
		}

		fieldsList = append(fieldsList, fieldStr)
	}

	return fieldsList, nil
}

func (app *ormService) writeSchemaTableField(
	isPrimaryKey bool,
	field field,
) (string, error) {
	if field.kind.pConnection != nil {
		return "", nil
	}

	notNullStr := ""
	if !field.canBeNil {
		notNullStr = " NOT NULL"
	}

	primaryKeyStr := ""
	if isPrimaryKey {
		primaryKeyStr = " PRIMARY KEY"
	}

	kindStr, err := app.writeSchemaTableFieldKind(field.kind)
	if err != nil {
		return "", err
	}

	fieldStr := fmt.Sprintf(
		"%s %s %s%s",
		field.name,
		kindStr,
		primaryKeyStr,
		notNullStr,
	)

	return fieldStr, nil
}

func (app *ormService) writeSchemaTableFieldKind(
	kind kind,
) (string, error) {
	if kind.pForeignKey != nil {
		return app.writeSchemaTableFieldKind(kind.pForeignKey.localField.kind)
	}

	if kind.pList != nil {
		return app.writeSchemaTableFieldKindNative(resources.NativeBytes), nil
	}

	return app.writeSchemaTableFieldKindNative(*kind.pSingle), nil
}

func (app *ormService) writeSchemaConnectionTables(
	tables []table,
	parentName string,
) ([]string, error) {
	output := []string{}
	for _, oneTable := range tables {
		retTables, err := app.writeSchemaConnectionTable(oneTable, parentName)
		if err != nil {
			return nil, err
		}

		output = append(output, retTables...)
	}

	return output, nil
}

func (app *ormService) writeSchemaConnectionTable(
	table table,
	parentName string,
) ([]string, error) {
	output := []string{}
	for _, oneField := range table.fields {
		retFieldTable, err := app.writeSchemaConnectionField(oneField, parentName)
		if err != nil {
			return nil, err
		}

		if retFieldTable == "" {
			continue
		}

		output = append(output, retFieldTable)
	}

	return output, nil
}

func (app *ormService) writeSchemaConnectionField(
	field field,
	parentName string,
) (string, error) {
	if field.kind.pConnection == nil {
		return "", nil
	}

	pConnection := field.kind.pConnection
	fromForeignKey := pConnection.from
	fromForeignKeyStr := app.writeSchemaForeignKey(fromForeignKey)
	fromFieldStr, err := app.writeSchemaTableField(false, fromForeignKey.localField)
	if err != nil {
		return "", nil
	}

	toForeignKey := pConnection.to
	toForeignKeyStr := app.writeSchemaForeignKey(toForeignKey)
	toFieldStr, err := app.writeSchemaTableField(false, toForeignKey.localField)
	if err != nil {
		return "", nil
	}

	fieldsStr := strings.Join(
		[]string{
			fromFieldStr,
			toFieldStr,
			fromForeignKeyStr,
			toForeignKeyStr,
		},
		fmt.Sprintf("%s%s", ",", endOfLine),
	)

	tableName := pConnection.middleTableName
	dropTableStr := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	createTableStr := fmt.Sprintf("CREATE TABLE %s (%s%s%s);", tableName, endOfLine, fieldsStr, endOfLine)
	tableSchema := strings.Join(
		[]string{
			dropTableStr,
			createTableStr,
			endOfLine,
		},
		endOfLine,
	)

	return tableSchema, nil
}

func (app *ormService) writeSchemaTableFieldsForeignKeysList(
	fields []field,
) []string {
	output := []string{}
	for _, oneField := range fields {
		foreignKey := app.writeSchemaTableFieldForeignKeysList(oneField)
		if foreignKey == "" {
			continue
		}

		output = append(output, foreignKey)
	}

	return output
}

func (app *ormService) writeSchemaTableFieldForeignKeysList(
	field field,
) string {
	kind := field.kind
	if kind.pForeignKey == nil {
		return ""
	}

	return app.writeSchemaForeignKey(*kind.pForeignKey)
}

func (app *ormService) writeSchemaForeignKey(
	foreignKey foreignKey,
) string {
	return fmt.Sprintf(
		"FOREIGN KEY(%s) REFERENCES %s(%s)",
		foreignKey.localField.name,
		foreignKey.foreignTableName,
		foreignKey.foreignField.name,
	)
}

func (app *ormService) writeSchemaTableFieldKindNative(
	native uint8,
) string {
	if resources.NativeInteger == native {
		return "INTEGER"
	}

	if resources.NativeFloat == native {
		return "REAL"
	}

	if resources.NativeString == native {
		return "TEXT"
	}

	return "BLOB"
}

func (app *ormService) generateTables(
	resources resources.Resources,
	allResources resources.Resources,
	allConnections connections.Connections,
) ([]table, error) {
	output := []table{}
	list := resources.List()
	for _, oneResource := range list {
		pTable, err := app.generateTable(oneResource, allResources, allConnections)
		if err != nil {
			return nil, err
		}

		output = append(output, *pTable)
	}

	return output, nil
}

func (app *ormService) generateTable(
	resource resources.Resource,
	resources resources.Resources,
	connections connections.Connections,
) (*table, error) {
	name := resource.Name()
	key := resource.Key()
	keyName := key.Name()
	pCreatedKey, err := app.generateField(keyName, key, resources, connections)
	if err != nil {
		str := fmt.Sprintf("there was a problem while generating the key field for table (resource: %s): %s", name, err.Error())
		return nil, errors.New(str)
	}

	fields := resource.Fields()
	createdKeys, err := app.generateFields(fields, resources, connections)
	if err != nil {
		str := fmt.Sprintf("there was a problem while generating a field for table (resource: %s): %s", name, err.Error())
		return nil, errors.New(str)
	}

	output := table{
		name:   name,
		key:    *pCreatedKey,
		fields: createdKeys,
	}

	if resource.HasChildren() {
		children := resource.Children()
		tables, err := app.generateTables(
			children,
			resources,
			connections,
		)

		if err != nil {
			return nil, err
		}

		output.children = tables
	}

	return &output, nil
}

func (app *ormService) generateFields(
	fields resources.Fields,
	resources resources.Resources,
	connections connections.Connections,
) ([]field, error) {
	output := []field{}
	list := fields.List()
	for _, oneField := range list {
		fieldName := oneField.Name()
		pCreatedField, err := app.generateField(fieldName, oneField, resources, connections)
		if err != nil {
			return nil, err
		}

		output = append(output, *pCreatedField)
	}

	return output, nil
}

func (app *ormService) generateField(
	name string,
	fieldIns resources.Field,
	resources resources.Resources,
	connections connections.Connections,
) (*field, error) {
	kind := fieldIns.Kind()
	createdKind, err := app.generateFieldKind(name, kind, resources, connections)
	if err != nil {
		return nil, err
	}

	return &field{
		name:     name,
		kind:     *createdKind,
		canBeNil: fieldIns.HasCondition(),
	}, nil
}

func (app *ormService) generateFieldKind(
	fieldName string,
	kindIns resources.Kind,
	resources resources.Resources,
	connections connections.Connections,
) (*kind, error) {
	output := kind{}
	if kindIns.IsConnection() {
		if connections == nil {
			str := fmt.Sprintf("the field (name: %s) contains a connection, but the skeleton does not contain any", fieldName)
			return nil, errors.New(str)
		}

		connectionName := kindIns.Connection()
		retConnection, err := connections.Fetch(connectionName)
		if err != nil {
			return nil, err
		}

		pCreatedConn, err := app.createConnection(resources, retConnection, connections)
		if err != nil {
			return nil, err
		}

		output.pConnection = pCreatedConn
	}

	if kindIns.IsReference() {
		reference := kindIns.Reference()
		pForeignKey, err := app.createForeignKey(resources, reference, fieldName, connections)
		if err != nil {
			return nil, err
		}

		output.pForeignKey = pForeignKey
	}

	if kindIns.IsNative() {
		native := kindIns.Native()
		if native.IsSingle() {
			pValue := native.Single()
			output.pSingle = pValue
		}

		if native.IsList() {
			nativeList := native.List()
			output.pList = &list{
				value:     nativeList.Value(),
				delimiter: nativeList.Delimiter(),
			}
		}
	}

	return &output, nil
}

func (app *ormService) createConnection(
	resources resources.Resources,
	connectionIns connections.Connection,
	connections connections.Connections,
) (*connection, error) {
	from := connectionIns.From()
	fromName := from.Name()
	fromPath := from.Path()

	toField := connectionIns.To()
	toName := toField.Name()
	toPath := toField.Path()
	tableName := createConnectionTableName(
		fromPath,
		toPath,
	)

	pFrom, err := app.createForeignKey(resources, fromPath, fromName, connections)
	if err != nil {
		return nil, err
	}

	pTo, err := app.createForeignKey(resources, toPath, toName, connections)
	if err != nil {
		return nil, err
	}

	return &connection{
		middleTableName: tableName,
		from:            *pFrom,
		to:              *pTo,
	}, nil
}

func (app *ormService) createForeignKey(
	resources resources.Resources,
	path []string,
	localFieldName string,
	connections connections.Connections,
) (*foreignKey, error) {
	foreignTableName := strings.Join(path, resourceNameDelimiter)
	retResource, err := resources.FetchByPath(path)
	if err != nil {
		return nil, err
	}

	key := retResource.Key()
	pLocalField, err := app.generateField(localFieldName, key, resources, connections)
	if err != nil {
		return nil, err
	}

	foreignFieldName := key.Name()
	pForeignField, err := app.generateField(foreignFieldName, key, resources, connections)
	if err != nil {
		return nil, err
	}

	return &foreignKey{
		localField:       *pLocalField,
		foreignTableName: foreignTableName,
		foreignField:     *pForeignField,
	}, nil
}

// Delete deletes an instance
func (app *ormService) Delete(path []string, hash hash.Hash) error {
	return nil
}
