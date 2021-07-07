package httprouter

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestParams_Get(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		ps   Params
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.Get(tt.args.name); got != tt.want {
				t.Errorf("Params.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamsFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want Params
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamsFromContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParamsFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_GET(t *testing.T) {
	type fields struct {
		trees        map[string]*node
		PanicHandler func(http.ResponseWriter, *http.Request, interface{})
	}
	type args struct {
		path   string
		handle RouteHandler
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
			r := &Router{
				trees:        tt.fields.trees,
				PanicHandler: tt.fields.PanicHandler,
			}
			r.GET(tt.args.path, tt.args.handle)
		})
	}
}

func TestRouter_POST(t *testing.T) {
	type fields struct {
		trees        map[string]*node
		PanicHandler func(http.ResponseWriter, *http.Request, interface{})
	}
	type args struct {
		path   string
		handle RouteHandler
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
			r := &Router{
				trees:        tt.fields.trees,
				PanicHandler: tt.fields.PanicHandler,
			}
			r.POST(tt.args.path, tt.args.handle)
		})
	}
}

func TestRouter_PUT(t *testing.T) {
	type fields struct {
		trees        map[string]*node
		PanicHandler func(http.ResponseWriter, *http.Request, interface{})
	}
	type args struct {
		path   string
		handle RouteHandler
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
			r := &Router{
				trees:        tt.fields.trees,
				PanicHandler: tt.fields.PanicHandler,
			}
			r.PUT(tt.args.path, tt.args.handle)
		})
	}
}

func TestRouter_OPTIONS(t *testing.T) {
	type fields struct {
		trees        map[string]*node
		PanicHandler func(http.ResponseWriter, *http.Request, interface{})
	}
	type args struct {
		path   string
		handle RouteHandler
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
			r := &Router{
				trees:        tt.fields.trees,
				PanicHandler: tt.fields.PanicHandler,
			}
			r.OPTIONS(tt.args.path, tt.args.handle)
		})
	}
}

func TestRouter_AddRouteHandler(t *testing.T) {
	type fields struct {
		trees        map[string]*node
		PanicHandler func(http.ResponseWriter, *http.Request, interface{})
	}
	type args struct {
		method string
		path   string
		handle RouteHandler
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
			r := &Router{
				trees:        tt.fields.trees,
				PanicHandler: tt.fields.PanicHandler,
			}
			r.AddRouteHandler(tt.args.method, tt.args.path, tt.args.handle)
		})
	}
}

func TestRouter_recv(t *testing.T) {
	type fields struct {
		trees        map[string]*node
		PanicHandler func(http.ResponseWriter, *http.Request, interface{})
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
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
			r := &Router{
				trees:        tt.fields.trees,
				PanicHandler: tt.fields.PanicHandler,
			}
			r.recv(tt.args.w, tt.args.req)
		})
	}
}

func Test_enableCors(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enableCors(tt.args.w, tt.args.req)
		})
	}
}

func TestRouter_ServeHTTP(t *testing.T) {
	type fields struct {
		trees        map[string]*node
		PanicHandler func(http.ResponseWriter, *http.Request, interface{})
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
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
			r := &Router{
				trees:        tt.fields.trees,
				PanicHandler: tt.fields.PanicHandler,
			}
			r.ServeHTTP(tt.args.w, tt.args.req)
		})
	}
}
