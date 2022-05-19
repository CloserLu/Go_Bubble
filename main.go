package main

import (
	"bubble/dao"
	"bubble/routers"
)

func main() {
	dao.InitDb()
	r := routers.InitRouter()
	r.Run(":9000")
}
