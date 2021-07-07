package userclient

import (
	"net"
	"time"

	// local packages
	"mmc.com/landingtask/be/internal/common"
	"mmc.com/landingtask/be/internal/lightrpc"
)

// Represents the user service RPC client
type UserClient struct {
	conn net.Conn
	rpc  lightrpc.Conn
}

// Creates a user service RPC client
func NewClient() *UserClient {
	return &UserClient{}
}

// Connects to the user service RPC server
func (s *UserClient) Connect(conString string) bool {
	c, err := net.DialTimeout("tcp", conString, 3*time.Second)
	if nil != err {
		common.LogE("Failed to connect to server:%s", err.Error())
		return false
	}
	s.conn = c
	s.rpc = lightrpc.LightRpcConn(c)
	return true
}

// Closes the RPC connections
func (s *UserClient) Close() error {
	return s.conn.Close()
}

// Performs the RPC client invocation
func (s *UserClient) RpcInvoke(req []byte) ([]byte, error) {
	// rpc.send
	err := s.rpc.Send(req)
	if nil != err {
		return nil, err
	}

	// rpc.receive
	return s.rpc.Receive()
}
