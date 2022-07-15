package slice

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		src []int
		dst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				src: []int{1, 2, 3},
				dst: 1,
			},
			want: true,
		},
		{
			args: args{
				src: []int{1, 2, 3},
				dst: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.src, tt.args.dst); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	type args struct {
		src []int
		dst int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				src: []int{1, 2, 3},
				dst: 1,
			},
			want: 0,
		},
		{
			args: args{
				src: []int{1, 2, 3},
				dst: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.args.src, tt.args.dst); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}
