package main

import (
	"fmt"
	"github.com/spf13/viper"
	"listener/router"
	"os"
	"path/filepath"
)

func main() {
	p, _ := filepath.Abs(fmt.Sprintf("%s/%s", filepath.Dir(os.Args[0]), "./public"))
	fmt.Println(p)
	viper.SetDefault("public", p)
	r := router.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
