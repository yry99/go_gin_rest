package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	// "github.com/pressly/goose/v3"

	"sqlc/api/controllers"
	"sqlc/api/repository"
	"sqlc/api/routes"
)

func main() {
	// Initialize database connection pool
	// dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:123456@127.0.0.1:5432/go_commerce")

	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:123456@127.0.0.1:5432/go_commerce")

	// fmt.Print(dbpool.Stat())
	// var totalConns = dbpool.Stat()

	// fmt.Printf("totalConns: %v\n", totalConns)
	// // fmt.Print(totalConns)

	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Print("Connected to database with ping\n")

	defer dbpool.Close()

	/* Stats */
	// // Acquire a connection and perform a simple query to ensure the pool is initialized
	// conn, err := dbpool.Acquire(context.Background())
	// if err != nil {
	// 	log.Fatalf("Unable to acquire a connection: %v", err)
	// }
	// defer conn.Release()

	// // Fetch and print pool statistics
	// stats := dbpool.Stat()
	// fmt.Printf("Total Connections: %d\n", stats.TotalConns())
	// fmt.Printf("Idle Connections: %d\n", stats.IdleConns())
	// fmt.Printf("Acquired Connections: %d\n", stats.AcquiredConns())
	// fmt.Printf("Max Connections: %d\n", stats.MaxConns())
	// fmt.Printf("Total Queries: %d\n", stats.AcquireCount())
	// fmt.Printf("Average Connection Acquire Duration: %s\n", stats.AcquireDuration())

	// Initialize repository
	authorRepo := repository.NewAuthorRepository(dbpool)

	// Initialize controller
	authorController := controllers.NewAuthorController(authorRepo)

	router := gin.Default()

	routes.SetupAuthorRoutes(router, authorController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}

	// Run migrations
	// if err := runMigrations(); err != nil {
	// 	log.Fatalf("Unable to run migrations: %v\n", err)
	// }

	// Initialize sqlc queries
	// 	queries := db.New(dbpool)

	// 	// Set up Gin router
	// 	r := gin.Default()

	// 	// Define routes
	// 	r.POST("/users", createUser(queries))
	// 	r.GET("/users/:id", getUser(queries))
	// 	r.PUT("/users/:id", updateUser(queries))
	// 	r.DELETE("/users/:id", deleteUser(queries))

	// 	// Start server
	// 	if err := r.Run(":8080"); err != nil {
	// 		log.Fatalf("Failed to start server: %v\n", err)
	// 	}
	// }

	// func createUser(q *db.Queries) gin.HandlerFunc {
	// 	return func(c *gin.Context) {
	// 		var user db.CreateUserParams
	// 		if err := c.ShouldBindJSON(&user); err != nil {
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 			return
	// 		}

	// 		conn, err := q.GetConnection(c)
	// 		if err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get database connection"})
	// 			return
	// 		}
	// 		defer conn.Release()

	// 		createdUser, err := q.CreateUser(c, conn, user)
	// 		if err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 			return
	// 		}

	// 		c.JSON(http.StatusCreated, createdUser)
	// 	}
}
