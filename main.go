package main

import (
	"github.com/jtyler139/blogaggregator/internal/config"
	"fmt"
)

func main() {
	cfg := config.Config{}
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
	cfg.SetUser("Tyler")
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}