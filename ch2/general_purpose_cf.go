package ch2

import (
	"fmt"
	"log"
)

func unitConversion() string {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
	return Fahrenheit(3).String()
	// for _, arg := range os.Args[1:] {
	// 	t, err := strconv.ParseFloat(arg, 64)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	f := tempconv.Fahrenheit(t)
	// 	c := tempconv.Celsius(t)
	// 	fmt.Printf("%s = %s, %s = %s\n",
	// 		f, tempconv.FToC(f), c, tempconv.CToF(c))
	// }
}
