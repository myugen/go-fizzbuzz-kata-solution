package fizzbuzz

import "strconv"

type evaluator struct {
	value int
}

func (f evaluator) Print() string {
	return strconv.Itoa(f.value)
}

func NewEvaluator(value int) *evaluator {
	return &evaluator{value: value}
}
