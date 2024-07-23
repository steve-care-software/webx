# Target
The target represents all the possible target formats the CareScript can be compiled to.

## Example
```
-> core:layer {
    myVariable = byte.join [
        [34, 54, 67, 76],
        [33, 21, 23, 45],
    ];

    return myVariable|continue;
};


```

## Reference
1. [Core](core/readme.md)