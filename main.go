package main

import (
	coon "./database"
	router "./routers"
)

func main() {
	defer coon.Orm.Clone()
	defer coon.Db.Close()
	routers := router.InitRouter()
	routers.Run(":8081")
}
