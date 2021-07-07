package serverhandler

import "mmc.com/landingtask/be/cmd/userserver/datastorage"

// Handles the request of getting profile
func GetProfile(args map[string]interface{}, r map[string]interface{}) (int, string) {
	s := datastorage.GetInstance()

	username := args["username"]

	nick, avatar, err := s.GetUserProfile(username.(string))
	if nil != err {
		return 10001, "ResourceNotFound"
	}

	r["nickname"] = nick
	r["avatar"] = avatar

	return 0, ""
}
