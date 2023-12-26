package http_server

import "github.com/labstack/echo/v4"

func (app *app) publicRoute(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	userApi := apiV1.Group("/user")
	{
		userApi.GET("/detail", app.userHandler.Detail)
		userApi.PUT("/register", app.userHandler.Register)
		userApi.POST("/update", app.userHandler.Update)
	}

	profileApi := apiV1.Group("/profile")
	{
		profileApi.GET("/:profile_user_id", app.profileHander.Profile)
		profileApi.POST("/:follow_user_id/follow", app.profileHander.Follow)
		profileApi.DELETE("/:unfollow_user_id/unfollow", app.profileHander.Unfollow)
	}

}
