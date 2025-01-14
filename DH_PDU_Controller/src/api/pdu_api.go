package api

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Pdu struct {
	Ip           string
	Port         int
	VoltageLimit float64    // 电压上限
	CurrentLimit float64    // 电流上限
	IsEA         bool       // 是否为EA厂家的电源
	Conn         *PduSocket // 连接
	OutputStatus bool       // 输出状态 0: 电源输出关闭, 1: 电源输出打开
}

const (
	CMD_SET_VOLTAGE_LIMIT = "VOLT %.3f"      //设置允许电压的上限值
	CMD_GET_VOLTAGE_LIMIT = "VOLT?"          //读取允许电压的上限值
	CMD_SET_CURRENT_LIMIT = "CURR %.3f"      //设置允许电流的上限值
	CMD_GET_CURRENT_LIMIT = "CURR?"          //读取允许电流的上限值
	CMD_OUTPUT_ON         = "OUTP ON"        //打开设备输出
	CMD_OUTPUT_OFF        = "OUTP OFF"       //关闭设备输出
	CMD_OUTPUT_STATE      = "OUTP?"          // 0: OFF, 1: ON
	CMD_MEASURE_VOLTAGE   = "MEAS:VOLT?"     //读取实际电压
	CMD_MEASURE_CURRENT   = "MEAS:CURR?"     //读取实际电流
	CMD_IDN               = "*IDN?"          //查询设备的IDN 信息
	CMD_SET_LOCK_STATUS   = "SYST:LOCK %s"   //设置锁定状态, ON: 锁定, OFF: 解锁
	CMD_GET_ARRAY         = "MEASure:ARRay?" //EA电源指令，可以直接获取电压、电流、功率
)

// 初始化结构体参数
func (pdu *Pdu) InitPara() {
	pdu.CheckPDUManufactory()
	res, _ := pdu.Get(CMD_GET_VOLTAGE_LIMIT)
	if res != "" {
		res = FormatResponse(res)
		pdu.VoltageLimit, _ = strconv.ParseFloat(res, 64)
	}
	res, _ = pdu.Get(CMD_GET_CURRENT_LIMIT)
	if res != "" {
		res = FormatResponse(res)
		pdu.CurrentLimit, _ = strconv.ParseFloat(res, 64)
	}
	res, _ = pdu.Get(CMD_OUTPUT_STATE)
	if res != "" {
		if res == "1" || res == "ON" {
			pdu.OutputStatus = true
		} else {
			pdu.OutputStatus = false
		}
	}
}

// 该方法用于获取设备各种参数信息
func (pdu *Pdu) Get(cmd string) (string, error) {
	if pdu.IsEA {
		(*pdu.Conn.Conn).Close()
		pdu.Conn.Connect(pdu.Ip, pdu.Port)
	}
	res := pdu.Conn.Send(cmd)
	if res {
		reply, err := pdu.Conn.Receive()
		if err != nil {
			PrintErrors(err.Error())
		} else {
			return strings.Trim(reply, "\n"), nil
		}
	}
	return "", fmt.Errorf("failed to get %s", cmd)
}

// 该方法用于设置设备参数, 一般设置设备参数没有返回值
func (pdu *Pdu) Set(cmd string) bool {
	if pdu.IsEA {
		(*pdu.Conn.Conn).Close()
		pdu.Conn.Connect(pdu.Ip, pdu.Port)
	}
	res := pdu.Conn.Send(cmd)
	if !res {
		PrintErrors("Send command fail...")
	}

	return res
}

// 自动执行电压遍历时的设置电压并显示电压的部分
func (pdu *Pdu) traverseVoltage(v float64) {
	res := pdu.Set(fmt.Sprintf(CMD_SET_VOLTAGE_LIMIT, v))
	if res {
		currentVol, err := pdu.Get(CMD_MEASURE_VOLTAGE)
		if err != nil {
			PrintErrors(fmt.Sprintf("traverse voltage errors when get voltage value: %s", err.Error()))
			return
		}
		PrintMsg(fmt.Sprintf("Measure voltage: %s", currentVol))
		pdu.VoltageLimit = math.Round(v*100) / 100
	}
}

// 遍历电压
func (pdu *Pdu) SetVoltageByStep(input string) {
	input = strings.Trim(input, " \n")
	// 按照一个或者多个空格split字符串
	str_lst := strings.Fields(input)
	if len(str_lst) != 4 {
		PrintErrors("The number of parameters is incorrect. Please enter four parameters.")
		return
	}
	start, err := strconv.ParseFloat(str_lst[0], 64)
	if err != nil {
		PrintErrors(fmt.Sprintf("Parse parameters error: %s", err.Error()))
		return
	}

	stop, err := strconv.ParseFloat(str_lst[1], 64)
	if err != nil {
		PrintErrors(fmt.Sprintf("Parse parameters error: %s", err.Error()))
		return
	}
	step, err := strconv.ParseFloat(str_lst[2], 64)
	if err != nil {
		PrintErrors(fmt.Sprintf("Parse parameters error: %s", err.Error()))
		return
	}
	delay, err := strconv.ParseFloat(str_lst[3], 64)
	if err != nil {
		PrintErrors(fmt.Sprintf("Parse parameters error: %s", err.Error()))
		return
	}

	if start <= 0 || stop <= 0 || step <= 0 || delay <= 0 {
		PrintErrors("Please enter the correct value: [start|stop|step|delay>0]")
		return
	}

	if start <= stop {
		for v := start; v <= stop; v += step {
			pdu.traverseVoltage(v)
			time.Sleep(time.Duration(delay) * time.Second)
		}
	} else {
		for v := start; v >= stop; v -= step {
			pdu.traverseVoltage(v)
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}
}

// 核查PDU是否为EA厂家的，如果是EA厂家的，在初始化的时候就需要执行锁定操作
func (pdu *Pdu) CheckPDUManufactory() {
	idn, err := pdu.Get(CMD_IDN)
	if err != nil {
		PrintErrors(err.Error())
		return
	}
	PrintMsg(strings.Trim(idn, " \n"))
	if strings.Contains(idn, "EA") {
		pdu.IsEA = true
		pdu.Set(fmt.Sprintf(CMD_SET_LOCK_STATUS, "ON"))
	}
}

// EA 电源查询电压电流的方法，返回值：电压、电流、Error
func (pdu *Pdu) EAPowerGetVolCur(cmd string) (string, string, error) {
	(*pdu.Conn.Conn).Close()
	pdu.Conn.Connect(pdu.Ip, pdu.Port)
	res, err := pdu.Get(cmd)
	if err == nil {
		res := strings.Split(res, ",")
		vol_str := res[0][:len(res[0])-2]
		cur_str := res[1][:len(res[1])-2]
		return vol_str, cur_str, nil
	}
	return "", "", err
}
