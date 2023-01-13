package route

import (
	"github.com/MicBun/TestJavan/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	familyRoute := r.Group("/families")
	familyRoute.POST("", controllers.CreateFamily)
	familyRoute.GET("", controllers.GetFamilies)
	familyRoute.PUT("/:id", controllers.UpdateFamily)
	familyRoute.DELETE("/:id", controllers.DeleteFamily)
	familyRoute.GET("/totalPriceAssetByApi", controllers.TotalPriceAssetByApi)

	assetRoute := r.Group("/assets")
	assetRoute.POST("", controllers.CreateAsset)
	assetRoute.GET("", controllers.GetAssets)
	assetRoute.PUT("/:id", controllers.UpdateAsset)
	assetRoute.DELETE("/:id", controllers.DeleteAsset)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
