# Assignment
An assignment reference a variable name and a assignable.  After executing the assignable command, it assigns its value to the provided variable name.

## Byte
The [byte](assignable/byte/readme.md) assignment executes the byte command then returns its value and stores it in the provided variable.
```
myVariable = byte.join [
    firstVariabke,
    "This is a string",
    342, // this is an unsigned integer
    -2, // this is an integer
    34.09, // this is a float
    [20, 34, 54, 67], // this is bytes
];
```

## Reference
1. [Assignable](assignable/readme.md)