package servercore

import (
	"net/http"
	"reflect"
	"testing"

	"mmc.com/landingtask/be/cmd/gateserver/httprouter"
	"mmc.com/landingtask/be/cmd/gateserver/userclient"
)

func Test_userClientFactory(t *testing.T) {
	tests := []struct {
		name    string
		want    *userclient.UserClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userClientFactory()
			if (err != nil) != tt.wantErr {
				t.Errorf("userClientFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userClientFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRPCClientPool(t *testing.T) {
	tests := []struct {
		name string
		want *userclient.ClientPool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRPCClientPool(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRPCClientPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateServer(t *testing.T) {
	type args struct {
		rpcAddress string
		cache      string
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
			if got := CreateServer(tt.args.rpcAddress, tt.args.cache); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	type fields struct {
		router *httprouter.Router
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
				router: tt.fields.router,
			}
			s.Start(tt.args.connString)
		})
	}
}

func Test_getHello(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getHello(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}

func Test_postHello(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postHello(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}

func Test_login(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			login(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}

func Test_getProfile(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getProfile(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}

func Test_putProfile(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			putProfile(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}
