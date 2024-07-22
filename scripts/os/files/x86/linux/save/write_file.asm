section .bss
    fd resb 4                             ; File descriptor

section .text
    global _start

_start:
    ; Set up the stack frame
    xor ebp, ebp                          ; Clear base pointer
    pop ebx                               ; Pop argc (argument count)
    pop ebx                               ; Pop argv[0] (program name)
    pop ebx                               ; Pop argv[1] (file path)
    pop ecx                               ; Pop argv[2] (data)
    pop edx                               ; Pop argv[3] (file mode)

    ; Open the file (O_WRONLY | O_CREAT | O_TRUNC, mode from argv[3])
    push edx                              ; Push mode for alignment
    mov eax, 5                            ; syscall number for sys_open
    mov ecx, 0x241                        ; flags (O_WRONLY | O_CREAT | O_TRUNC)
    mov edx, [esp]                        ; mode from argv[3]
    int 0x80                              ; call kernel

    mov [fd], eax                         ; store the file descriptor

    ; Write data to the file
    mov eax, 4                            ; syscall number for sys_write
    mov ebx, [fd]                         ; file descriptor
    mov ecx, ecx                          ; pointer to data (argv[2])
    mov edx, dword [ecx-4]                ; length of data (4 bytes before argv[2] in argv[1] length)
    int 0x80                              ; call kernel

    ; Close the file
    mov eax, 6                            ; syscall number for sys_close
    mov ebx, [fd]                         ; file descriptor
    int 0x80                              ; call kernel

    ; Exit the program
    mov eax, 1                            ; syscall number for sys_exit
    xor ebx, ebx                          ; status = 0
    int 0x80                              ; call kernel