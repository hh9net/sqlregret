package client

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/siddontang/go-log/log"
)

func Slave1() {
	h, _ := log.NewStreamHandler(os.Stdout)
	logger := log.NewDefault(h)
	conn := NewMysqlConnection("192.168.0.135", "logreader", "123456", "test", 3306, 123, logger)

	if err := conn.Connect(); nil != err {
		fmt.Println("connect failed:", err.Error())
	}

	if err := conn.Register(); nil != err {
		fmt.Println("register failed:", err.Error())
	}

	fmt.Println("begin dump")
	conn.Dump(107, "mysql-bin.000001", EntryCallback)
	fmt.Println("end dump")
}

func Slave2() {
	h, _ := log.NewStreamHandler(os.Stdout)
	logger := log.NewDefault(h)
	conn := NewMysqlConnection("192.168.0.135", "logreader", "123456", "test", 3306, 122, logger)

	if err := conn.Connect(); nil != err {
		fmt.Println("connect failed:", err.Error())
	}

	if err := conn.Register(); nil != err {
		fmt.Println("register failed:", err.Error())
	}

	fmt.Println("begin dump")
	conn.Dump(107, "mysql-bin.000001", EntryCallback)
	fmt.Println("end dump")
}

func TestMysqlConn(t *testing.T) {
	//go Slave1()
	go Slave2()

	for {
		time.Sleep(time.Second)
	}
}
