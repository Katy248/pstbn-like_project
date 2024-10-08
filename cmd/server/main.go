package main

import (
	h "cdecode/pkg/handlers"
	"cdecode/pkg/storage"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

const defaultUsername = "admin"

var db *sql.DB

func main() {
	// db setup
	db = storage.InitDB()

	defer db.Close()

	// server setup
	r := gin.Default()
	log.SetOutput(gin.DefaultErrorWriter)

	// routes setup
	aliases := r.Group("/alias")
	{
		aliases.GET("/", h.GetAliases(db))
		aliases.POST("/new", h.CreateAliases(db))
	}
	users := r.Group("/user")
	{
		users.GET("/", h.GetUsers(db))                // List users
		users.GET("/data/:id", h.GetUserDataById(db)) // Get user data
		users.GET("/data/", h.GetUserData(db))        // Get current authorized user data
		users.POST("/new", h.CreateUser(db))          // Register
		users.POST("/change/password", dummyHandler)  // Change password
		users.POST("/change/name", dummyHandler)      // Change username
	}

	r.POST("/auth", h.Authenticate(db))

	resources := r.Group("resource")
	{
		resources.GET("/", h.GetResources(db))                // Get users resources
		resources.POST("/create", h.CreateResource(db))       // Create new resource
		resources.DELETE("/delete/:id", h.DeleteResource(db)) // Delete existing resource
		resources.POST("/edit", h.UpdateResource(db))         // Edit existing resource content
	}

	r.Run(":1488")
}

func dummyHandler(ctx *gin.Context) {
	log.Printf("Dummy handler")
	ctx.JSON(200, "Dummy handler")
}
