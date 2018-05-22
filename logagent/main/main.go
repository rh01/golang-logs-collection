package main

import "fmt"

func main() {

	// 读取配置
	filename := "./conf/logcollect.conf"
	err := loadConf("ini",filename)
	if err != nil {
		fmt.Println("load conf failed, err: %v", err)
		panic(err)
		return
	}

	//初始化配置
	err = initLogger()
	if err != nil {
		fmt.Println("load conf failed, err: %v", err)
		panic(err)
		return
	}

}
