package main

import "fmt"

type nomadCoder struct {
	name    string
	age     int
	job     string
	favFood map[string]string
}

func praticeMap() {
	favFood := map[string]string{"vegetable": "apple"}
	m := nomadCoder{"Thai", 22, "Developer", favFood}
	n := nomadCoder{name: "Dương"}
	fmt.Println(m)
	fmt.Println(m.favFood)
	fmt.Println(n)

}
