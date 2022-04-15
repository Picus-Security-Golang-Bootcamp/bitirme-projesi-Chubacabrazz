package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Chubacabrazz/picus-storeApp/pkg/config"
	db "github.com/Chubacabrazz/picus-storeApp/pkg/database"

	//graceful "github.com/Chubacabrazz/picus-storeApp/pkg/gracefulExit"
	"github.com/Chubacabrazz/picus-storeApp/pkg/logger"
	"github.com/Chubacabrazz/picus-storeApp/storage/handlers"
	"github.com/Chubacabrazz/picus-storeApp/storage/repo"
	"github.com/Chubacabrazz/picus-storeApp/storage/user"
	"github.com/gin-gonic/gin"
)

func main() {

	log.Println("PicusStore Basket Service starting...")
	//Setting envs.
	cfg, err := config.LoadConfig("./pkg/config/local-cfg")

	if err != nil {
		log.Fatalf("Load Config failed %v", err)
	}

	logger.NewLogger(cfg)
	defer logger.Close()

	DB := db.Connect(cfg)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	/* r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})) */
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int(time.Second)),
	}

	// Router group
	rootRouter := r.Group(cfg.ServerConfig.RoutePrefix)
	//orderRouter := rootRouter.Group("/order")
	//cartRouter := rootRouter.Group("/cart")
	categoryRouter := rootRouter.Group("/category")
	//productRouter := rootRouter.Group("/product")
	userRouter := rootRouter.Group("/user")

	// Category Repository
	CategoryRepo := repo.NewCategoryRepository(DB)
	CategoryRepo.Migration()
	//CategoryRepo.InsertData()
	handlers.CategoryHandler(categoryRouter, CategoryRepo)
	// Product Repository
	ProductRepo := repo.NewProductRepository(DB)
	ProductRepo.Migration()
	// User Repository
	UserRepo := user.NewUserRepository(DB)
	UserRepo.Migration()
	user.NewUserHandler(userRouter, UserRepo)
	// Order Repository
	OrderDetailsRepo := repo.NewOrderDetailsRepository(DB)
	OrderDetailsRepo.Migration()
	// Cart Repository
	CartRepo := repo.NewCartRepository(DB)
	CartRepo.Migration()
	/* SessionRepo := repo.NewSession(DB)
	SessionRepo.Migration() */

	/* go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}() */
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	/* log.Println("Book store service started")
	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int(time.Second))) */
}
