package model


import(
    "ServerScan/models"
)



/*
1. 主要用于运行插件 - 主机漏洞检测
*/


var Plugins map[string]Need



func init() {
    /* 设定容器的容量 */
    Plugins = make(map[string]Need)
}




// 设定插在此容器的插件的样子
type Need interface {
    /* 它必须告诉容器它是否生病 */
    Flag() string
    /* 它必须得有启动方法 */
    Start(models.Service)
}



// 启动这个容器中所有的插件
func Start(IP models.Service) {
    for _, plugin := range Plugins {
        /* 查看插件是否是启用状态 */
        // plugin.Flag() 这里的判断可以改成判断运行的模块
        if IP.Protocol == plugin.Flag() {
            // 这里之前是go 并发的
            plugin.Start(IP)
            //fmt.Printf("加载 %s\n", name)
        } 
    }
}



// 插件做完之后必须得插入到容器中
func Regist(name string, plugin Need) {
    Plugins[name] = plugin
}