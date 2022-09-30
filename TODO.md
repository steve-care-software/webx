1. Finish composer structs
2. Finish finish composer syntax application
3. Create database application
    1. Contains a grammar to validate the file structure
    2. An insert trx contain a criteria to fetch an element from the database, a composer to add the new data at the right place, then use the compiler to replace the old data with the new data
    3. A delete trx contains a criteria
    4. An update contains a criteria + a composer + newdata
    5. Support atomic transactions


4. Finish blockchain daemon
5. Create blockchain repository and service on top of database application
6. Create blockchain module application
7. Create blockchain language application
8. Finish the syntax's program and execute methods in the syntax action application
9. Create tests in scripts to test the blockchain application
----
1. Finish the identity action application
    1. Assets transfers must first be pushed on blockchain, wait for confirmation by daemon, then sent to the other party
2. Finish the identity module
3. Finish the identity language
4. Create tests in scripts to test the identity application
----
1. Finish the syntax modules application
2. Finish the syntax language language
3. Create tests in scripts to test the syntax application
----
1. Create the REST server
2. Create the REST client for all applications
----
1. Create module to start/stop the daemon application
2. In the context of a started application, the scripting language must use the REST client applications, instead of local applications
