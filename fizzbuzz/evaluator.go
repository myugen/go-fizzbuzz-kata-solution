package fizzbuzz

import "strconv"

type evaluator struct {
	rules []PrinterRule
}

func (e evaluator) anyRuleIsSatisfied(value int) bool {
	for _, rule := range e.rules {
		if rule.Satisfy(value) {
			return true
		}
	}
	return false
}

func (e evaluator) Print(value int) string {
	if e.anyRuleIsSatisfied(value) {
		word := ""
		for _, rule := range e.rules {
			word += rule.Print(value)
		}
		return word
	}
	return strconv.Itoa(value)
}

func NewEvaluator() *evaluator {
	return &evaluator{rules: []PrinterRule{NewMultipleOfThreePrinterRule(), NewMultipleOfFivePrinterRule()}}
}
