package routers

import (
	_ "blog-service/docs"
	"blog-service/global"
	"blog-service/internal/middleware"
	v1 "blog-service/internal/routers/api/v1"
	"blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")

	//
	//配置中间件
	//
	//限流中间件
	var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})
	r.Use(middleware.RateLimiter(methodLimiters))
	//超时中间件
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	//为api注册中间件：鉴权
	apiv1.Use(middleware.JWT())
	//为api注册中间件：日志
	apiv1.Use(middleware.AccessLog())


	//
	//配置路由
	//
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := v1.NewUpload()
	//上传文件接口的路由
	r.POST("/upload/file", upload.UploadFile)
	//设置文件服务去提供静态资源的访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	//获取JWT Token
	r.POST("/auth", v1.GetAuth)
	//设置api的路由
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}


