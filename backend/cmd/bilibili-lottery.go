package main

import (
	"github.com/flamego/binding"
	"github.com/flamego/cors"
	"github.com/flamego/flamego"

	"github.com/flamego-examples/bilibili-lottery/context"
	"github.com/flamego-examples/bilibili-lottery/form"
	"github.com/flamego-examples/bilibili-lottery/route"
)

func main() {
	f := flamego.Classic()
	f.Use(cors.CORS(), context.Contexter())

	f.Group("/api", func() {
		lotteryHandler := route.NewLotteryHandler()
		f.Post("/lottery", binding.JSON(form.Lottery{}), lotteryHandler.Lottery)
	})

	f.Run()
}
