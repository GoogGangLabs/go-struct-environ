## go-struct-environ

It's a package that helps you set environment variables easily in go.<br>

It is also fully compatible with the process environment variable (os.Getenv).

<br>

## Usage

### Install Package

```Shell
$ go get https://github.com/GoogGangLabs/go-struct-environ
```

<br>

### Import Package

```Go
import (
  environ "https://github.com/GoogGangLabs/go-struct-environ"
)
```

<br>

### Function Declaration

```Go
/*
  (Parameter 1) The path variable allows both absolte and relative paths.

  (Parameter 2) The environStrute variable must be of the pointer struct type.

  (Return) If an error occurs, an error type is returned.
*/
func Load(path string, environStruct interface{}) (error)
```

<br>

## Example

<br>

`.env`

```Plain Text
SERVER_HOST=127.0.0.1
SERVER_PORT=5000
```

<br>

`main.go`

```Go
package main

import (
  "fmt"

  environ "https://github.com/GoogGangLabs/go-struct-environ"
)

type Environ struct {
  SERVER_HOST string
  SERVER_PORT int
}

func main() {
  envStruct := Environ{}

  fmt.Println(envStruct)

  /* Load environment variables from Environment file */
  err := environ.LoadFromFile("./.env", &envStruct)
  /* Or load from OS Environment variables */
  err := environ.LoadFromEnvironment(&envStruct)
  if err != nil { panic(err) }

  fmt.Println(envStruct)
}
```

<br>

`Result`

```Shell
$ go run main.go
{ 0} # First output
{127.0.0.1 5000} # Second output
```
