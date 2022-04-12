package main

import "fmt"

func main() {
	MyArry := [5]string{"I", "am", "stupid", "and", "week"}
	for k, _ := range MyArry {
		switch MyArry[k] {
		case "stupid":
			MyArry[k] = "smart"
		case "week":
			MyArry[k] = "strong"
		}
		fmt.Println(MyArry[k])
	}
}
