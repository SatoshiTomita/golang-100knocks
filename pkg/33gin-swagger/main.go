package main

import (
	"net/http"

	docs "golang-100knocks/pkg/33gin-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
swaggerのdocsを作成するには以下のコマンドを実行する
swag init .

swaggerのUIを確認するには以下のURLにアクセスする
http://localhost:8080/swagger/index.html
*/

// @Summary ハローワールド
// @Schemes
// @Description 文字列を返すだけ
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @Summary ハローワールド
// @Schemes
// @Description POSTされたら文字列を返すだけ
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [post]
func HelloworldPost(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	r := gin.Default()

	// swaggerの設定
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
			eg.POST("/helloworld", HelloworldPost)
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
