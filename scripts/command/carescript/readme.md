# CareScript
CareScript is a scripting language that can be lexed by the engine and compiled to a target format the engine understands.

## Example
```
carescript -> core:layer {
    myVariable = byte.join [
        [34, 54, 67, 76],
        [33, 21, 23, 45],
    ];

    return myVariable|continue;
};


```

## Reference
1. [Target](target/readme.md)