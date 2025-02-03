package main

import (
	"fmt"
	"log"

	"github.com/bzelaznicki/gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("Bartek")

	if err != nil {
		log.Fatal(err)
	}

	newCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newCfg.DbURL)

}
