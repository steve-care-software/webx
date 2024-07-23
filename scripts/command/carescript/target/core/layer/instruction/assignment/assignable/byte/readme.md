# Byte
The byte code compiles a CareScript command into core json that can be embedded into instructions that can be executed by the core engine.

## Possible inputs
The byte command can take any kind of data directly, or variables.  When using data directly, it creates the assignments in json then put the generated variables into the join command in json.

## Join
The join command takes the input data and joins them together. 
```
// Source:
join [
    firstVariabke,
    "This is a string",
    342, // this is an unsigned integer
    -2, // this is an integer
    34.09, // this is a float
    [20, 34, 54, 67], // this is bytes
];
```

## Compare
The compare command takes the input data and compares them together.
```
// Source:
compare [
    firstVariable,
    "this is a string",
];
```

## Hash
The hash command takes the input parameter and returns its hash.
```
// Source:
hash myVariable;
hash "this is a string";
hash [34, 32, 34, 54];
```