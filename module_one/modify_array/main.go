package main

import "fmt"

func main() {
	var me = [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println(me)
	for index, value := range me {
		if value == "stupid" {
			me[index] = "smart"
		} else if value == "weak" {
			me[index] = "strong"
		}
	}
	fmt.Println(me)
}
