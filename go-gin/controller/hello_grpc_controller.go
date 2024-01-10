package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go1/services/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"

	pb "go1/hello"
)

func HelloGrpc(c *gin.Context) {

	//	addr := "localhost:50051"
	addr := "go-grpc-server:50051"

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.DefaultLogger.Error("did not connect", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "did not connect"})
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.DefaultLogger.Error("did not close connect", err)
		}
	}(conn)

	grpcClient := pb.NewHiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user := pb.HiUser{
		Name: "gi",
		Sex:  44,
	}

	r, err := grpcClient.SayHello(ctx, &user)
	if err != nil {
		logger.DefaultLogger.Error("did not HiReply", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "did not HiReply"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": r.GetMessage(),
		"user":    r.GetUser(),
	})
}

func HelloPhpGrpc(c *gin.Context) {

	//	addr := "localhost:50051"
	addr := "gserver:9503"

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.DefaultLogger.Error("did not connect", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "did not connect"})
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.DefaultLogger.Error("did not close connect", err)
		}
	}(conn)

	grpcClient := pb.NewHiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user := pb.HiUser{
		Name: "gi",
		Sex:  44,
	}

	r, err := grpcClient.SayHello(ctx, &user)
	if err != nil {
		logger.DefaultLogger.Error("did not HiReply", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "did not HiReply"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": r.GetMessage(),
		"user":    r.GetUser(),
	})
}
