package swagger

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

//go:embed swagger-ui
var swaggerUI embed.FS

func Serve(r *gin.Engine) {
	// Serve Swagger UI from embedded files
	subFS, err := fs.Sub(swaggerUI, "swagger-ui")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/swagger", http.FS(subFS))

	// Serve swagger.json and swagger.yaml from the generated docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
