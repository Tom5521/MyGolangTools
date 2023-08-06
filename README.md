# MyGolangTools
My little golang tool kit


### Commands

This Go (Golang) library is called `tools` and provides functionalities for executing commands on different operating systems and shells. Below, I'll explain each of its features:

1. `Sh` struct:
   This structure represents a configuration for command execution. It has different fields to configure the execution on different operating systems and shells.

2. `Windows` field:
   - `PowerShell`: A boolean indicating whether to use PowerShell as the default shell on Windows systems. If `true`, PowerShell will be used; if `false`, the cmd shell will be used.

3. `Linux` field:
   - `CustomSh`: A nested field that allows configuring a custom shell on Linux systems.
      - `Enable`: A boolean indicating whether to enable the use of the custom shell defined in `ShName` and `ShArg`. If `true`, the custom shell specified in `ShName` and `ShArg` will be used.
      - `ShName`: The name of the custom shell to be used.
      - `ShArg`: The argument for executing the custom shell.

   - `Bash`: A boolean indicating whether to use Bash as the default shell on Linux systems. If `true`, Bash will be used; if `false`, the sh shell will be used.

4. `CustomStd` field:
   - `Enable`: A boolean indicating whether to configure the standard input, output, and error redirection options in a custom way. If `true`, the configurations defined for these fields will be used; if `false`, the default configurations will be used.
   - `Stdin`: A boolean indicating whether to redirect the standard input of the executed command.
   - `Stdout`: A boolean indicating whether to redirect the standard output of the executed command.
   - `Stderr`: A boolean indicating whether to redirect the standard error output of the executed command.

5. `Cmd(input string) (string, error)` method:
   This method takes an `input` argument, which is a string representing the command to be executed. It returns two values: the combined output of the command (stdout and stderr) as a string and an error if any issue occurred during command execution.

   Inside the method, it determines the operating system being used and selects the appropriate shell and arguments based on the configuration set in the `Sh` struct.

   If `CustomStd.Enable` is `true`, the standard input, output, and error of the command are configured based on the values of `Stdin`, `Stdout`, and `Stderr` fields. Otherwise, the default standard input, output, and error redirections (associated with the current console) are used.

6. `Out(input string) (string, error)` method:
   This method takes an `input` argument, which is a string representing the command to be executed. It returns two values: the combined output of the command (stdout and stderr) as a string and an error if any issue occurred during command execution.

   This method is a utility that calls the `Cmd` method with a specific configuration to not redirect the standard output (`Stdout` is disabled). In case of an error during execution, it returns the error. If there were no errors, it returns the combined output of the command.

In summary, this library provides a way to execute commands on different operating systems and shells, allowing the customization of standard input, output, and error options. It also offers specific methods to obtain only the output of the command without redirecting the standard output.
