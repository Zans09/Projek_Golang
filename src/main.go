package main

import (
	"fmt"
)

type Operation interface {
	Calculate() int
}

type Addition struct {
	a, b int
}

func (add Addition) Calculate() int {
	return add.a + add.b
}

type Multiplication struct {
	a, b int
}

func (mul Multiplication) Calculate() int {
	return mul.a * mul.b
}

func performOperation(op Operation, resultChan chan int) {
	result := op.Calculate()
	resultChan <- result
}

func main() {
	a := 3
	b := 5

	addition := Addition{a, b}
	multiplication := Multiplication{a, b}

	resultChan := make(chan int, 2)

	go performOperation(addition, resultChan)
	go performOperation(multiplication, resultChan)

	addResult := <-resultChan
	mulResult := <-resultChan
	
	fmt.Printf("Penjumlahan: %d\n", addResult)
	fmt.Printf("Perkalian: %d\n", mulResult)
}

// cara menguji hasil codingan golang 
// 1. Buka Terminal
// 2. Akses "New Terminal" 
// 3. Setelah itu, jalankan command "go run main.go"