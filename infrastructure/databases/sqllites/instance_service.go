package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
	commit_resources "github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/resources"
)

type instanceService struct {
	callMethodsOnInstances            map[string]callMethodOnInstanceFn
	listInstanceToElementHashesListFn map[string]listInstanceToElementHashesListFn
	repository                        instances.Repository
	hashAdapter                       hash.Adapter
	pointerBuilder                    pointers.Builder
	resourceBuilder                   commit_resources.Builder
	actionsBuilder                    actions.Builder
	actionBuilder                     actions.ActionBuilder
	contentBuilder                    commits.ContentBuilder
	commitBuilder                     commits.Builder
	skeleton                          skeletons.Skeleton
	signer                            signers.Signer
	pDB                               *sql.DB
	contextsMap                       map[uint]*transaction
	contextCpt                        int
}

func createInstanceService(
	callMethodsOnInstances map[string]callMethodOnInstanceFn,
	listInstanceToElementHashesListFn map[string]listInstanceToElementHashesListFn,
	repository instances.Repository,
	hashAdapter hash.Adapter,
	pointerBuilder pointers.Builder,
	resourceBuilder commit_resources.Builder,
	actionsBuilder actions.Builder,
	actionBuilder actions.ActionBuilder,
	contentBuilder commits.ContentBuilder,
	commitBuilder commits.Builder,
	skeleton skeletons.Skeleton,
	signer signers.Signer,
	pDB *sql.DB,
) instances.Service {
	out := instanceService{
		callMethodsOnInstances:            callMethodsOnInstances,
		listInstanceToElementHashesListFn: listInstanceToElementHashesListFn,
		repository:                        repository,
		hashAdapter:                       hashAdapter,
		pointerBuilder:                    pointerBuilder,
		resourceBuilder:                   resourceBuilder,
		actionsBuilder:                    actionsBuilder,
		actionBuilder:                     actionBuilder,
		contentBuilder:                    contentBuilder,
		commitBuilder:                     commitBuilder,
		skeleton:                          skeleton,
		signer:                            signer,
		pDB:                               pDB,
		contextsMap:                       map[uint]*transaction{},
		contextCpt:                        0,
	}

	return &out
}

