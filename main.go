package main

import (
	"api/src/core"
	"api/src/order/infraestructure/adapters"
	"api/src/order/infraestructure/routes"
	productadapter "api/src/product/infrastructure/adapters"
	productroutes "api/src/product/infrastructure/routes"
	useradapter "api/src/user/infrastructure/adapters"
	userroutes "api/src/user/infrastructure/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := core.ConnectToDataBase()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	productRepo := productadapter.NewMySQLRepository(db)
	userRepo := useradapter.NewMySQLRepository(db)
	orderRepo := adapters.NewMySQLRepository(db)

	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Cambia esto según tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	productroutes.SetupProductRoutes(router, productRepo)
	userroutes.SetupUserRoutes(router, userRepo)
	routes.SetupOrderRoutes(router, orderRepo)

	log.Println("Iniciando el Servidor en el puerto 8080...")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
