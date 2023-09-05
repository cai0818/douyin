package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"os"
	"simpledouyin/Dao/vo"
	"simpledouyin/model"
	"simpledouyin/router"
)

func Init(arg string) {
	dsn := "root:123456789mk@tcp(192.168.1.176:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		println(err)
	}
	model.DB = db
	model.Arg = arg
}

var (
	Arg1 string
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	args := os.Args[1:]
	os.MkdirAll("./static/picture/", os.ModePerm)
	r := gin.Default()
	Init(args[0])

	router.InitRouter(r)
	var db = model.DB
	db.AutoMigrate(&vo.VideoList{}, &vo.Favorite{}, &vo.IsFollower{}, &vo.Comment{}, &vo.Message{})
	r.Run()
	fmt.Println(getClientIp())
}
func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}

	return "", errors.New("can not find the client ip address")

}
