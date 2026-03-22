package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nitesh-mhatre/go-rest-api/db"
	"github.com/nitesh-mhatre/go-rest-api/routes"

)

func main() {
	if err := db.InitDB(); err != nil {
		panic(err)
	}
	r := gin.Default()
	routes.ResgisterRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}



