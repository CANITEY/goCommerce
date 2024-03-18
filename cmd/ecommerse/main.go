package main

import (
	"log"
	_ "ecommerce/internal/routes"
)


func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()


}

