package math

import "testing"

func TestCountDigits(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "single digit", args: args{i: 5}, want: 1},
		{name: "multi-digit", args: args{i: 50}, want: 2},
		{name: "long", args: args{i: 123456789}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountDigits(tt.args.i); got != tt.want {
				t.Errorf("CountDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
