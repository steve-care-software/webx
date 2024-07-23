# Layer
A layer represents a script that contains instructions that can be executed.  The engine learns how to execute layers based on remaining data in the input.

## Example
Here is a layer example that executes a [byte's join](instruction/assignment/assignable/byte/readme.md) instruction, then return its joined values as output.
```
layer {
    myVariable = byte.join [
        [34, 54, 67, 76],
        [33, 21, 23, 45],
    ];

    return myVariable|continue;
};


```

## Reference
1. [Instruction](instruction/readme.md)