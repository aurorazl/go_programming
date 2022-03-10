package controller

import (
	"fmt"
	"go_programming/chatRoom/common/message"
	"go_programming/chatRoom/common/utils"
	"go_programming/chatRoom/server/service"
	"io"
	"log"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (p *Processor) ProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		userProcess := service.UserProcess{
			Conn: p.Conn,
		}
		userProcess.ProcessLogin(mes)
	case message.RegisterMesType:
		userProcess := service.UserProcess{
			Conn: p.Conn,
		}
		userProcess.ProcessRegister(mes)
	case message.SmsMesType:
		process := &service.SmsProcess{}
		process.SendGroupSms(mes)
	default:
		fmt.Println("mes type error")
	}
	return
}

func (p *Processor) Process() {
	fmt.Println(p.Conn.RemoteAddr().String() + "connected!")
	defer p.Conn.Close()
	tf := utils.Transfer{Conn: p.Conn}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit")
				return
			}
			fmt.Println("read mes error", err)
			return
		}
		fmt.Println(mes)
		err = p.ProcessMes(&mes)
		if err != nil {
			log.Fatal(err)
		}
	}
}
