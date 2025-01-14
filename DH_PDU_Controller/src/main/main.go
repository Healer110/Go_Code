package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"zw.com/DH_PDU/src/api"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// func init() {
// 	// 设置使用的最大核数
// 	runtime.GOMAXPROCS(1)
// 	// 手动回收垃圾
// 	runtime.GC()
// }

func menu(socketStatus bool, outputstatus bool) {
	var c_stat, o_stat string
	if socketStatus {
		c_stat = color.GreenString("[Connected]")
	} else {
		c_stat = color.RedString("[Disconnected]")
	}

	if outputstatus {
		o_stat = color.GreenString("[ON]")
	} else {
		o_stat = color.RedString("[OFF]")
	}
	fmt.Println(color.CyanString("1. Connect into device ") + c_stat)
	fmt.Println(color.CyanString("2. ON/OFF device ") + o_stat)
	color.Cyan("3. Get device voltage and current")
	color.Cyan("4. Set device voltage")
	color.Cyan("5. Set device current")
	color.Cyan("6. Set device voltage by step")
	color.Cyan("7. Exit")
}

func showInfo(table *tablewriter.Table, data *[][]string) {
	for _, v := range *data {
		table.Append(v)
	}
	table.Render()
	table.ClearRows()

}

func initTable() (*tablewriter.Table, *[][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	data := make([][]string, 2)
	row1, row2 := make([]string, 3), make([]string, 3)
	row1[0] = "MEASURE"
	row2[0] = "LIMIT"
	data[0] = row1
	data[1] = row2
	return table, &data
}

func updateTableHeadAndLimit(table *tablewriter.Table, data *[][]string, pdu *api.Pdu) {
	table.SetHeader([]string{fmt.Sprintf("%s:%d", pdu.Ip, pdu.Port), "Voltage", "Current"})
	(*data)[1][1] = fmt.Sprintf("%.2f", pdu.VoltageLimit)
	(*data)[1][2] = fmt.Sprintf("%.2f", pdu.CurrentLimit)
}

func checkSocketState(socket *api.PduSocket) bool {
	if !socket.ConnetStatus {
		color.Red(">> Device not connected, please connect firstly.")
		return false
	}
	return true
}

// 将控制台接收到的字符串转换为相应的int
func formatOption(revStr string, option *int) bool {
	revStr = strings.Trim(revStr, " \r\n")
	res, err := strconv.Atoi(revStr)
	if err != nil {
		return false
	}
	*option = res
	return true
}

func main() {
	table, data := initTable()
	var discard string
	var input int
	var str string
	var inputFloat float64
	pduScoket := &api.PduSocket{}
	pdu := &api.Pdu{
		Ip:           "",
		Port:         -1,
		VoltageLimit: 0,
		CurrentLimit: 0,
		IsEA:         false,
		Conn:         pduScoket,
		OutputStatus: false,
	}

	// 判断是否存在文件，保存着上次连接的ip,port信息,尝试连接
	str, input = api.CheckIPPortFile("./auth.cfg")
	if input != -1 {
		res := pduScoket.Connect(str, input)
		if res {
			pdu.Ip = str
			pdu.Port = input
			pdu.InitPara()
			updateTableHeadAndLimit(table, data, pdu)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		menu(pduScoket.ConnetStatus, pdu.OutputStatus)
		fmt.Print(color.YellowString("Input option number: "))
		str, err := reader.ReadString('\n')
		if err != nil {
			color.Red(">> read input error")
			continue
		}
		if !formatOption(str, &input) {
			color.Red(">> invalid option")
			continue
		}
		switch input {
		case 1:
			if pduScoket.ConnetStatus {
				color.Red("Device connected, please disconnect firstly.")
				continue
			}
			fmt.Print("Input IP address: ")
			fmt.Scanln(&str)
			fmt.Print("Input port: ")
			_, err = fmt.Scanln(&input)
			if err != nil {
				color.Red(">> invalid port")
				// 清空缓冲区，防止出现多次打印
				fmt.Scanln(&discard)
				continue
			}
			res := pduScoket.Connect(str, input)
			if res {
				pdu.Ip = str
				pdu.Port = input
				pdu.InitPara()
				updateTableHeadAndLimit(table, data, pdu)
				// 第一次连接成功后保存ip和port到本地文件
				api.SaveIPPoort(str, strconv.Itoa(input), "./auth.cfg")
			}
		case 2:
			if !checkSocketState(pduScoket) {
				continue
			}
			// ON/OFF
			if pdu.OutputStatus {
				res := pdu.Set(api.CMD_OUTPUT_OFF)
				if res {
					pdu.OutputStatus = false
				}
			} else {
				res := pdu.Set(api.CMD_OUTPUT_ON)
				if res {
					pdu.OutputStatus = true
				}
			}
		case 3:
			if !checkSocketState(pduScoket) {
				continue
			}
			if pdu.IsEA {
				measVol, measCur, err := pdu.EAPowerGetVolCur(api.CMD_GET_ARRAY)
				if err == nil {
					(*data)[0][1] = measVol
					(*data)[0][2] = measCur
				}
			} else {

				// 获取设备信息, 主要是当前电压和电流
				vol, err := pdu.Get(api.CMD_MEASURE_VOLTAGE)
				if err == nil {
					(*data)[0][1] = api.FormatResponse(vol)
				} else {
					continue
				}
				cur, err := pdu.Get(api.CMD_MEASURE_CURRENT)
				if err == nil {
					(*data)[0][2] = api.FormatResponse(cur)
				} else {
					continue
				}
			}
			(*data)[1][1] = fmt.Sprintf("%.2f", pdu.VoltageLimit)
			showInfo(table, data)
		case 4:
			if !checkSocketState(pduScoket) {
				continue
			}
			// 设置电压上限, 设置完成后，需要重新获取电压上限并赋值给结构体中的电压上限
			fmt.Print("Input voltage limit: ")
			_, err := fmt.Scanln(&inputFloat)
			if err != nil {
				color.Red(">> invalid voltage")
				// 清空缓冲区，防止出现多次打印
				fmt.Scanln(&discard)
				continue
			} else {
				cmd := fmt.Sprintf(api.CMD_SET_VOLTAGE_LIMIT, inputFloat)
				res := pdu.Set(cmd)
				if res {
					res, err := pdu.Get(api.CMD_GET_VOLTAGE_LIMIT)
					if err == nil {
						res = api.FormatResponse(res)
						inputFloat, _ = strconv.ParseFloat(res, 64)
						pdu.VoltageLimit = inputFloat
						(*data)[1][1] = fmt.Sprintf("%.2f", pdu.VoltageLimit)
					}
				}
			}
		case 5:
			if !checkSocketState(pduScoket) {
				continue
			}
			// 设置电流上限，设置完成后，需要重新获取电流上限并赋值给结构体中的电流上限
			fmt.Print("Input current limit: ")
			_, err := fmt.Scanln(&inputFloat)
			if err != nil {
				color.Red(">> invalid current")
				// 清空缓冲区，防止出现多次打印
				fmt.Scanln(&discard)
				continue
			} else {
				cmd := fmt.Sprintf(api.CMD_SET_CURRENT_LIMIT, inputFloat)
				res := pdu.Set(cmd)
				if res {
					res, err := pdu.Get(api.CMD_GET_CURRENT_LIMIT)
					if err == nil {
						res = api.FormatResponse(res)
						inputFloat, _ = strconv.ParseFloat(res, 64)
						pdu.CurrentLimit = inputFloat
						(*data)[1][2] = fmt.Sprintf("%.2f", pdu.CurrentLimit)
					}
				}
			}
		case 6:
			// 按照指定的电压以及步进等时间间隔自动设置电压值
			fmt.Println("Specify [start stop step time] values separated by spaces")
			fmt.Print("[PS~]$ ")
			str, err := reader.ReadString('\n')
			if err != nil {
				api.PrintErrors(fmt.Sprintf("reader error: %s", err.Error()))
			}
			if strings.Trim(str, " \r\n") == "0" {
				continue
			}
			pdu.SetVoltageByStep(str)
		case 7:
			if !checkSocketState(pduScoket) {
				continue
			}
			pduScoket.Disconnect()
		default:
			color.Red(">> invalid option")
		}
	}
}
