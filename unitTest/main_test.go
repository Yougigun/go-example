package main

import (
	"testing"
)

func Test_reString_UnmarshalJSON(t *testing.T) {
	type args struct {
		bs []byte
	}
	r := reString("")
	tests := []struct {
		name    string
		tr      *reString
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "wo error",
			tr:      &r,
			args:    args{bs: []byte(`"/d+"`)},
			wantErr: false,
		},
		{
			name:    "w/ error case in json unmarshal ",
			tr:      &r,
			args:    args{bs: []byte(`"aaa`)},
			wantErr: true,
		},
		{
			name:    "w/ error case in regexp compile ",
			tr:      &r,
			args:    args{bs: []byte(`"[d+"`)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.UnmarshalJSON(tt.args.bs); (err != nil) != tt.wantErr {
				t.Errorf("reString.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "just for coverage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
