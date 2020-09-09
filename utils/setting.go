/*
 * @Author: YangQi
 * @Date: 2020-08-23 23:21:23
 * @LastEditors: YangQi
 * @LastEditTime: 2020-09-03 11:17:05
 */
package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var(
	AppMode string
	HttpPort string

	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
)

func init(){
	file,err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查文件路径:",err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File){
	AppMode = file.Section("server").Key("AppMode3").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File){
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}