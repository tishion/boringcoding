package usermessage

import (
	"strconv"
	"time"

	"mmc.com/landingtask/be/internal/common"
)

type ServiceRequestMessage struct {
	Name      string
	RequestId string
	Timestamp string
	Data      map[string]interface{}
}

func CreateRequestMessage(name string, params map[string]interface{}) ServiceRequestMessage {
	return ServiceRequestMessage{
		Name:      name,
		RequestId: common.NewUuid(),
		Timestamp: strconv.FormatInt(time.Now().UnixNano(), 10),
		Data:      params,
	}
}

type ServiceResponseMessage struct {
	Code  int
	Error string
	Data  map[string]interface{}
}
