package main

import "fmt"

func main() {

	var kelvin float32 = 373.0
	var celcius float32

	celcius = kelvin - 273

	fmt.Printf("O ponto de abulição da agua em Celsius é %.2f", celcius)

}
