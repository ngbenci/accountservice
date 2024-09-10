package service

import "net/http"

type Route struct {
	Name        string           // 路由名称
	Method      string           //请求方法
	Pattern     string           //匹配模式
	HandlerFunc http.HandlerFunc //处理器
}

type Routes []Route //不用初始化, Routes不是nil，实际上这里没有必要这样定义类型别名

var routes = Routes{

	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		GetAccount,
	},

	Route{
		"hello",
		"GET",
		"/",
		Hello,
	},

	Route{
		"test",
		"GET",
		"/test/*",
		Test,
	},
}
