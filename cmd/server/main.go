package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"workflow_engine/internal/config"
	deliveryGrpc "workflow_engine/internal/delivery/grpc"
	"workflow_engine/internal/delivery/grpc/pb"
	"workflow_engine/internal/delivery/http/api/auth"
	"workflow_engine/internal/delivery/http/api/document"
	"workflow_engine/internal/delivery/http/middleware"
	"workflow_engine/internal/domain/usecase"
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

	// infra
	passHasher := service.NewBcryptHasher()
	jwtService := jwt.NewJWTService(cfg.JWTSecret)

	// repos
	userRepo := postgres.NewUserRepository(db)
	docRepo := postgres.NewDocumentRepository(db)
	workflowRepo := postgres.NewWorkFlowRepository(db)

	// usecases
	docUseCase := usecase.NewDocumentUseCase(docRepo, workflowRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo, passHasher)
	workflowUseCase := usecase.NewWorkFlowRepository(workflowRepo, docRepo)

	// HTTP
	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	docHandler := document.NewDocumentHandler(docUseCase)
	authHandler := auth.NewAuthHandler(authUseCase, jwtService)

	mux := http.NewServeMux()
	auth.RegisterRoutes(mux, authHandler)
	document.RegisterRoutes(mux, docHandler, authMiddleware)

	// gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWorkflowServer(
		grpcServer,
		deliveryGrpc.NewWorkflowGRPCServer(workflowUseCase),
	)

	go func() {
		log.Println("gRPC server started on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC error: %v", err)
		}
	}()

	log.Printf("HTTP server started on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("HTTP error: %v", err)
	}
}
