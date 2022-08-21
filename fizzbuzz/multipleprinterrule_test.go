package fizzbuzz_test

import (
	"testing"

	"github.com/myugen/go-fizzbuzz-kata-solution/fizzbuzz"
	"github.com/stretchr/testify/assert"
)

func Test_MultipleOfThreePrinterRule_Print(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return an empty word when value is not multiple of 3",
			args: args{value: 2},
			want: "",
		},
		{
			name: "should return the Fizz word when value is not multiple of 3",
			args: args{value: 3},
			want: "Fizz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fizzbuzz.NewMultipleOfThreePrinterRule().Print(tt.args.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_MultipleOfFivePrinterRule_Print(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return an empty word when value is not multiple of 5",
			args: args{value: 4},
			want: "",
		},
		{
			name: "should return the Buzz word when value is not multiple of 5",
			args: args{value: 5},
			want: "Buzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fizzbuzz.NewMultipleOfFivePrinterRule().Print(tt.args.value)
			assert.Equal(t, tt.want, got)
		})
	}
}
