package leetcode

import "testing"

func Test_climbStairs01(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs01(tt.args.n); got != tt.want {
				t.Errorf("climbStairs01() = %v, want %v", got, tt.want)
			}
		})
	}
}
