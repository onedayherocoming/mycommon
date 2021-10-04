package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
	"strconv"
)



//设置配置中心
func GetConsulConfig(host string,port int64,prefix string)(config.Config,error){
	consulSource := consul.NewSource(
		//设置配置中心地址
		consul.WithAddress(host+":"+strconv.FormatInt(port,10)),
		//设置前缀，不设置默认前缀 /micro/config
		consul.WithPrefix(prefix),
		//设置是否不带前缀
		consul.StripPrefix(true),
		)
	config,err := config.NewConfig()
	if err !=nil{
		return config,err
	}
	//加载配置
	err = config.Load(consulSource)
	return config,err
}
