// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)
func main() {
	println(os.Args[0], os.Args[1])
    for idx, i := range os.Args{
		fmt.Printf("%d%s\n",idx,i)
	}
}
