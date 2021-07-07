package servercore

import (
	"net"
	"reflect"
	"sync"
	"testing"
)

func TestCreateServer(t *testing.T) {
	type args struct {
		db string
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateServer(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	type fields struct {
		listener net.Listener
		clients  sync.Map
		pending  chan net.Conn
		leaving  chan net.Conn
	}
	type args struct {
		connString string
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
			s := &Server{
				listener: tt.fields.listener,
				clients:  tt.fields.clients,
				pending:  tt.fields.pending,
				leaving:  tt.fields.leaving,
			}
			s.Start(tt.args.connString)
		})
	}
}

func TestServer_Stop(t *testing.T) {
	type fields struct {
		listener net.Listener
		clients  sync.Map
		pending  chan net.Conn
		leaving  chan net.Conn
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				listener: tt.fields.listener,
				clients:  tt.fields.clients,
				pending:  tt.fields.pending,
				leaving:  tt.fields.leaving,
			}
			s.Stop()
		})
	}
}

func TestServer_prepareChannels(t *testing.T) {
	type fields struct {
		listener net.Listener
		clients  sync.Map
		pending  chan net.Conn
		leaving  chan net.Conn
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				listener: tt.fields.listener,
				clients:  tt.fields.clients,
				pending:  tt.fields.pending,
				leaving:  tt.fields.leaving,
			}
			s.prepareChannels()
		})
	}
}

func TestServer_listenAndServe(t *testing.T) {
	type fields struct {
		listener net.Listener
		clients  sync.Map
		pending  chan net.Conn
		leaving  chan net.Conn
	}
	type args struct {
		connString string
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
			s := &Server{
				listener: tt.fields.listener,
				clients:  tt.fields.clients,
				pending:  tt.fields.pending,
				leaving:  tt.fields.leaving,
			}
			s.listenAndServe(tt.args.connString)
		})
	}
}

func TestServer_createSession(t *testing.T) {
	type fields struct {
		listener net.Listener
		clients  sync.Map
		pending  chan net.Conn
		leaving  chan net.Conn
	}
	type args struct {
		conn net.Conn
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
			s := &Server{
				listener: tt.fields.listener,
				clients:  tt.fields.clients,
				pending:  tt.fields.pending,
				leaving:  tt.fields.leaving,
			}
			s.createSession(tt.args.conn)
		})
	}
}

func TestServer_removeSession(t *testing.T) {
	type fields struct {
		listener net.Listener
		clients  sync.Map
		pending  chan net.Conn
		leaving  chan net.Conn
	}
	type args struct {
		conn net.Conn
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
			s := &Server{
				listener: tt.fields.listener,
				clients:  tt.fields.clients,
				pending:  tt.fields.pending,
				leaving:  tt.fields.leaving,
			}
			s.removeSession(tt.args.conn)
		})
	}
}
