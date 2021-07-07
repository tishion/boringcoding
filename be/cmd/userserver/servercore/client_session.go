package servercore

import (
	"encoding/json"
	"net"

	// local packages

	"mmc.com/landingtask/be/cmd/userserver/serverhandler"
	"mmc.com/landingtask/be/internal/common"
	"mmc.com/landingtask/be/internal/lightrpc"
	"mmc.com/landingtask/be/internal/usermessage"
)

// Represents the disconnection handler
type DisconnectHandler func(c net.Conn)

// Represents the client session
type ClientSession struct {
	conn    net.Conn
	rpc     lightrpc.Conn
	handler DisconnectHandler
}

// Creates a ClientSession object
func NewClientSession(c net.Conn) *ClientSession {
	return &ClientSession{
		conn: c,
		rpc:  lightrpc.LightRpcConn(c),
	}
}

// Sets the handler for client session finish
func (s *ClientSession) setEndHandler(h DisconnectHandler) {
	s.handler = h
}

// Starts the session loop
func (s *ClientSession) start() {
	go s.sessionLoop()
}

// Runs the loop to receive the request from client
func (s *ClientSession) sessionLoop() {
	common.LogI("Enter session reading loop...")
	for {
		// receive a message from the remote
		data, err := s.rpc.Receive()
		if nil != err {
			common.LogE("Failed to scan packet, EOF reached or IO Error occurred")
			break
		}

		// process the message
		s.processMessage(data)
	}

	// session end, call the handler
	if nil != s.handler {
		s.handler(s.conn)
	}
}

// Process the client message
func (s *ClientSession) processMessage(msg []byte) {
	common.LogD("message received, len: %d, text: %s", len(msg), string(msg))
	reqMsg := new(usermessage.ServiceRequestMessage)
	err := json.Unmarshal(msg, &reqMsg)
	if nil != err {
		common.LogE("Failed to deserialize the request")
		e := []byte(`{"code":-1, "error":"Failed to serialize the response"}`)
		s.send(e)
		return
	}

	// parse the request message
	serviceName := reqMsg.Name
	errCode := 0
	errMsg := ""
	result := make(map[string]interface{})
	switch {
	case serviceName == "login":
		{
			common.LogI("request: login, data len %d", len(msg))
			errCode, errMsg = serverhandler.Login(reqMsg.Data, result)
		}
	case serviceName == "getProfile":
		{
			common.LogI("request: getProfile, data len %d", len(msg))
			errCode, errMsg = serverhandler.GetProfile(reqMsg.Data, result)
		}
	case serviceName == "updateProfile":
		{
			common.LogI("request: updateProfile, data len %d", len(msg))
			errCode, errMsg = serverhandler.UpdateProfile(reqMsg.Data, result)
		}
	default:
		common.LogE("Unknown request name")
		e := []byte(`{"code":-2, "error":"Unknown request name"}`)
		s.send(e)
		return
	}

	// build the response message
	res := usermessage.ServiceResponseMessage{
		Code:  errCode,
		Error: errMsg,
		Data:  result,
	}

	// serialize the response message
	resBuf, err := json.Marshal(res)
	if nil != err {
		common.LogE("Failed to serialize the response")
		e := []byte(`{"code":-3, "error":"Failed to serialize the response"}`)
		s.send(e)
		return
	}

	// send the response
	s.send(resBuf)
}

// Sends data to remote
func (s *ClientSession) send(data []byte) {
	s.rpc.Send(data)
}
