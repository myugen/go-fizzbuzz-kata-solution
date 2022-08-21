package fizzbuzz_test

import (
	"testing"

	"github.com/myugen/go-fizzbuzz-kata-solution/fizzbuzz"
	"github.com/stretchr/testify/assert"
)

func TestEvaluator_Print(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return the value when value is not multiple of 3 or multiple of 5",
			args: args{value: 1},
			want: "1",
		},
		{
			name: "should return the Fizz word when value is multiple of 3",
			args: args{value: 3},
			want: "Fizz",
		},
		{
			name: "should return the Buzz word when value is not multiple of 5",
			args: args{value: 5},
			want: "Buzz",
		},
		{
			name: "should return the FizzBuzz word when value is not multiple of 3 and multiple of 5",
			args: args{value: 15},
			want: "FizzBuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fizzbuzzEvaluator := fizzbuzz.NewEvaluator()
			got := fizzbuzzEvaluator.Print(tt.args.value)
			assert.Equal(t, tt.want, got)
		})
	}
}
