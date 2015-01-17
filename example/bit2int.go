package main

import (
	"flag"
	"fmt"
	"github.com/ChaimHong/bit2int"
)

func main() {
	var data = flag.Int("v", 0, "bits")
	flag.Parse()

	fmt.Println(bit2int.BToI(*data))
}
