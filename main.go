/*
 * @Author: YangQi
 * @Date: 2020-08-23 22:34:53
 * @LastEditors: YangQi
 * @LastEditTime: 2020-09-03 11:06:09
 */
package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main()  {
	model.InitDb()
	routes.InitRouter()
}