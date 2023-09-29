

# Commands

## Overview

This Go module provides functions to execute shell commands and processes in Go. It supports running commands on Windows, Linux, and custom shells.

The main features are:

- Run commands with default shell or directly without a shell
- Configure options like shell type, sudo, standard streams 
- Execute and get output or start long-running processes
- Cross-platform support for Windows and Linux

## Usage

Import the module:

```go
import "github.com/Tom5521/MyGolangTools/commands"
```

Create a Sh struct to configure options:

```go
sh := commands.Sh{
  RunWithShell: true,
  Windows: {
    PowerShell: true, 
  },
  Linux: {
    RunWithSudo: true,
  },
}
```

Run a command:

```go 
err := sh.Cmd("ls -l")
```

Get command output:

```go
out, err := sh.Out("date")
```

Start a long-running process:

```go
sh.Start("sleep 10") 
```

## Sh Options

The Sh struct has the following configuration options:

- RunWithShell - Run the command with default system shell 
- Windows.Silent - Hide window when running commands on Windows
- Windows.PowerShell - Use PowerShell instead of cmd on Windows 
- Linux.RunWithSudo - Prefix commands with sudo on Linux
- Linux.Bash - Use bash instead of sh on Linux
- Linux.CustomSh - Configure a custom shell on Linux
- CustomStd - Configure standard streams

## Functions

The module has the following functions:

- Cmd - Execute a command
- Out - Execute and return output 
- Start - Start a long-running process

Cmd and Out return an error. Out also returns the command output string.

## Examples

Run a PowerShell command on Windows:

```go
sh := commands.Sh{
  RunWithShell: true,
  Windows: {
    PowerShell: true,
  },
}

sh.Cmd("Get-Process")
```

Run a command with sudo on Linux:

```go
sh := commands.Sh{
  RunWithShell: true,
  Linux: {
    RunWithSudo: true,
  },
}

sh.Cmd("apt update") 
```

Start a background process:

```go
sh := commands.Sh{} 

sh.Start("./app")
```

This provides a basic Go module to run shell commands and processes in a cross-platform way. The options and functions allow customizing the execution environment.


# File

## Overview

This Go module provides utility functions for working with files in Go.

The main features are:

- Get file size for files and folders 
- Check if a file exists
- Read file contents into a byte slice
- Write string data to a file  
- Get executable binary directory

## Functions

The module contains the following functions:

### FileSize

Get the size of a file or folder recursively.

```go
size, err := file.FileSize("folder")
```

Returns size in bytes and an error.

### CheckFile

Check if a file exists.

```go 
exists, err := file.CheckFile("file.txt")
``` 

Returns a boolean and an error.

### ReadFileCont

Read a file into a byte slice.

```go
data, err := file.ReadFileCont("file.txt") 
```

Returns the file contents as a byte slice and an error.

### ReWriteFile

Write string data to a file, overwriting existing contents.

```go
err := file.ReWriteFile("file.txt", "hello world")
```

Returns an error.

### GetBinaryDir

Get the directory of the running executable binary. 

```go 
dir, err := file.GetBinaryDir()
```

Returns the directory string and an error.

## Examples

Read a file:

```go
data, _ := file.ReadFileCont("data.bin")
```

Check if a file exists:

```go 
exists, _ := file.CheckFile("file.txt")
if exists {
  // do something
}
```

Get executable directory:

```go
dir, _ := file.GetBinaryDir()
fmt.Println(dir)
```

This provides basic file manipulation functions in Go. The functions return errors for robust handling.
