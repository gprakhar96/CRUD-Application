package app

import (
	"mypackage/controllers/ping"
	"mypackage/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)

	router.POST("/user", users.CreateUser)            // create a user
	router.GET("/user/:user_id", users.GetUser)       // get user by id
	router.PUT("/user/:user_id", users.UpdateUser)    // update user details based on user_id
	router.DELETE("/user/:user_id", users.DeleteUser) // delete user based on user_id
}
