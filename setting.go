package main

import (
	"github.com/astaxie/beego"

	"devops-api/common"
	_ "devops-api/routers"
)

// 初始化一些配置
func setting() {

	// 初始化获取命令行参数
	common.InitCli()

	// 是否启用 定时生成验证密码 功能
	if ok, _ := beego.AppConfig.Bool("authpassword::enableCrontabAuthPassword"); ok {
		common.CronGenAuthPassword()
	}

	// 是否启用 定时清除验证密码 功能
	if ok, _ := beego.AppConfig.Bool("authpassword::enableManualGenAuthPassword"); ok {
		common.CronClearAuthPassword()
	}

}
