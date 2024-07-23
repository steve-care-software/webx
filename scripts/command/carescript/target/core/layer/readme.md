# Layer
A layer represents a script that contains instructions that can be executed.  The engine learns how to execute layers based on remaining data in the input.

## Example
Here is a layer example that executes a [byte's join](instruction/assignment/assignable/byte/readme.md) instruction, then return its joined values as output.
```
layer {
    route {
       // the first omission is global, the second is local.  If there is no local, remove the @ as well.  Both are optional
       " "|" " @ " "|" ": [
            {
                // before the brackets, this is the omission before and after this specific token (optional)
                // the token is valid if any of these element matches
                " "|" ": [
                    // this is a token hash
                    361434c52eda44645178675eac8d630a942630504425a24aac1fa07fd74dc497f268ef585bcebded3f5b50facae47b431901345abe5354c84aec455c1baa9b85,
                    
                    // this is bytes
                    [23, 43, 45],

                    // this is a string:
                    "this is a string",
                ];
            }, {1,}, // This is the cardinality, this means that there must be at least 1 match, but can have unlimited.  If there is no cardinality, there must be exactly 1 match
            {
                " "|" ": [
                    361434c52eda44645178675eac8d630a942630504425a24aac1fa07fd74dc497f268ef585bcebded3f5b50facae47b431901345abe5354c84aec455c1baa9b85,
                    [23, 43, 45],
                    "this is a string",
                ];
            } {1,},
        ];
    };

    myVariable = byte.join [
        [34, 54, 67, 76],
        [33, 21, 23, 45],
    ];

    return myVariable|continue;
};

```

## Reference
1. [Instruction](instruction/readme.md)