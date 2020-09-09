/*
 * @Author: YangQi
 * @Date: 2020-09-03 09:52:00
 * @LastEditors: YangQi
 * @LastEditTime: 2020-09-03 11:04:13
 */
package routes

import(
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//用户模块的路由接口
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		router.PUT("user/:id",v1.EditUser)
		router.DELETE("user/:id",v1.DeleteUser)
		//分类模块的路由接口

		//文章模块的路由接口
	}

	r.Run(utils.HttpPort)
}