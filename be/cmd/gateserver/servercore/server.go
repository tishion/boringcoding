package servercore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// local packages
	"mmc.com/landingtask/be/cmd/gateserver/httprouter"
	"mmc.com/landingtask/be/cmd/gateserver/userclient"
	"mmc.com/landingtask/be/internal/common"
	"mmc.com/landingtask/be/internal/session"
	"mmc.com/landingtask/be/internal/usermessage"
)

// The RPC service connection string
var rpcConn string

// Create a user service RCP client
func userClientFactory() (*userclient.UserClient, error) {
	var e error
	client := userclient.NewClient()
	client.Connect(rpcConn)

	common.LogI("New rpc client connected")
	return client, e
}

// RPC service client pool
var rcpClientPool *userclient.ClientPool

// Gets the RPC service client pool
func getRPCClientPool() *userclient.ClientPool {
	return rcpClientPool
}

// Represents the server object
type Server struct {
	router *httprouter.Router
}

// Creates a server
func CreateServer(rpcAddress string, cache string) *Server {
	rpcConn = rpcAddress
	rcpClientPool = userclient.NewClientPool(128, 256, userClientFactory)

	// init the session manager
	session.InitSessionManager(cache)

	// create the http server router
	r := httprouter.New()
	server := &Server{
		router: r,
	}
	return server
}

// Starts the server
func (s *Server) Start(connString string) {
	// add routers
	s.router.POST("/api/v1/login", login)
	s.router.GET("/api/v1/profile/:username", getProfile)
	s.router.PUT("/api/v1/profile/:username", putProfile)

	// start and fly
	log.Fatal(http.ListenAndServe(connString, s.router))
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	common.LogI(">> POST /api/v1/login, content len %d", r.ContentLength)

	// parse the form parameters
	err := r.ParseForm()
	if nil != err {
		common.LogE("Failed to parse the form data")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// botain the username
	username, ok := r.Form["username"]
	if !ok {
		common.LogE("No user name found")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// obtain the password
	password, ok := r.Form["password"]
	if !ok {
		common.LogE("No password found")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// build the parameters object for RCP invocation
	params := make(map[string]interface{})
	params["username"] = username[0]
	params["password"] = password[0]
	reqMsg := usermessage.CreateRequestMessage("login", params)
	reqBuf, err := json.Marshal(reqMsg)
	if nil != err {
		common.LogE("Failed to marshal the RPC request: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// acquire the RCP client session from the pool
	rpcClient, err := getRPCClientPool().Acquire()
	if nil != err {
		common.LogE("Failed to acquire the rpc client: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer getRPCClientPool().Release(rpcClient)

	// invoke the RPC service
	resBuf, err := rpcClient.RpcInvoke(reqBuf)
	if nil != err {
		common.LogE("Failed to invoke rpc service: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	common.LogD("\t << RPC response: %s", string(resBuf))

	// check the RCP result
	var resMsg usermessage.ServiceResponseMessage
	err = json.Unmarshal(resBuf, &resMsg)
	if nil != err {
		common.LogE("Failed to unmarshal the RPC response: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// login success, create session
	if resMsg.Code == 0 {
		seed := fmt.Sprintf("%s.%s.%d", username[0], r.RemoteAddr, time.Now().UTC().Nanosecond())
		session := session.GlobalManager().StartSession(w, seed)
		session.Set("username", username[0])
	}

	// respond the HTTP request
	w.Write(resBuf)
}

// Handles the acquiring profile request
func getProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	common.LogI(">> GET /api/v1/profile")

	// validate the session
	session := session.GlobalManager().QuerySession(r)
	if nil == session {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	// obtain the username
	username := params.Get("username")
	if username == "" {
		common.LogE("No username found")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// build the parameters object for RCP invocation
	data := make(map[string]interface{})
	data["username"] = username
	reqMsg := usermessage.CreateRequestMessage("getProfile", data)
	reqBuf, err := json.Marshal(reqMsg)
	if nil != err {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// acquire the RCP client session from the pool
	rpcClient, _ := getRPCClientPool().Acquire()
	if nil != err {
		common.LogE("Failed to acquire the rpc client: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer getRPCClientPool().Release(rpcClient)

	// invoke the RPC service
	resBuf, err := rpcClient.RpcInvoke(reqBuf)
	if nil != err {
		common.LogE("Failed to invoke rpc service: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	common.LogD("\t << RPC response: %s", string(resBuf))

	// respond the HTTP request
	w.Write(resBuf)
}

// Handles the updating profile request
func putProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	common.LogI(">> PUT /api/v1/profile, content length %d", r.ContentLength)

	// validate the session
	session := session.GlobalManager().QuerySession(r)
	if nil == session {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	// obtain the username
	username := params.Get("username")
	if username == "" {
		common.LogE("No username found")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// build the parameters object for RCP invocation
	data := make(map[string]interface{})
	data["username"] = username
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	reqMsg := usermessage.CreateRequestMessage("updateProfile", data)
	reqBuf, err := json.Marshal(reqMsg)
	if nil != err {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// acquire the RCP client session from the pool
	rpcClient, _ := getRPCClientPool().Acquire()
	if nil != err {
		common.LogE("Failed to acquire the rpc client: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer getRPCClientPool().Release(rpcClient)

	// invoke the RPC service
	resBuf, err := rpcClient.RpcInvoke(reqBuf)
	if nil != err {
		common.LogE("Failed to invoke rpc service: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	common.LogD("\t << RPC response: %s", string(resBuf))

	// respond the HTTP request
	w.Write(resBuf)
}
