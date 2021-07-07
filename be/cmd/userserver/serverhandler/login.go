package serverhandler

import (
	"mmc.com/landingtask/be/cmd/userserver/datastorage"
)

// Handles the request of logging
func Login(args map[string]interface{}, r map[string]interface{}) (int, string) {
	s := datastorage.GetInstance()

	username := args["username"]

	psw, err := s.GetUserPsw(username.(string))
	if nil != err {
		return 10001, "InvlaidUsernameOrPassword"
	}

	password := args["password"]
	if psw != password {
		return 10002, "InvlaidUsernameOrPassword"
	}

	return 0, ""
}
