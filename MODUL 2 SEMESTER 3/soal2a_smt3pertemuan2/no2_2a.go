package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)
	if tahun%4 == 0 && tahun%100 != 0 {
		fmt.Print("Kabisat : True")
	} else {
		fmt.Print("Kabisat : False")
	}
}