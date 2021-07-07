package serverhandler

import (
	"mmc.com/landingtask/be/cmd/userserver/datastorage"
)

// Handles the request of updating the profile
func UpdateProfile(args map[string]interface{}, r map[string]interface{}) (int, string) {
	s := datastorage.GetInstance()

	username := args["username"]
	nickname := args["nickname"]
	avatar := args["avatar"]

	_, _, err := s.UpdateProfile(username.(string), nickname.(string), avatar.(string))
	if nil != err {
		return 10001, "ResourceNotFound"
	}

	return 0, ""
}
