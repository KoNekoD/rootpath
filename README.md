# RootPath

Go library to find the path to the root of the project where the `go.mod` file is located. 
Can also change the current directory to the project directory path

## Usage

```go
package main

import (
	"fmt"
	"github.com/KoNekoD/rootpath/pkg/rootpath"
)

func main() {
	fmt.Printf("root path: %s\n", rootpath.MustDir())
}
```