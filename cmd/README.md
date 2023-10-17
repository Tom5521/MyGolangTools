
# Unix Library

The Unix library provides a set of functions for executing commands on Unix and Unix-like systems. It allows you to run both regular commands and superuser (sudo) commands.

## Main Functions

### `Cmd(command string) UnixCmd`

This function creates an instance of `UnixCmd` for running a regular command.

Example:

```go
sh := unix.Cmd("ls -l")
err := sh.Run()
```

### `Sudo_Cmd(command string, optional_password ...string) UnixSudoCmd`

This function creates an instance of `UnixSudoCmd` for running a command with superuser (sudo) privileges.

Example:

```go
sudoSh := unix.Sudo_Cmd("apt-get update", "your_password")
err := sudoSh.Run()
```

## Configuration Functions

### `SetInput(input string)`

This function sets the input for the command.

Example:

```go
sh := unix.Cmd("echo Hello")
sh.SetInput("echo World")
```

### `SetPath(path string)`

This function sets the working directory for the command.

Example:

```go
sh := unix.Cmd("ls -l")
sh.SetPath("/path/to/directory")
```

### `RunWithShell(set bool)`

This function enables or disables running the command with a shell.

Example:

```go
sh := unix.Cmd("echo Hello")
sh.RunWithShell(true)
```

### `CustomStd(Stdin, Stdout, Stderr bool)`

This function allows you to customize the input, output, and standard error of the command.

Example:

```go
sh := unix.Cmd("ls -l")
sh.CustomStd(true, true, true)
```

### `CustomShell(Shell_Name, Exec_Arg string)`

This function allows you to configure a custom shell to execute the command.

Example:

```go
sh := unix.Cmd("echo Hello")
sh.CustomShell("sh", "-c")
```

### `UseBashShell(set bool)`

This function enables or disables the use of the Bash shell to run the command.

Example:

```go
sh := unix.Cmd("echo Hello")
sh.UseBashShell(true)
```

## Execution Functions

The following functions execute the previously configured command:

### `Run() error`

Executes the command and returns an error if one occurs.

Example:

```go
sh := unix.Cmd("ls -l")
err := sh.Run()
```

### `Out() (string, error)`

Executes the command and returns the standard output as a string.

Example:

```go
sh := unix.Cmd("echo Hello")
out, err := sh.Out()
```

### `CombinedOut() (string, error)`

Executes the command and returns the combined standard output and error as a string.

Example:

```go
sh := unix.Cmd("ls non_existent_directory")
out, err := sh.CombinedOut()
```

### `Start() error`

Initiates the execution of the command in the background and returns an error if one occurs.

Example:

```go
sh := unix.Cmd("sleep 5")
err := sh.Start()
```

## Sudo Functions (for superuser commands)

The functions in `UnixSudoCmd` are similar to those in `UnixCmd` but are designed to run commands with superuser privileges.

Example of running a sudo command:

```go
sudoSh := unix.Sudo_Cmd("apt-get update", "your_password")
err := sudoSh.Run()
```

## Notes

- Ensure that your application has the necessary permissions to run sudo commands.
- Be cautious when using passwords in your source code, as this can pose a security risk.
- Make sure that the Unix library is imported and available in your Go project.

We hope this library is helpful for programmatically executing commands on Unix systems. If you encounter any issues or have suggestions, please don't hesitate to contact us.
```
