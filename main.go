package main

import (
	"fmt"

	"github.com/myugen/go-fizzbuzz-kata-solution/fizzbuzz"
)

func main() {
	evaluator := fizzbuzz.NewEvaluator()
	for counter := 1; counter <= 100; counter++ {
		fmt.Println(evaluator.Print(counter))
	}
}
