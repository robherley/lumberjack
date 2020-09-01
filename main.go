package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	docs "github.com/robherley/lumberjack/docs"
	"github.com/robherley/lumberjack/es"
	"github.com/robherley/lumberjack/routes"
	"github.com/robherley/lumberjack/util"
)

// @title Lumberjack
// @version 1.0
// @description This experimental API interacts with an elasticsearch instance to return logs over websocket
// @contact.name Rob Herley
// @contact.email robert.herley@ibm.com
// @BasePath /
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
func main() {
	var host string
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.Schemes = []string{"https"}
		host = "0.0.0.0"
	} else {
		gin.SetMode(gin.DebugMode)
		docs.SwaggerInfo.Schemes = []string{"http"}
		host = "localhost"
	}

	// connect to elastic
	es6, err := es.GetClient()
	if err != nil {
		util.Log.Fatalln("unable to connect to elasticsearch:", err.Error())
	}

	_, err = es6.Ping()
	if err != nil {
		util.Log.Fatalln(err.Error())
	}
	util.Log.Infoln("Successfully Connected to Elastic Search")

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.LoadHTMLGlob("html/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/healthz", routes.Healthz)
	v1 := r.Group("/api/v1")
	{
		logs := v1.Group("/logs")
		{
			logs.GET("/:index/:project", routes.GetProjectLogs)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, map[string]interface{}{
			"code":    404,
			"error":   "Request Not Found",
			"message": "docs: http://" + c.Request.Host + "/swagger/index.html",
		})
	})

	addr := net.JoinHostPort(host, os.Getenv("PORT"))
	r.Run(addr)
}
