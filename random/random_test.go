package random

import (
	"testing"
)

func TestChoice(t *testing.T) {
	type args struct {
		seq []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				seq: []any{"string_1_1", "string_1_2"},
			},
		},
		{
			name: "test2",
			args: args{
				seq: []any{111111, 222222},
			},
		},
		{
			name: "test3",
			args: args{
				seq: []any{333333, "string_3_1", 3.3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Choice(tt.args.seq)
			t.Logf("Choice() = %v", got)
		})
	}
}

func TestChoices(t *testing.T) {
	type args struct {
		seq []string
		k   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				seq: []string{"string_1_1", "string_1_2"},
				k:   0,
			},
		},
		{
			name: "test2",
			args: args{
				seq: []string{"string_1_1", "string_1_2"},
				k:   1,
			},
		},
		{
			name: "test3",
			args: args{
				seq: []string{"string_1_1", "string_1_2"},
				k:   5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Choices(tt.args.seq, tt.args.k)
			t.Logf("Choice() = %v", got)
		})
	}
}
