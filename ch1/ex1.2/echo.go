// Ex1.2 prints its commandline arguments and index 1 per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
}
