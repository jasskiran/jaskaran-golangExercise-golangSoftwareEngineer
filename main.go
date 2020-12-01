package main

import (
	"Ayoconnect/bianrySearchTree"
	r "Ayoconnect/robbery"
	"fmt"
)



func main() {
	fmt.Println("Binary search tree: ")
	arr, err := getInputArray()
	if err != nil{
		panic(err)
	}
	BST.BinarySearchTree(arr)


	fmt.Printf("Robbery: ")
	arr, err = getInputArray()
	if err != nil{
		panic(err)
	}
	res := r.Robbery(arr)
	fmt.Println(res)
}

func getInputArray() ([]int, error) {
	length := 0
	fmt.Println("Enter values to create array: ")
	_, err := fmt.Scanln(&length)
	if err != nil {
		return nil, err
	}
	arr := make([]int, length)
	fmt.Println(arr, "arr")
	for i := 0; i < length; i++ {
		_, err := fmt.Scanln(&arr[i])
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}