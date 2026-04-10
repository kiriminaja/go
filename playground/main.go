package main

import (
	"fmt"
	"log"
	"os"

	kiriminaja "github.com/kiriminaja/go"
)

func main() {
	client := kiriminaja.New(kiriminaja.Config{
		Env:    kiriminaja.EnvSandbox,
		APIKey: os.Getenv("KIRIMINAJA_API_KEY"),
	})

	provinces, err := client.Address.Provinces()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Provinces: %+v\n", provinces)
}
