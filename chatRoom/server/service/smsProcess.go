package service

import (
	"encoding/json"
	"fmt"
	"go_programming/chatRoom/common/message"
)

type SmsProcess struct {

}

func (s *SmsProcess) SendGroupSms(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return err
	}
	for userId,userProcess := range userMgr.getOnlineUsers() {
		if userId == smsMes.UserId {
			continue
		}
		// 转发
		err = userProcess.SendMes(mes)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return err
}