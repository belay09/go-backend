package main
//nodemon --exec go run main.go --signal SIGTERM
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
	"go_backend/routes"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	routes := r.Group("/")
	router.InitializeRoutes(routes)
	port := os.Getenv("GO_PORT")
	if port == "" {
		fmt.Println("Port Not Found")
		port = "3000"
	}
	r.Run(":" + port)
}
