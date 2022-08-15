package fizzbuzz_test

import (
	"testing"

	"github.com/myugen/fizzbuzz-kata/fizzbuzz"
	"github.com/stretchr/testify/assert"
)

func TestEvaluator_Print(t *testing.T) {
	type fields struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should return the value when value is not multiple of 3 or multiple of 5",
			fields: fields{
				value: 1,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fizzbuzzEvaluator := fizzbuzz.NewEvaluator(tt.fields.value)
			got := fizzbuzzEvaluator.Print()
			assert.Equal(t, tt.want, got)
		})
	}
}
