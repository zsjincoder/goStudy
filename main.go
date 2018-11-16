package main

import (
	db "./database"
	router "./routers"
)

func main() {
	defer db.Orm.Clone()
	routers := router.InitRouter()
	routers.Run(":8081")
}
