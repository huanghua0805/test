package main

import (
	"fmt"
	"github.com/konimarti/opc"
	"golang.org/x/net/websocket"
	"time"
)


var origin = "http://202.120.162.6:80/"
var url = "ws://202.120.162.6:80/MachineWebSocket"

func main() {
	client := opc.NewConnection(
		"PhoenixContact.AX-Server.21", // ProgId
		[]string{"localhost"},         //  OPC servers nodes
		[]string{"V1_MG1.Web_MachineOff_Button",
			"V1_MG1.Web_MachineUnlock_Button",
			"V1_MG1.Web_MachineOn_Button",
			"V1_MG1.Web_St1_Prestop_Forward_Stop",
			"V1_MG1.Web_St4_Prestop_Forward_Stop",
			"V1_MG1.Web_St1_Prestop_Backward_Move",
			"V1_MG1.Web_St4_Prestop_Backward_Move",
			"V2_MG1.opc_iOrderNo",
			"V2_MG1.opc_xRead_NewOrder",
			"V1_MG2.Web_Small_Kuka_Control",

		},                             // slice of OPC tags
	)
	fmt.Printf("OPC Server Connection OK")
	defer client.Close()

	//建立WebSocket连接
	ws, err := websocket.Dial(url, "", origin)
	fmt.Printf("WebScoket Connection OK")
	checkErr(err)
	defer ws.Close()//关闭连接


	for {
		//发送message给服务端
		message := []byte("cdhaw")
		_, err = ws.Write(message)
		checkErr(err)
		fmt.Printf("Send: %s\n", message)


		//从服务端接收消息
		msg := make([]byte,1024)
		m, err := ws.Read(msg)
		checkErr(err)
		fmt.Printf("%s\n",msg[:m])

		switch string(msg[:m]){
		case "V1_Reset_and_Run":
			client.Write("V1_MG1.Web_MachineUnlock_Button", 1)
			time.Sleep(time.Millisecond*1500)
			client.Write("V1_MG1.Web_MachineOn_Button", 1)
			//client.Write("V1_MG1.Web_St1_Prestop_Forward_Stop", 1)
			//client.Write("V1_MG1.Web_St4_Prestop_Forward_Stop", 1)
		default:
			break


		case "V2_Load_WebOrder":
			client.Write("V2_MG1.opc_iOrderNo", 1)
			time.Sleep(time.Millisecond * 1000)
			client.Write("V2_MG1.opc_xRead_NewOrder", 1)

		case "V1_Small_Robot_Grip":
			client.Write("V1_MG2.Web_Small_Kuka_Control", 1)

		case "V1_Dialog_Repair":
			client.Write("V1_MG1.Web_St1_Prestop_Forward_Stop", 0)
			client.Write("V1_MG1.Web_St1_Prestop_Backward_Move", 1)
			client.Write("V1_MG1.Web_MachineUnlock_Button", 1)
			time.Sleep(time.Millisecond *1500)
			client.Write("V1_MG1.Web_MachineOn_Button", 1)

		}

		time.Sleep(time.Millisecond *50)
	}

}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
