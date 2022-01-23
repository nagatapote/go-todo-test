package route

import (
	"go-gorm-test/interface/controllers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type UserRouter interface {
	UserRouting(e *echo.Echo)
	userAuthRouting(e *echo.Echo)
	userCertificationRouting(e *echo.Echo)
}

type userRouter struct {
	Uc controllers.UserController
}

func NewUserRouter(uc controllers.UserController) UserRouter {
	return userRouter{uc}
}

func (ur userRouter) UserRouting(e *echo.Echo) {
	ur.userAuthRouting(e)
	ur.userCertificationRouting(e)
}

func (ur userRouter) userAuthRouting(e *echo.Echo) {
	r := e.Group("/user")
	r.POST("/login", ur.Uc.UserLogin)
}

func (ur userRouter) userCertificationRouting(e *echo.Echo) {
	r := e.Group("/user")
	r.Use(middleware.JWT([]byte(os.Getenv("SIGNINGKEY"))))
	r.GET("/all", ur.Uc.UserGet)
	r.POST("/create", ur.Uc.UserCreate)
	r.POST("/upload", ur.Uc.UserUpload)
	r.PUT("/update/:id", ur.Uc.UserUpdate)
	r.DELETE("/delete/:id", ur.Uc.UserDelete)
}
