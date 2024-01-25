package http_server

func (app *app) routeVer1() {

	apiV1 := app.httpServer.GroupRouter("/api/v1")

	// unauthorized api group
	userApi := apiV1.Group("/user")
	{
		userApi.POST("/login", app.userHandler.Login)
		userApi.PUT("/register", app.userHandler.Register)
	}

	// authorized api group
	authApiV1 := apiV1.Group("", app.authMiddleware.ValidateToken)
	authUserApi := authApiV1.Group("/user")
	{
		authUserApi.GET("/:user_id", app.userHandler.Detail)
		authUserApi.POST("/update", app.userHandler.Update)
	}

	authProfileApi := authApiV1.Group("/profile")
	{
		authProfileApi.GET("/:profile_user_id", app.profileHander.Profile)
		authProfileApi.POST("/:follow_user_id/follow", app.profileHander.Follow)
		authProfileApi.DELETE("/:unfollow_user_id/unfollow", app.profileHander.Unfollow)
	}
}
