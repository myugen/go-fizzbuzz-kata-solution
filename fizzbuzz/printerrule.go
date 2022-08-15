package fizzbuzz

type PrinterRule interface {
	Satisfy(value int) bool
	Print(value int) string
}
