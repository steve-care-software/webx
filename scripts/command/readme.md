# Command
Commands are scripts that are built-in inside our base logic.  They enable the execution of core commands into our engine in order to interact with others on the network, manipulate data, retrieve information or modify our internal database.

## CareScript Example
```
command: carescript -> core:layer {
    myVariable = byte.join [
        [34, 54, 67, 76],
        [33, 21, 23, 45],
    ];

    return myVariable|continue;
};


```

## Current sections:
1. [CareScript](carescript/readme.md)
2. [Identity](identity/readme.md)
3. [Network](network/readme.md)