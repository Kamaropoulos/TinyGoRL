package main

import (
	"os"

	goctapus "github.com/Kamaropoulos/goctapus-mongo"
	"github.com/Kamaropoulos/TinyGoRL/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	goctapus.Init(os.Args, "debug", "goapp")

	goctapus.File("/", "public/index.html")
	goctapus.GET("/:url", handlers.GetURL())
	goctapus.PUT("/urls", handlers.PutURL())
	goctapus.GET("/urls", handlers.GetURLs())
	// goctapus.DELETE("/urls", handlers.DeleteURL())

	goctapus.Start()
}
