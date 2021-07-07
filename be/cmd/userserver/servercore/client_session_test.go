package servercore

import (
	"net"
	"reflect"
	"testing"

	"mmc.com/landingtask/be/internal/lightrpc"
)

func TestNewClientSession(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name string
		args args
		want *ClientSession
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientSession(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientSession_setEndHandler(t *testing.T) {
	type fields struct {
		conn    net.Conn
		rpc     lightrpc.Conn
		handler DisconnectHandler
	}
	type args struct {
		h DisconnectHandler
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
			s := &ClientSession{
				conn:    tt.fields.conn,
				rpc:     tt.fields.rpc,
				handler: tt.fields.handler,
			}
			s.setEndHandler(tt.args.h)
		})
	}
}

func TestClientSession_start(t *testing.T) {
	type fields struct {
		conn    net.Conn
		rpc     lightrpc.Conn
		handler DisconnectHandler
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ClientSession{
				conn:    tt.fields.conn,
				rpc:     tt.fields.rpc,
				handler: tt.fields.handler,
			}
			s.start()
		})
	}
}

func TestClientSession_sessionLoop(t *testing.T) {
	type fields struct {
		conn    net.Conn
		rpc     lightrpc.Conn
		handler DisconnectHandler
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ClientSession{
				conn:    tt.fields.conn,
				rpc:     tt.fields.rpc,
				handler: tt.fields.handler,
			}
			s.sessionLoop()
		})
	}
}

func TestClientSession_processMessage(t *testing.T) {
	type fields struct {
		conn    net.Conn
		rpc     lightrpc.Conn
		handler DisconnectHandler
	}
	type args struct {
		msg []byte
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
			s := &ClientSession{
				conn:    tt.fields.conn,
				rpc:     tt.fields.rpc,
				handler: tt.fields.handler,
			}
			s.processMessage(tt.args.msg)
		})
	}
}

func TestClientSession_send(t *testing.T) {
	type fields struct {
		conn    net.Conn
		rpc     lightrpc.Conn
		handler DisconnectHandler
	}
	type args struct {
		data []byte
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
			s := &ClientSession{
				conn:    tt.fields.conn,
				rpc:     tt.fields.rpc,
				handler: tt.fields.handler,
			}
			s.send(tt.args.data)
		})
	}
}
