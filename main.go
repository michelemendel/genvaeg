// Playground

package main

import (
	"fmt"

	"github.com/michelemendel/genvaeg/util"
)

func main() {
	randomString := util.GenerateRandomString(10)
	fmt.Println("Random String:", randomString)
}
