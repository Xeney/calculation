package main

func initializeRoutes() {
	router.GET("/", SessionCheck, GetHome)
	router.POST("/", SessionCheck, PostHome)
	router.GET("/login", SessionCheck, GetLogin)
	router.POST("/login", SessionCheck, PostLogin)
	router.GET("/lk", SessionCheck, GetLk)
	router.POST("/lk", SessionCheck, PostLk)
	router.GET("/exit", SessionCheck, GetExit)
	router.GET("/result/:Name", SessionCheck, GetResult)
}
