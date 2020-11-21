package main

import (
	"fmt"

	"github.com/AyanChoudhary/transponster/pkg/utils"
)

type Transponster struct {
}

func main() {
	config := utils.GenerateConfigReader("./config.json")
	fmt.Println(config.String("writer"))
}
