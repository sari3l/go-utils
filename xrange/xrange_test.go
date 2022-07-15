package xrange

import (
	"reflect"
	"testing"
)

func TestXRange(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test",
			args: args{
				1, 3,
			},
			want: []int64{1, 2},
		},
		{
			name: "test2",
			args: args{
				1, 1,
			},
			want: []int64{},
		},
		{
			name: "test3",
			args: args{
				5, 1,
			},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XRange(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXRangeWithStep(t *testing.T) {
	type args struct {
		min  int64
		max  int64
		step int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			args: args{1, 10, 2},
			want: []int64{1, 3, 5, 7, 9},
		},
		{
			args: args{
				5, 1, 0,
			},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XRangeWithStep(tt.args.min, tt.args.max, tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XRangeWithStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				1, 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := range Range(tt.args.min, tt.args.max) {
				t.Log(i)
			}
		})
	}
}

func TestRangeWithStep(t *testing.T) {
	type args struct {
		min  int64
		max  int64
		step int64
	}
	tests := []struct {
		name string
		args args
		want <-chan int64
	}{
		{
			args: args{
				1, 10, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := range RangeWithStep(tt.args.min, tt.args.max, tt.args.step) {
				t.Log(i)
			}
		})
	}
}
