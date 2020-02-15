package main

import (
	"fmt"

	"github.com/linkITworld/goexample0/04_scope/01_package-scope/vis"
)
linkITworld/goexample0/start/model
func main() {
	fmt.Println(vis.CatName + " and " + vis.MouseName)
	vis.PrintVar()
}
