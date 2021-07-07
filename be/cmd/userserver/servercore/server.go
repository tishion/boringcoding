package servercore

import (
	"net"
	"sync"

	// local package
	"mmc.com/landingtask/be/cmd/userserver/datastorage"
	"mmc.com/landingtask/be/internal/common"
)

// Represents the server
type Server struct {
	listener net.Listener
	clients  sync.Map

	pending chan net.Conn
	leaving chan net.Conn
}

// Creates a Server object
func CreateServer(db string) *Server {
	datastorage.InitDataStore(db)

	server := &Server{
		pending: make(chan net.Conn),
		leaving: make(chan net.Conn),
	}
	return server
}

// Starts the server
func (s *Server) Start(connString string) {
	s.prepareChannels()
	s.listenAndServe(connString)
}

// Stops the server
func (s *Server) Stop() {
	s.listener.Close()
}

// Prepares the channels
func (s *Server) prepareChannels() {
	go func() {
		for {
			select {
			case conn := <-s.pending:
				s.createSession(conn)
			case conn := <-s.leaving:
				s.removeSession(conn)
			}
		}
	}()
}

// Starts to listen on the server socket
func (s *Server) listenAndServe(connString string) {
	s.listener, _ = net.Listen("tcp", connString)
	common.LogI("Server %p is running on: %s\n", s, connString)

	for {
		conn, err := s.listener.Accept()

		if err != nil {
			common.LogE("Failed to accept the connection:", err.Error())
			return
		}

		common.LogI("A new connection was accepted:%s\n", conn.RemoteAddr().String())
		s.pending <- conn
	}
}

// Creates a client session
func (s *Server) createSession(conn net.Conn) {
	clientSession := NewClientSession(conn)
	clientSession.setEndHandler(func(c net.Conn) {
		s.removeSession(c)
	})
	s.clients.Store(conn, clientSession)

	common.LogI("Session created for client: %s", conn.RemoteAddr().String())

	go func() {
		clientSession.start()
	}()
}

// Removes the client session
func (s *Server) removeSession(conn net.Conn) {
	if conn != nil {
		common.LogI("Session removed for client: %s", conn.RemoteAddr().String())
		conn.Close()
		s.clients.Delete(conn)
	}
}
