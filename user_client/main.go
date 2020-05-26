package main

import (
	restful "github.com/emicklei/go-restful/v3"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/zjjt/biblio/user_client/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("biblio.client.user"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	api := new(handler.API)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	wc.EnableContentEncoding(true)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/user")
	ws.Route(ws.POST("/signup").Consumes("application/x-www-form-urlencoded").To(api.Signup))
	ws.Route(ws.POST("/login").Consumes("application/x-www-form-urlencoded").To(api.Login))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
