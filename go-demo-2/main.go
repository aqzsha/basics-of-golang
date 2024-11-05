package main

import "fmt"

func main() {
	arr := []int{}
	for {
		var num int
		fmt.Print("Input the number: ")
		fmt.Scan(&num)
		if(num==0){
			break
		}
		arr = append(arr, num)
	}

	// for index, item :=range arr{
	// 	fmt.Println(index, item)
	// }
}