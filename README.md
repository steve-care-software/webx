# Syntax
The syntax application makes it easy to create a grammar, use that grammar to extract an AST from data, create criterias to extract data from an AST, combine AST's, create a compiler to compile a custom script to a program and/or execute a program and receive its output.

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

## Grammar syntax

    root: myRoot;
    - myFirstChannel prev:next;
    - mySecondChannel :next;
    - myThirdChannel prev;
    - myFourthChannel;

    myRevTokenInstance: @myToken; // @ = reverse
    myEverythingExceptInstance: -(firstException secondException thirdException):(firstEscape secondEscape);
    myEverythingExceptInstance: -(firstException secondException thirdException); // no escape

    myEscape: firstEscape secondEscape
    		;

    myToken : myExternal
    		 | myValue
    		 ;

    myExternal: 79314fd8d3d2c692c06862b5ea0b3fd7fde18cfd272303824a7a51eab56e0334b857590ba7963d0e1c5c38e46d2a3a578b9ea9e2267299662a949346ea3c7ffd;
    myValue: 45;

    ===
    myValue:
    	- this \; is some data;
    	- this is other valid data;
    	! this is invalid data \; ;
    	! Some other invalid data;
    ;
