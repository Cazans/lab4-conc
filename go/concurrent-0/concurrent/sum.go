package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Result struct {
	sum int
	path string
}

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filePath string, canal chan Result) {
	data, _ := readFile(filePath)
	

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	canal <- Result{_sum, filePath}
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64

	canal := make(chan Result, len(os.Args[1]))
	
	sums := make(map[int][]string)

	for _, path := range os.Args[1:] {
		go sum(path, canal)
	}

	for range os.Args[1:] {
		result_n := <- canal 
			
		sums[result_n.sum] = append(sums[result_n.sum], result_n.path)
		totalSum += int64(result_n.sum)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}
