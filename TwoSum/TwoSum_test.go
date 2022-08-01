package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		array   []int
		target  int
		expect  []int
		wantErr bool
	}{
		{
			name:    "Success",
			array:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:  9,
			expect:  []int{3, 4},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TwoSum(test.array, test.target)
			if !reflect.DeepEqual(got, test.expect) && test.wantErr == false {
				t.Fatalf("Error, Err := %v want no err", got)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	var buf bytes.Buffer
	for n := 1; n < 1024; n *= 2 {
		b.Run(fmt.Sprintf("%v", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf.Reset()
				arr := []int{1, 4, 2, 3, 5, 6, 7, 8, 4, 2, 5, 6}
				target := 15
				TwoSum(arr, target)
			}
		})
	}
}
