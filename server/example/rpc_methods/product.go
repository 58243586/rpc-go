package rpc_methods

import (
	"fmt"
	"rpc-go/server/config"
	"rpc-go/server/service/register"
	"rpc-go/server/transport"
)

func init() {
	register.RegisterHandler("Example/a/b.sayHello", RpcTest2Handler)
}

func RpcTest2Handler(conn *transport.JumeiConn, request interface{}) (response string, err error) {
	response = fmt.Sprintf("RpcTest2Handler response :%s", request)
	config.Logger.Info("RpcTest2Handler response:", response)

	return
}
