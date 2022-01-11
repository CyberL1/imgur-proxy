package main

import (
	"embed"
	"fmt"
	"os"
	"imgurproxy/config"
	"imgurproxy/routers"

	"github.com/gin-gonic/gin"
)

//go:embed resources
var resources embed.FS

func main() {
	// Check for all required resources	
	if _, err := os.Stat("./public"); os.IsNotExist(err) {
		fmt.Println("public dir doesn't exist, creating...")
		os.Mkdir("public", 0775)
	}

	if _, err := os.Stat("./html"); os.IsNotExist(err) {
		fmt.Println("html dir doesn't exist, creating...")
		os.Mkdir("html", 0775)
		htmlDir, _ := resources.ReadDir("resources/html")
		for _, f := range htmlDir {
			hf, _ := resources.ReadFile(fmt.Sprintf("resources/html/%v", f.Name()))
			os.WriteFile(fmt.Sprintf("html/%v", f.Name()), hf, 0664)
		}
	}

	if _, err := os.Stat("./config.json"); os.IsNotExist(err) {
		cfgExample, _ := resources.ReadFile("resources/config.json.example")
		fmt.Println("config.json doesn't exist, creating...")
		os.WriteFile("config.json", cfgExample, 0775)
		return
	}

	// Set release mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize the gin router
	r := gin.Default()

	// Serve HTML files
	r.LoadHTMLGlob("./html/*.html")

	// Serve assets
	r.Static("/_p", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	// Register all routers
	routers.PagesRouters(r)
	routers.ResourcesRouters(r)

	// Run the proxy
	port := config.GetConfig().Proxy.Port
	r.Run(port)
}