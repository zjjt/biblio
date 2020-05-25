package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/zjjt/biblio/user_service/config"
	"github.com/zjjt/biblio/user_service/handler"
	userProto "github.com/zjjt/biblio/user_service/proto/user"
	"github.com/zjjt/biblio/user_service/repository"
)

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetETCDConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
func init() {
	if os.Getenv("ENV") != "PROD" || os.Getenv("ENV") != "TEST" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatalf("Couldnt load .env file %v", err)
		}
	}

}
func main() {
	//handling database connection and deferring its closing
	db, err := config.CreatePostgresDBConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	} else {
		log.Info("Connected to DB successfully")
	}
	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&userProto.User{})
	repo := repository.InitRepository(db)
	//create registery with etcd
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("biblio.service.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()
	handler := handler.InitService(repo)
	// Register Handler
	userProto.RegisterUserServiceHandler(service.Server(), handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
