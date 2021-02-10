package main

import (
	"github.com/WesleyCastilho/codepix/application/grpc"
	"github.com/WesleyCastilho/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
	"os"
)


var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}