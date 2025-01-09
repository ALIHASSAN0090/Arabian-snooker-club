package router

import "arabian-snooker/middleware"

func (r *Router) SetupRoutes() {
	rateLimiter := middleware.NewRateLimiter(1, 5)
	r.Engine.Use(middleware.EnableCors(), rateLimiter.Limit())

	adminGroup := r.Engine.Group("/admin")
	{
		adminGroup.Use(middleware.Auth([]string{"admin"}))

		adminGroup.GET("/health-check", r.HealthCheck)

	}

	sellerGroup := r.Engine.Group("/seller")
	{
		sellerGroup.Use(middleware.Auth([]string{"seller", "admin"}))
		sellerGroup.Use(middleware.StatusCheck(r.Engine.HandleContext))
		sellerGroup.POST("/match-rates", r.CreateRates)
		sellerGroup.PATCH("/match-rates/:id", r.UpdateRate)

	}

	publicGroup := r.Engine.Group("/public")
	{

		publicGroup.GET("/health-check", r.HealthCheck)
		// publicGroup.POST("/signup", r.SignUp)
		publicGroup.POST("/login", r.Login)

	}
}
