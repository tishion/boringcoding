package session

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestSession_Set(t *testing.T) {
	type fields struct {
		id    string
		kv    map[string]string
		ctx   context.Context
		store *redis.Client
	}
	type args struct {
		k string
		v string
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
			s := &Session{
				id:    tt.fields.id,
				kv:    tt.fields.kv,
				ctx:   tt.fields.ctx,
				store: tt.fields.store,
			}
			s.Set(tt.args.k, tt.args.v)
		})
	}
}

func TestSession_Get(t *testing.T) {
	type fields struct {
		id    string
		kv    map[string]string
		ctx   context.Context
		store *redis.Client
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				id:    tt.fields.id,
				kv:    tt.fields.kv,
				ctx:   tt.fields.ctx,
				store: tt.fields.store,
			}
			if got := s.Get(tt.args.k); got != tt.want {
				t.Errorf("Session.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitSessionManager(t *testing.T) {
	type args struct {
		cache string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitSessionManager(tt.args.cache)
		})
	}
}

func TestGlobalManager(t *testing.T) {
	tests := []struct {
		name string
		want *SessionManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GlobalManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GlobalManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_sessionId(t *testing.T) {
	type fields struct {
		cookieName  string
		maxLifetime int
		ctx         context.Context
		rdb         *redis.Client
	}
	type args struct {
		seed string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := &SessionManager{
				cookieName:  tt.fields.cookieName,
				maxLifetime: tt.fields.maxLifetime,
				ctx:         tt.fields.ctx,
				rdb:         tt.fields.rdb,
			}
			if got := manager.sessionId(tt.args.seed); got != tt.want {
				t.Errorf("SessionManager.sessionId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_createSession(t *testing.T) {
	type fields struct {
		cookieName  string
		maxLifetime int
		ctx         context.Context
		rdb         *redis.Client
	}
	type args struct {
		seed string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SessionManager{
				cookieName:  tt.fields.cookieName,
				maxLifetime: tt.fields.maxLifetime,
				ctx:         tt.fields.ctx,
				rdb:         tt.fields.rdb,
			}
			if got := m.createSession(tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionManager.createSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_getExistSession(t *testing.T) {
	type fields struct {
		cookieName  string
		maxLifetime int
		ctx         context.Context
		rdb         *redis.Client
	}
	type args struct {
		sid string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SessionManager{
				cookieName:  tt.fields.cookieName,
				maxLifetime: tt.fields.maxLifetime,
				ctx:         tt.fields.ctx,
				rdb:         tt.fields.rdb,
			}
			if got := m.getExistSession(tt.args.sid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionManager.getExistSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_deleteSession(t *testing.T) {
	type fields struct {
		cookieName  string
		maxLifetime int
		ctx         context.Context
		rdb         *redis.Client
	}
	type args struct {
		sid string
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
			m := &SessionManager{
				cookieName:  tt.fields.cookieName,
				maxLifetime: tt.fields.maxLifetime,
				ctx:         tt.fields.ctx,
				rdb:         tt.fields.rdb,
			}
			m.deleteSession(tt.args.sid)
		})
	}
}

func TestSessionManager_StartSession(t *testing.T) {
	type fields struct {
		cookieName  string
		maxLifetime int
		ctx         context.Context
		rdb         *redis.Client
	}
	type args struct {
		w    http.ResponseWriter
		seed string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SessionManager{
				cookieName:  tt.fields.cookieName,
				maxLifetime: tt.fields.maxLifetime,
				ctx:         tt.fields.ctx,
				rdb:         tt.fields.rdb,
			}
			if got := m.StartSession(tt.args.w, tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionManager.StartSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_QuerySession(t *testing.T) {
	type fields struct {
		cookieName  string
		maxLifetime int
		ctx         context.Context
		rdb         *redis.Client
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SessionManager{
				cookieName:  tt.fields.cookieName,
				maxLifetime: tt.fields.maxLifetime,
				ctx:         tt.fields.ctx,
				rdb:         tt.fields.rdb,
			}
			if got := m.QuerySession(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionManager.QuerySession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_EndSession(t *testing.T) {
	type fields struct {
		cookieName  string
		maxLifetime int
		ctx         context.Context
		rdb         *redis.Client
	}
	type args struct {
		s *Session
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
			m := &SessionManager{
				cookieName:  tt.fields.cookieName,
				maxLifetime: tt.fields.maxLifetime,
				ctx:         tt.fields.ctx,
				rdb:         tt.fields.rdb,
			}
			m.EndSession(tt.args.s)
		})
	}
}
