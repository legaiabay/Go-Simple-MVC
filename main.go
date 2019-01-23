package main

//remote : https://github.com/legaiabay/go-sandbox.git
import (
	"gotest.com/go-sandbox/config"

	"github.com/kataras/iris"
)

func main() {
	config.DBConnect()

	app := iris.Default()
	DefineRoutes(app)
	app.Run(iris.Addr(":8787"))
}
