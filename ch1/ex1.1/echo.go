// Ex1.1 prints its commandline arguments. including the name of the command to start running it
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
