package enforcer

import (
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func EnforcerTool() *casbin.Enforcer {

	a, err := xormadapter.NewAdapter("mysql", "root:mxymxy@tcp(172.19.128.1)/web?charset=utf8", true)
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
	}
	//创建一个casbin决策器需要一个模板文件和策略文件为参数
	e, err := casbin.NewEnforcer("conf/keymatch.conf", a)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
	}

	e.EnableAutoSave(true)
	//从数据库加载策略
	e.LoadPolicy()
	return e
}
