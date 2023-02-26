package main

import (
	"net/http"

	"github.com/flamego/binding"
	"github.com/flamego/cors"
	"github.com/flamego/flamego"

	"github.com/flamego-examples/bilibili-lottery/context"
	"github.com/flamego-examples/bilibili-lottery/embed"
	"github.com/flamego-examples/bilibili-lottery/form"
	"github.com/flamego-examples/bilibili-lottery/route"
)

func main() {
	f := flamego.New()
	f.Use(flamego.Logger())
	f.Use(flamego.Recovery())
	f.Use(cors.CORS(), context.Contexter())

	if flamego.Env() == flamego.EnvTypeProd {
		f.Use(flamego.Static(
			flamego.StaticOptions{
				FileSystem: http.FS(embed.WebAssets()),
			},
		))
	}

	f.Group("/api", func() {
		lotteryHandler := route.NewLotteryHandler()
		f.Post("/lottery", binding.JSON(form.Lottery{}), lotteryHandler.Lottery)
	})

	f.Run()
}
