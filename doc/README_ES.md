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
  * [Documentación de la librería "file"](#documentaci-n-de-la-librer-a--file-)
    + [Importación](#importaci-n)
    + [Función `CheckFile(file string) (bool, error)`](#funci-n--checkfile-file-string---bool--error--)
    + [Función `ReadFileCont(file string) ([]byte, error)`](#funci-n--readfilecont-file-string-----byte--error--)
    + [Función `ReWriteFile(file, text string) error`](#funci-n--rewritefile-file--text-string--error-)
    + [Función `GetBinaryDir() (string, error)`](#funci-n--getbinarydir----string--error--)

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

## Documentación de la librería "file"

El paquete `file` proporciona funciones para trabajar con archivos en Go (Golang).

### Importación

```go
import (
	"github.com/Tom5521/MyGolangTools/file"
)
```

### Función `CheckFile(file string) (bool, error)`

La función `CheckFile` verifica si un archivo existe en la ubicación proporcionada.

Parámetros:
- `file` (string): Ruta al archivo que se desea verificar.

Devuelve:
- `bool`: `true` si el archivo existe, `false` si no existe.
- `error`: Error si ocurre algún problema al verificar la existencia del archivo.

Ejemplo de uso:

```go
exists, err := file.CheckFile("archivo.txt")
if err != nil {
    // Manejo de error
}
if exists {
    fmt.Println("El archivo existe.")
} else {
    fmt.Println("El archivo no existe.")
}
```

### Función `ReadFileCont(file string) ([]byte, error)`

La función `ReadFileCont` lee el contenido de un archivo.

Parámetros:
- `file` (string): Ruta al archivo del que se desea leer el contenido.

Devuelve:
- `[]byte`: Contenido del archivo en forma de bytes.
- `error`: Error si ocurre algún problema al leer el archivo.

Ejemplo de uso:

```go
content, err := file.ReadFileCont("archivo.txt")
if err != nil {
    // Manejo de error
}
fmt.Println(string(content))
```

### Función `ReWriteFile(file, text string) error`

La función `ReWriteFile` crea un nuevo archivo o sobrescribe un archivo existente con el texto proporcionado.

Parámetros:
- `file` (string): Ruta al archivo que se creará o sobrescribirá.
- `text` (string): Texto que se escribirá en el archivo.

Devuelve:
- `error`: Error si ocurre algún problema al crear o escribir en el archivo.

Ejemplo de uso:

```go
err := file.ReWriteFile("nuevo_archivo.txt", "Contenido del archivo")
if err != nil {
    // Manejo de error
}
```

### Función `GetBinaryDir() (string, error)`

La función `GetBinaryDir` obtiene el directorio en el que se encuentra el ejecutable binario.

Devuelve:
- `string`: Ruta al directorio del ejecutable.
- `error`: Error si ocurre algún problema al obtener la ruta del ejecutable.

Ejemplo de uso:

```go
dir, err := file.GetBinaryDir()
if err != nil {
    // Manejo de error
}
fmt.Println("Directorio del ejecutable:", dir)
```
