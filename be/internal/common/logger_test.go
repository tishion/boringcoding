package common

import (
	"reflect"
	"testing"
)

func Test_getLevleString(t *testing.T) {
	type args struct {
		lvl Level
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLevleString(tt.args.lvl); got != tt.want {
				t.Errorf("getLevleString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateLogger(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLogger(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_logWithLevel(t *testing.T) {
	type fields struct {
		Name   string
		MinLvl Level
	}
	type args struct {
		lvl    Level
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Name:   tt.fields.Name,
				MinLvl: tt.fields.MinLvl,
			}
			l.logWithLevel(tt.args.lvl, tt.args.format, tt.args.v...)
		})
	}
}

func TestLogger_D(t *testing.T) {
	type fields struct {
		Name   string
		MinLvl Level
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Name:   tt.fields.Name,
				MinLvl: tt.fields.MinLvl,
			}
			l.D(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogger_I(t *testing.T) {
	type fields struct {
		Name   string
		MinLvl Level
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Name:   tt.fields.Name,
				MinLvl: tt.fields.MinLvl,
			}
			l.I(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogger_W(t *testing.T) {
	type fields struct {
		Name   string
		MinLvl Level
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Name:   tt.fields.Name,
				MinLvl: tt.fields.MinLvl,
			}
			l.W(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogger_E(t *testing.T) {
	type fields struct {
		Name   string
		MinLvl Level
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Name:   tt.fields.Name,
				MinLvl: tt.fields.MinLvl,
			}
			l.E(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogger_F(t *testing.T) {
	type fields struct {
		Name   string
		MinLvl Level
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Name:   tt.fields.Name,
				MinLvl: tt.fields.MinLvl,
			}
			l.F(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogD(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogD(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogI(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogI(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogW(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogW(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogE(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogE(tt.args.format, tt.args.v...)
		})
	}
}

func TestLogF(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogF(tt.args.format, tt.args.v...)
		})
	}
}
