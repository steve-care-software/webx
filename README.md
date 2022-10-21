# WebX
The webX application makes it easy to create a grammar, use that grammar to extract an AST from data, create criterias to extract data from an AST, combine AST's, create a compiler to compile a custom script to a program and/or execute a program and receive its output.

## VM syntax
    // module declaration:
    module @myModule;;

    // application declaration:
    @myModule $myApplication;;

    // this is an input parameter:
    -> $myInput;;

    // this is an output variable:
    <- $myOutput;;

    // assignment:
    $myVariable = ANY VALUE EXCEPT NON-ESCAPED SEMI-COLON;;

    // attach variable to application:
    attach $myDataVariable:$data $myAppVariable;;

    // execute module application and return its output:
    $myOutput = execute $myAppVariable;;

# Legit DAO
The documentation of the Legit DAO is coming soon
