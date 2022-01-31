package database

//以下是配置Redis数据库
var (
	addrRedis = "180.184.74.143:6379" //redis地址
	pwd       = ""                    //redis密码
	dbnum     = 0                     //redis数据库编号
)

//以下是配置mysql数据库
var (
	addrMYSQL = "127.0.0.1:3306" //mysql地址
	account   = "root"           //mysql账号
	password  = "root"           //mysql密码
	dbName    = "bytedancecamp"  //mysql数据库
)
