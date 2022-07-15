package general

import (
	"reflect"
	"testing"
)

type test struct {
	Node1 string
	node2 int
	Node3 []int
}

var test1 = test{node2: 3}

func TestKeys(t *testing.T) {
	type args struct {
		T any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			args: args{
				map[any]int{"1": 2, 3: 4},
			},
			want: []any{"1", 3},
		},
		{
			args: args{
				test{},
			},
			want: []any{"node1", "node2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		src          any
		dst          string
		defaultValue any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			args: args{
				src:          map[string]int{"x": 1},
				dst:          "x",
				defaultValue: nil,
			},
			want: 1,
		},
		{
			args: args{
				src:          map[string]int{"t": 1},
				dst:          "x",
				defaultValue: nil,
			},
			want: nil,
		},
		{
			args: args{
				src:          test{node2: 2},
				dst:          "node2",
				defaultValue: nil,
			},
			want: nil,
		},
		{
			args: args{
				src:          test{Node3: []int{1, 2}},
				dst:          "Node3",
				defaultValue: nil,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Get(tt.args.src, tt.args.dst, tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			args: args{
				src: map[int]string{1: "1", 2: "2", 3: "3"},
			},
			want: []any{"1", "2", "3"},
		},
		{
			args: args{
				src: test1,
			},
			want: []any{test1.Node1, test1.Node3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Values(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
