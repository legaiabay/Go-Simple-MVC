package main

//remote : https://github.com/legaiabay/go-sandbox.git
import (
	"strconv"

	"gotest.com/go-sandbox/config"

	"github.com/kataras/iris"
)

func startNewIrisServer(port int) {
	config.DBConnect()
	app := iris.Default()
	DefineRoutes(app)
	app.Run(iris.Addr(":" + strconv.Itoa(port)))
}

func main() {
	// Testing webserver with external routes (Iris)
	startNewIrisServer(8787)
}
