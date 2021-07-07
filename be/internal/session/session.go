package session

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Session struct {
	id string
	kv map[string]string

	ctx   context.Context
	store *redis.Client
}

func (s *Session) Set(k string, v string) {
	_, err := s.store.HSet(s.ctx, s.id, k, v).Result()
	if err != nil {
		panic(err)
	}
	s.kv[k] = v
}

func (s *Session) Get(k string) string {
	v, err := s.store.HGet(s.ctx, s.id, k).Result()
	if err != nil {
		panic(err)
	}
	return v
}

type SessionManager struct {
	cookieName  string
	maxLifetime int

	ctx context.Context
	rdb *redis.Client
}

var globalSM *SessionManager

func InitSessionManager(cache string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cache,
		Password: "",
		DB:       0,
	})

	globalSM = &SessionManager{
		cookieName:  "SID",
		maxLifetime: 99999,

		ctx: context.Background(),
		rdb: rdb,
	}
}

func GlobalManager() *SessionManager {
	return globalSM
}

func (manager *SessionManager) sessionId(seed string) string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	d := md5.Sum([]byte(seed))
	copy(b[16:], d[:])
	return base64.URLEncoding.EncodeToString(b)
}

func (m *SessionManager) createSession(seed string) *Session {
	sid := m.sessionId(seed)
	kv := make(map[string]string)
	kv["create"] = strconv.FormatInt(time.Now().Unix(), 10)

	_, err := m.rdb.HSet(m.ctx, sid, kv).Result()
	if err != nil {
		panic(err)
	}
	return &Session{
		id: sid,
		kv: kv,

		ctx:   m.ctx,
		store: m.rdb,
	}
}

func (m *SessionManager) getExistSession(sid string) *Session {
	exist, err := m.rdb.Exists(m.ctx, sid).Result()
	if err != nil {
		panic(err)
	}
	if exist == 0 {
		return nil
	}

	kv, err := m.rdb.HGetAll(m.ctx, sid).Result()
	if err != nil {
		panic(err)
	}
	return &Session{
		id: sid,
		kv: kv,

		ctx:   m.ctx,
		store: m.rdb,
	}
}

func (m *SessionManager) deleteSession(sid string) {
	m.rdb.HDel(m.ctx, sid)
}

func (m *SessionManager) StartSession(w http.ResponseWriter, seed string) *Session {
	session := m.createSession(seed)
	sid := session.id
	cookie := http.Cookie{Name: m.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(m.maxLifetime)}
	http.SetCookie(w, &cookie)
	return session
}

func (m *SessionManager) QuerySession(r *http.Request) *Session {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return nil
	}

	sid, _ := url.QueryUnescape(cookie.Value)
	return m.getExistSession(sid)
}

func (m *SessionManager) EndSession(s *Session) {
	m.deleteSession(s.id)
}
