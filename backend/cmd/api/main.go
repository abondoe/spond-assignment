package main

import (
	"abondoe/spond-assignment/internal/handler"
	"abondoe/spond-assignment/internal/repository"
	"abondoe/spond-assignment/internal/service"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/spond?sslmode=disable"
	}

	// 2. Hent port fra miljøvariabel, fall tilbake til 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Fatal("could not connect to database:", err)
	}

	// Dependency injection
	formRepo := repository.NewFormRepository()
	formService := service.NewFormService(formRepo)
	formHandler := handler.NewFormHandler(formService)

	registrationRepo := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepo, formRepo)
	registrationHandler := handler.NewRegistrationHandler(registrationService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			log.Printf("Health check failed: %v", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	if os.Getenv("APP_ENV") == "test" {
		mux.HandleFunc("POST /api/test/reset", func(w http.ResponseWriter, r *http.Request) {
			_, err := db.Exec("TRUNCATE registrations RESTART IDENTITY CASCADE")
			if err != nil {
				log.Printf("Reset failed: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			log.Println("Database reset successful")
			w.WriteHeader(http.StatusOK)
		})
	}

	mux.HandleFunc("GET /api/forms/{id}", formHandler.GetForm)
	mux.HandleFunc("POST /api/registrations", registrationHandler.CreateRegistration)

	log.Printf("Server starting on :%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
