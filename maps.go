package main

import "fmt"

func main() {

	fmt.Println("----- MAPS -----")
	mp := make(map[string]int)

	mp["marky"] = 15
	sampleKey := "id"
	mp[sampleKey] = 124

	fmt.Println(mp)
	delete(mp, sampleKey)

	val, ok := mp["marky"]

	fmt.Println(val, ok)
	fmt.Println(mp)
}
