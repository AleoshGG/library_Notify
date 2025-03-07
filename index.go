package main

import (
	"library-Notify/src/infrastructure"
	"library-Notify/src/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Lanzar dependencias
	godotenv.Load()
	infrastructure.GoDependences()

	// Crear el router
	r := gin.Default()
	r.Use(cors.Default())

	// Registrar las rutas
	routes.RegisterRouter(r)

	r.Run(":8001") // Sirve y escucha peticiones en 0.0.0.0:8080
}