// Init initializes the service
func (app *instanceService) Init() error {
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

	_, err = app.pDB.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

// Begin begins a transaction
func (app *instanceService) Begin() (*uint, error) {
	newContext := app.contextCpt + 1
	if newContext%100 == 0 {
		maxIndex := uint(0)
		for idx := range app.contextsMap {
			if idx < maxIndex {
				continue
			}

			maxIndex = idx
		}

		newContext = int(maxIndex + 1)
	}

	pTx, err := app.pDB.Begin()
	if err != nil {
		return nil, err
	}

	casted := uint(newContext)
	app.contextsMap[casted] = &transaction{
		actionsList: []actions.Action{},
		pTX:         pTx,
	}

	return &casted, nil
}

// Insert inserts an instance
func (app *instanceService) Insert(context uint, ins instances.Instance, path []string) error {
	if app.skeleton.Blacklist().Contains(path) {
		pathStr := strings.Join(path, "/")
		str := fmt.Sprintf("instances of the path (%s) cannot be inserted in the database because it is blacklisted", pathStr)
		return errors.New(str)
	}

	if pTrx, ok := app.contextsMap[context]; ok {
		allResources := app.skeleton.Resources()
		resource, err := allResources.FetchByPath(path)
		if err != nil {
			return err
		}

		resourceToInsert, err := app.resourceBuilder.Create().
			WithPath(path).
			WithInstance(ins).
			Now()

		if err != nil {
			return err
		}

		connections := app.skeleton.Connections()
		tableName := strings.Join(path, resourceNameDelimiter)
		err = app.insertResource(
			tableName,
			resourceToInsert,
			resource,
			allResources,
			connections,
			pTrx.pTX,
		)

		if err != nil {
			return err
		}

		action, err := app.actionBuilder.Create().
			WithInsert(resourceToInsert).
			Now()

		if err != nil {
			return err
		}

		app.contextsMap[context].actionsList = append(app.contextsMap[context].actionsList, action)
		return nil
	}

	str := fmt.Sprintf(contextDoesNotExistsErrorStr, context)
	return errors.New(str)
}

// Delete deletes an instance
func (app *instanceService) Delete(context uint, path []string, hash hash.Hash) error {
	if pTrx, ok := app.contextsMap[context]; ok {
		allResources := app.skeleton.Resources()
		resource, err := allResources.FetchByPath(path)
		if err != nil {
			return err
		}

		keyName := resource.Key().Name()
		tableName := strings.Join(path, resourceNameDelimiter)
		queryStr := fmt.Sprintf("DROP TABLE %s where %s = ?", tableName, keyName)
		_, err = pTrx.pTX.Exec(queryStr, hash.Bytes())
		if err != nil {
			return err
		}

		return nil
	}

	str := fmt.Sprintf(contextDoesNotExistsErrorStr, context)
	return errors.New(str)
}

// Commit commits actions
func (app *instanceService) Commit(context uint) error {
	if pTrx, ok := app.contextsMap[context]; ok {
		content, err := app.buildCommitContent(pTrx.actionsList)
		if err != nil {
			return err
		}

		msg := content.Hash().String()
		signature, err := app.signer.Sign(msg)
		if err != nil {
			return err
		}

		commit, err := app.commitBuilder.Create().
			WithContent(content).
			WithSignature(signature).
			Now()

		if err != nil {
			return err
		}

		commitPath := app.skeleton.Commit()
		allResources := app.skeleton.Resources()
		resource, err := allResources.FetchByPath(commitPath)
		if err != nil {
			return err
		}

		resourceToInsert, err := app.resourceBuilder.Create().
			WithPath(commitPath).
			WithInstance(commit).
			Now()

		if err != nil {
			return err
		}

		connections := app.skeleton.Connections()
		tableName := strings.Join(commitPath, resourceNameDelimiter)
		err = app.insertResource(
			tableName,
			resourceToInsert,
			resource,
			allResources,
			connections,
			pTrx.pTX,
		)

		if err != nil {
			return err
		}

		err = pTrx.pTX.Commit()
		if err != nil {
			return err
		}

		app.contextsMap[context].pTX = nil
		delete(app.contextsMap, context)
		return nil
	}

	str := fmt.Sprintf(contextDoesNotExistsErrorStr, context)
	return errors.New(str)
}

// Cancel cancels a context
func (app *instanceService) Cancel(context uint) error {
	if _, ok := app.contextsMap[context]; ok {
		app.contextsMap[context].pTX = nil
		delete(app.contextsMap, context)
		return nil
	}

	str := fmt.Sprintf(contextDoesNotExistsErrorStr, context)
	return errors.New(str)
}

// Revert reverts the state of the last commit
func (app *instanceService) Revert() error {
	return app.revertToIndex(nil)
}

// Revert reverts the state of the commit to the provided index
func (app *instanceService) RevertToIndex(toIndex uint) error {
	return app.revertToIndex(&toIndex)
}

func (app *instanceService) revertToIndex(pIndex *uint) error {
	lastCommit, err := app.retrieveCommitAtIndex(pIndex)
	if err != nil {
		return err
	}

	pContext, err := app.Begin()
	if err != nil {
		return err
	}

	path := app.skeleton.Commit()
	hash := lastCommit.Hash()
	err = app.Delete(*pContext, path, hash)
	if err != nil {
		return err
	}

	return app.Commit(*pContext)
}

func (app *instanceService) retrieveCommitAtIndex(pIndex *uint) (commits.Commit, error) {
	commitPath := app.skeleton.Commit()
	commitHashesList, err := app.repository.ListByPath(commitPath)
	if err != nil {
		return nil, err
	}

	amountCommits := len(commitHashesList)
	if amountCommits <= 0 {
		return nil, nil
	}

	index := len(commitHashesList) - 1
	if pIndex != nil {
		index = int(*pIndex)
		if index >= amountCommits {
			str := fmt.Sprintf("the provided index (%d) exceeds the total commit amount (%d)", index, amountCommits)
			return nil, errors.New(str)
		}
	}

	lastHash := commitHashesList[index]
	lastCommit, err := app.repository.RetrieveByPathAndHash(commitPath, lastHash)
	if err != nil {
		return nil, err
	}

	if casted, ok := lastCommit.(commits.Commit); ok {
		return casted, nil
	}

	pathStr := strings.Join(commitPath, "/")
	str := fmt.Sprintf("the instance (path: %s, hash: %s) cannot be catsed to a Commit instance", pathStr, lastHash.String())
	return nil, errors.New(str)
}

func (app *instanceService) buildCommitContent(list []actions.Action) (commits.Content, error) {
	lastCommit, err := app.retrieveCommitAtIndex(nil)
	if err != nil {
		return nil, err
	}

	actions, err := app.actionsBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, err
	}

	builder := app.contentBuilder.Create().WithActions(actions)
	if lastCommit != nil {
		builder.WithPrevious(lastCommit)
	}

	return builder.Now()
}

