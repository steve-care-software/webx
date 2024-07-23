# Instruction
An instruction represents a single command that is executed.  It can be a command that puts an assignment to the stack memory, or execute a command to the engine.

## Byte
The [byte](assignment/assignable/byte/readme.md) instruction executes the byte command then returns its value and stores it in the provided variable.
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
1. [Assignment](assignment/readme.md)