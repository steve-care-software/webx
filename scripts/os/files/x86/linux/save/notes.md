# Command
```
"./this/is/a/path/myfile.json" 0644 [34, 54, 56, 65, 67, 76, 78]
```

# Compile the assembly to machine code:
```
nasm -f elf32 write_file.asm -o write_file.o
ld -m elf_i386 write_file.o -o write_file
```

# Execute the script
```
./write_file "output.txt" "Hello, World!" 0644
```