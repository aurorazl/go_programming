package service

import (
	"encoding/json"
	"fmt"
	"go_programming/chatRoom/common/message"
	"go_programming/chatRoom/common/utils"
)

type SmsProcess struct {
	UserId int
	Tf     *utils.Transfer
}

func (s *SmsProcess) SendGroupSms(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.UserId = s.UserId
	smsMes.UserStatus = message.UserOnline
	smsMes.Content = content
	data, err := json.Marshal(smsMes)
	if err != nil {
		return err
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return err
	}
	err = s.Tf.WritePkg(data)
	return
}

func (s *SmsProcess) recvGroupMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return
	}
	fmt.Printf("recv from %d group message: %s\n", smsMes.UserId, smsMes.Content)
	return
}