func (app *instanceService) insertResource(
	tableName string,
	resourceToInsert commit_resources.Resource,
	resource resources.Resource,
	allResources resources.Resources,
	allConnections connections.Connections,
	pTx *sql.Tx,
) error {
	key := resource.Key()
	instance := resourceToInsert.Instance()
	keyName, keyValue, err := app.fetchFieldValue(tableName, instance, key, allResources, allConnections)
	if err != nil {
		return err
	}

	fields := resource.Fields()
	fieldValues, connectionFieldValues, err := app.fetchFieldsValueList(tableName, instance, fields, allResources, allConnections)
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
	_, err = pTx.Exec(queryStr, allFieldValues...)
	if err != nil {
		return err
	}

	err = app.insertConnectionValues(
		tableName,
		instance,
		fields,
		allResources,
		allConnections,
		connectionFieldValues,
		pTx,
	)

	if err != nil {
		return err
	}

	return nil
}

func (app *instanceService) insertConnectionValues(
	table string,
	ins instances.Instance,
	fields resources.Fields,
	allResources resources.Resources,
	allConnections connections.Connections,
	fieldValues map[string]interface{},
	pTx *sql.Tx,
) error {
	fromBytes := ins.Hash().Bytes()
	list := fields.List()
	for _, oneField := range list {
		fieldName := oneField.Name()
		if value, ok := fieldValues[fieldName]; ok {
			if casted, ok := value.(instances.Instance); ok {
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
						_, err = pTx.Exec(queryStr, fromBytes, toBytes)
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

			str := fmt.Sprintf("field: %s: the field was expected to contain an Instance instance", oneField.Name())
			return errors.New(str)
		}
	}

	return nil
}

func (app *instanceService) fetchFieldsValueList(
	table string,
	ins instances.Instance,
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

func (app *instanceService) fetchFieldValue(
	tableName string,
	ins instances.Instance,
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

func (app *instanceService) writeSchema(
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

func (app *instanceService) writeSchemaTable(
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

func (app *instanceService) writeSchemaTableFieldsList(
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

func (app *instanceService) writeSchemaTableField(
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

func (app *instanceService) writeSchemaTableFieldKind(
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

func (app *instanceService) writeSchemaConnectionTables(
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

func (app *instanceService) writeSchemaConnectionTable(
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

func (app *instanceService) writeSchemaConnectionField(
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

func (app *instanceService) writeSchemaTableFieldsForeignKeysList(
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

func (app *instanceService) writeSchemaTableFieldForeignKeysList(
	field field,
) string {
	kind := field.kind
	if kind.pForeignKey == nil {
		return ""
	}

	return app.writeSchemaForeignKey(*kind.pForeignKey)
}

func (app *instanceService) writeSchemaForeignKey(
	foreignKey foreignKey,
) string {
	return fmt.Sprintf(
		"FOREIGN KEY(%s) REFERENCES %s(%s) ON DELETE CASCADE",
		foreignKey.localField.name,
		foreignKey.foreignTableName,
		foreignKey.foreignField.name,
	)
}

func (app *instanceService) writeSchemaTableFieldKindNative(
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

func (app *instanceService) generateTables(
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

func (app *instanceService) generateTable(
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

func (app *instanceService) generateFields(
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

func (app *instanceService) generateField(
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
		canBeNil: fieldIns.CanBeNil(),
	}, nil
}

func (app *instanceService) generateFieldKind(
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

func (app *instanceService) createConnection(
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

func (app *instanceService) createForeignKey(
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
