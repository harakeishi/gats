package gats

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "int",
			args: args{
				a: int(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "int8",
			args: args{
				a: int8(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "int16",
			args: args{
				a: int16(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "int32",
			args: args{
				a: int32(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "int64",
			args: args{
				a: int64(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "uint",
			args: args{
				a: uint(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "uint8",
			args: args{
				a: uint8(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "uint16",
			args: args{
				a: uint16(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "uint32",
			args: args{
				a: uint32(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "uint64",
			args: args{
				a: uint64(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "float32",
			args: args{
				a: float32(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "float64",
			args: args{
				a: float64(1),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "complex64",
			args: args{
				a: complex64(1),
			},
			want:    "(1+0i)",
			wantErr: false,
		},
		{
			name: "complex128",
			args: args{
				a: complex128(1),
			},
			want:    "(1+0i)",
			wantErr: false,
		},
		{
			name: "byte",
			args: args{
				a: []byte("test"),
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "bool",
			args: args{
				a: bool(true),
			},
			want:    "true",
			wantErr: false,
		},
		{
			name: "string",
			args: args{
				a: "test",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "nil",
			args: args{
				a: nil,
			},
			want:    "nil",
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				a: []int{1, 2},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleToString() {
	value := int16(11)
	result, err := ToString(value)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	//Output: 11
}
