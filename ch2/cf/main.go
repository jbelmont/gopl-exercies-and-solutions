package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jbelmont/gopl.io.challenges/ch2/tempconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		var response string
		fmt.Print("Enter Temperature: ")
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal(err)
		}
		temp, err := strconv.ParseFloat(response, 64)
		val := tempconv.Fahrenheit(temp).String()
		fmt.Println(val)
	} else {
		for _, val := range argsWithoutProg {
			temp, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}
			val := tempconv.Fahrenheit(temp).String()
			fmt.Println(val)
		}
	}
}
