package fizzbuzz

type multiplePrinterRule struct {
	multiple int
	word     string
}

func (f multiplePrinterRule) Satisfy(value int) bool {
	return value%f.multiple == 0
}

func (f multiplePrinterRule) Print(value int) string {
	if f.Satisfy(value) {
		return f.word
	}
	return ""
}

func NewMultipleOfThreePrinterRule() *multiplePrinterRule {
	return &multiplePrinterRule{
		multiple: 3,
		word:     "Fizz",
	}
}

func NewMultipleOfFivePrinterRule() *multiplePrinterRule {
	return &multiplePrinterRule{
		multiple: 5,
		word:     "Buzz",
	}
}
