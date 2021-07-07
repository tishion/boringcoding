package lightrpc

import (
	"io"
	"reflect"
	"testing"
)

func Test_packetHeaderSplitter(t *testing.T) {
	type args struct {
		data []byte
		eof  bool
	}
	tests := []struct {
		name        string
		args        args
		wantAdvance int
		wantToken   []byte
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdvance, gotToken, err := packetHeaderSplitter(tt.args.data, tt.args.eof)
			if (err != nil) != tt.wantErr {
				t.Errorf("packetHeaderSplitter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAdvance != tt.wantAdvance {
				t.Errorf("packetHeaderSplitter() gotAdvance = %v, want %v", gotAdvance, tt.wantAdvance)
			}
			if !reflect.DeepEqual(gotToken, tt.wantToken) {
				t.Errorf("packetHeaderSplitter() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func TestLightRpcConn(t *testing.T) {
	type args struct {
		rw io.ReadWriter
	}
	tests := []struct {
		name string
		args args
		want Conn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LightRpcConn(tt.args.rw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LightRpcConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_Send(t *testing.T) {
	type fields struct {
		io io.ReadWriter
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				io: tt.fields.io,
			}
			if err := c.Send(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Conn.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConn_Receive(t *testing.T) {
	type fields struct {
		io io.ReadWriter
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				io: tt.fields.io,
			}
			got, err := c.Receive()
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.Receive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.Receive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_sendPacketWithData(t *testing.T) {
	type fields struct {
		io io.ReadWriter
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				io: tt.fields.io,
			}
			if err := c.sendPacketWithData(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Conn.sendPacketWithData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
