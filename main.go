package main

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moura1001/time-manager/src/pkg/handlers"
	"github.com/moura1001/time-manager/src/pkg/logger"
	"github.com/moura1001/time-manager/src/pkg/util"
)

func main() {
	app, err := initApp()
	if err != nil {
		log.Fatal(err)
	}
	logger.Init()

	port := util.GetEnv("LISTEN_PORT", ":3000")
	logger.Log("msg", fmt.Sprintf("Server is listening on port %s...", port))
	log.Fatal(app.Run(port))
}

func initApp() (*gin.Engine, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.New()
	router.Static("/src/static/assets", "./src/static/assets")
	router.Static("/dist", "./dist")
	router.HTMLRender = createEngine()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		ctx.Set("Pragma", "no-cache")
		ctx.Set("Expires", "0")
		ctx.Set("Surrogate-Control", "no-store")
		ctx.Next()
	})
	router.Use(handlers.DefaultErrorHandler())

	router.GET("/", handlers.HandleGetHome)
	router.POST("/increase", handlers.HandleAddTiming)

	return router, nil
}

func createEngine() *ginview.ViewEngine {
	engine := goview.New(goview.Config{
		Root:      "src/static/views",
		Extension: ".html",
		Funcs: template.FuncMap{
			"css": func(name string) (res template.HTML) {
				filepath.Walk("./src/static/assets", func(path string, info fs.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.Name() == name {
						res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
					}
					return nil
				})
				return
			},
		},
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})

	return ginview.Wrap(engine)
}
