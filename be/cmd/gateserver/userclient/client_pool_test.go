package userclient

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewClientPool(t *testing.T) {
	type args struct {
		minOpen int
		maxOpen int
		factory ClientFactory
	}
	tests := []struct {
		name string
		args args
		want *ClientPool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientPool(tt.args.minOpen, tt.args.maxOpen, tt.args.factory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientPool_Acquire(t *testing.T) {
	type fields struct {
		Mutex   sync.Mutex
		pool    chan *UserClient
		maxOpen int
		numOpen int
		minOpen int
		closed  bool
		factory ClientFactory
	}
	tests := []struct {
		name    string
		fields  fields
		want    *UserClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ClientPool{
				Mutex:   tt.fields.Mutex,
				pool:    tt.fields.pool,
				maxOpen: tt.fields.maxOpen,
				numOpen: tt.fields.numOpen,
				minOpen: tt.fields.minOpen,
				closed:  tt.fields.closed,
				factory: tt.fields.factory,
			}
			got, err := p.Acquire()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientPool.Acquire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientPool.Acquire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientPool_getOrCreate(t *testing.T) {
	type fields struct {
		Mutex   sync.Mutex
		pool    chan *UserClient
		maxOpen int
		numOpen int
		minOpen int
		closed  bool
		factory ClientFactory
	}
	tests := []struct {
		name    string
		fields  fields
		want    *UserClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ClientPool{
				Mutex:   tt.fields.Mutex,
				pool:    tt.fields.pool,
				maxOpen: tt.fields.maxOpen,
				numOpen: tt.fields.numOpen,
				minOpen: tt.fields.minOpen,
				closed:  tt.fields.closed,
				factory: tt.fields.factory,
			}
			got, err := p.getOrCreate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientPool.getOrCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientPool.getOrCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientPool_Release(t *testing.T) {
	type fields struct {
		Mutex   sync.Mutex
		pool    chan *UserClient
		maxOpen int
		numOpen int
		minOpen int
		closed  bool
		factory ClientFactory
	}
	type args struct {
		client *UserClient
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
			p := &ClientPool{
				Mutex:   tt.fields.Mutex,
				pool:    tt.fields.pool,
				maxOpen: tt.fields.maxOpen,
				numOpen: tt.fields.numOpen,
				minOpen: tt.fields.minOpen,
				closed:  tt.fields.closed,
				factory: tt.fields.factory,
			}
			if err := p.Release(tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("ClientPool.Release() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClientPool_Close(t *testing.T) {
	type fields struct {
		Mutex   sync.Mutex
		pool    chan *UserClient
		maxOpen int
		numOpen int
		minOpen int
		closed  bool
		factory ClientFactory
	}
	type args struct {
		client *UserClient
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
			p := &ClientPool{
				Mutex:   tt.fields.Mutex,
				pool:    tt.fields.pool,
				maxOpen: tt.fields.maxOpen,
				numOpen: tt.fields.numOpen,
				minOpen: tt.fields.minOpen,
				closed:  tt.fields.closed,
				factory: tt.fields.factory,
			}
			if err := p.Close(tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("ClientPool.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClientPool_Shutdown(t *testing.T) {
	type fields struct {
		Mutex   sync.Mutex
		pool    chan *UserClient
		maxOpen int
		numOpen int
		minOpen int
		closed  bool
		factory ClientFactory
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
			p := &ClientPool{
				Mutex:   tt.fields.Mutex,
				pool:    tt.fields.pool,
				maxOpen: tt.fields.maxOpen,
				numOpen: tt.fields.numOpen,
				minOpen: tt.fields.minOpen,
				closed:  tt.fields.closed,
				factory: tt.fields.factory,
			}
			if err := p.Shutdown(); (err != nil) != tt.wantErr {
				t.Errorf("ClientPool.Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
