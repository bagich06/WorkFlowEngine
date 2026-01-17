package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"workflow_engine/internal/config"
	"workflow_engine/internal/delivery/http/api/auth"
	"workflow_engine/internal/delivery/http/api/document"
	"workflow_engine/internal/delivery/http/api/workflow"
	"workflow_engine/internal/delivery/http/middleware"
	"workflow_engine/internal/domain/usecase"
	workflowusecase "workflow_engine/internal/domain/usecase/workflow"
	"workflow_engine/internal/infrastructure/repository/postgres"
	"workflow_engine/internal/infrastructure/service"
	"workflow_engine/pkg/jwt"
)

func main() {
	_ = godotenv.Load()
	cfg := config.LoadConfig()

	db, err := postgres.NewDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	passHasher := service.NewBcryptHasher()

	userRepo := postgres.NewUserRepository(db)
	docRepo := postgres.NewDocumentRepository(db)

	docUseCase := usecase.NewDocumentUseCase(docRepo)
	wfUseCase := workflowusecase.NewWorkflowUseCase(docRepo, userRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo, passHasher)

	jwtService := jwt.NewJWTService(cfg.JWTSecret)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	docHandler := document.NewDocumentHandler(docUseCase)
	wfHandler := workflow.NewWorkflowHandler(wfUseCase)
	authHandler := auth.NewAuthHandler(authUseCase, jwtService)

	mux := http.NewServeMux()

	auth.RegisterRoutes(mux, authHandler)
	document.RegisterRoutes(mux, docHandler, authMiddleware)
	workflow.RegisterRoutes(mux, wfHandler, authMiddleware)

	log.Printf("Server started on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
