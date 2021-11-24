package main

import (
	"OnlineBank/controllers"
	"OnlineBank/database"
	"OnlineBank/repositories"
)

func main()  {
	controllers.InitLogger()
	database.ConnToDB()
	repositories.StartRoutes()
}
