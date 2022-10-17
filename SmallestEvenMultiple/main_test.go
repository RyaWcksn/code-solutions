package main

import "testing"

func Test_smallestEvenMultiple(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Success",
			args: args{
				n: 5,
			},
			want: 10,
		},
		{
			name: "Fail",
			args: args{
				n: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestEvenMultiple(tt.args.n); got != tt.want {
				t.Errorf("smallestEvenMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
