# Core
The core target takes a CareScript as input and compiles it to the json format that can be executed by the core engine.

## Example
```
core:layer {
    myVariable = byte.join [
        [34, 54, 67, 76],
        [33, 21, 23, 45],
    ];

    return myVariable|continue;
};


```

## Reference
1. [Layer](layer/readme.md)