package main

import (
	"crackpassword/dao"
	"crackpassword/router"
)

func main() {
	dao.InitDb()
	router.InitRouter()
}
