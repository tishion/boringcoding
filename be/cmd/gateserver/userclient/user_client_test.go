package userclient

import (
	"net"
	"reflect"
	"testing"

	"mmc.com/landingtask/be/internal/lightrpc"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *UserClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserClient_Connect(t *testing.T) {
	type fields struct {
		conn net.Conn
		rpc  lightrpc.Conn
	}
	type args struct {
		conString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserClient{
				conn: tt.fields.conn,
				rpc:  tt.fields.rpc,
			}
			if got := s.Connect(tt.args.conString); got != tt.want {
				t.Errorf("UserClient.Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserClient_Close(t *testing.T) {
	type fields struct {
		conn net.Conn
		rpc  lightrpc.Conn
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserClient{
				conn: tt.fields.conn,
				rpc:  tt.fields.rpc,
			}
			if err := s.Close(); (err != nil) != tt.wantErr {
				t.Errorf("UserClient.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserClient_RpcInvoke(t *testing.T) {
	type fields struct {
		conn net.Conn
		rpc  lightrpc.Conn
	}
	type args struct {
		req []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserClient{
				conn: tt.fields.conn,
				rpc:  tt.fields.rpc,
			}
			got, err := s.RpcInvoke(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserClient.RpcInvoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserClient.RpcInvoke() = %v, want %v", got, tt.want)
			}
		})
	}
}
