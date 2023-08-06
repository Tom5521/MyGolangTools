# MyGolangTools


- [MyGolangTools](#mygolangtools)
  * [Documentación de la librería "commands"](#documentaci-n-de-la-librer-a--commands-)
    + [Descripción](#descripci-n)
    + [Uso básico](#uso-b-sico)
    + [Métodos](#m-todos)
      - [1. `Cmd(input string) error`](#1--cmd-input-string--error-)
      - [2. `Out(input string) (string, error)`](#2--out-input-string---string--error--)
    + [Configuración avanzada](#configuraci-n-avanzada)
      - [1. Configuración del shell](#1-configuraci-n-del-shell)
      - [2. Configuración del shell personalizado en Linux](#2-configuraci-n-del-shell-personalizado-en-linux)
      - [3. Configuración de entrada/salida estándar personalizada](#3-configuraci-n-de-entrada-salida-est-ndar-personalizada)
    + [Ejemplo completo](#ejemplo-completo)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>



## Documentación de la librería "commands"

### Descripción

La librería "commands" proporciona una interfaz simple para ejecutar comandos en sistemas operativos Windows y Linux desde una aplicación escrita en Golang. La librería se encarga de manejar la elección del shell adecuado según el sistema operativo y permite la configuración de opciones personalizadas de entrada y salida estándar.

### Uso básico

1. Importar la librería en tu programa Golang:

```go
package main

import "github.com/Tom5521/MyGolangTools/commands"
```

2. Crear una instancia de la estructura `Sh` para utilizar los métodos proporcionados:

```go
sh := commands.Sh{}
```

### Métodos

#### 1. `Cmd(input string) error`

El método `Cmd` se utiliza para ejecutar un comando en el shell seleccionado según el sistema operativo.

Parámetros:
- `input` (string): El comando que se va a ejecutar.

Retorno:
- `error`: Si ocurre algún error durante la ejecución del comando.

Ejemplo de uso:

```go
sh.Windows.PowerShell = true // Utilizar PowerShell en Windows
err := sh.Cmd("ls -l")
if err != nil {
    fmt.Println("Error ejecutando el comando:", err)
}
```

#### 2. `Out(input string) (string, error)`

El método `Out` se utiliza para ejecutar un comando en el shell seleccionado y capturar la salida del mismo.

Parámetros:
- `input` (string): El comando que se va a ejecutar.

Retorno:
- `string`: La salida generada por el comando.
- `error`: Si ocurre algún error durante la ejecución del comando.

Ejemplo de uso:

```go
sh.Windows.PowerShell = false // Utilizar shell cmd en Windows
out, err := sh.Out("echo 'Hola, mundo!'")
if err != nil {
    fmt.Println("Error ejecutando el comando:", err)
} else {
    fmt.Println("Salida del comando:", out)
}
```

### Configuración avanzada

La librería "commands" permite configurar algunas opciones avanzadas para personalizar la ejecución de los comandos:

#### 1. Configuración del shell

Puedes especificar el shell que deseas utilizar en sistemas Windows y Linux mediante el campo `Windows.PowerShell` y `Linux.Bash`, respectivamente.

Ejemplo:

```go
sh.Windows.PowerShell = true // Utilizar PowerShell en Windows
sh.Linux.Bash = true // Utilizar Bash en Linux
```

#### 2. Configuración del shell personalizado en Linux

En sistemas Linux, puedes habilitar un shell personalizado y especificar su nombre y argumentos mediante el campo `Linux.CustomSh`.

Ejemplo:

```go
sh.Linux.CustomSh.Enable = true
sh.Linux.CustomSh.ShName = "/usr/bin/zsh"
sh.Linux.CustomSh.ShArg = "-c"
```

#### 3. Configuración de entrada/salida estándar personalizada

Puedes personalizar las opciones de entrada y salida estándar para los comandos mediante el campo `CustomStd`.

Ejemplo:

```go
sh.CustomStd.Enable = true
sh.CustomStd.Stdin = true // Habilitar entrada estándar
sh.CustomStd.Stdout = true // Habilitar salida estándar
sh.CustomStd.Stderr = true // Habilitar salida de error estándar
```

### Ejemplo completo

A continuación, se muestra un ejemplo completo de cómo utilizar la librería "commands" con diferentes configuraciones:

```go
package main

import (
	"fmt"
    "github.com/Tom5521/MyGolangTools/commands"	
)

func main() {
	sh := commands.Sh{}

	sh.Windows.PowerShell = true // Utilizar PowerShell en Windows
	err := sh.Cmd("ls -l")
	if err != nil {
		fmt.Println("Error ejecutando el comando:", err)
	}

	sh.Windows.PowerShell = false // Utilizar shell cmd en Windows
	out, err := sh.Out("echo 'Hola, mundo!'")
	if err != nil {
		fmt.Println("Error ejecutando el comando:", err)
	} else {
		fmt.Println("Salida del comando:", out)
	}
}
```
