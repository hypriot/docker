package main

import (
	"fmt"

	"github.com/dieterreuter/osarch"
)

func main() {
	fmt.Println("Arch()=", osarch.Arch())
	fmt.Println("OsArch()=", osarch.OsArch())
}
