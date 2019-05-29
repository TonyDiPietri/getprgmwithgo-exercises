//listing 7.5: time.go
package main

import (
	"fmt"
	"time"
)

func main() {
	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}
