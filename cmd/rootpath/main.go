package main

import (
	"fmt"
	"github.com/KoNekoD/rootpath/pkg/rootpath"
)

func main() {
	fmt.Printf("root path: %s\n", rootpath.MustDir())
}
