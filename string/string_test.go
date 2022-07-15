package string

import "testing"

func TestCapWords(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str: "test message",
				sep: " ",
			},
			want: "Test Message",
		},
		{
			args: args{
				str: "teSt meSsaGe",
				sep: " ",
			},
			want: "Test Message",
		},
		{
			args: args{
				str: "teST  meSsaGe",
				sep: "",
			},
			want: "Test Message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CapWords(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("CapWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
