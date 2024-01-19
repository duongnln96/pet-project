package http_server

func (app *app) routeVer1() {

	apiV1 := app.httpServer.GroupRouter("/api/v1")

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
