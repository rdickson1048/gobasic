package main

import (
	"fmt"

	"gowork/src/github.com/rdickson1048/gobasic/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
