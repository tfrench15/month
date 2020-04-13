package main

import (
	"fmt"
	"time"
)

func main() {
	_, month, _ := time.Now().Date()
	fmt.Println(month)
}
