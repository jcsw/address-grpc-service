package main

import (
	flag "flag"
	os "os"
	signal "os/signal"

	mongodb "github.com/jcsw/address-grpc-service/pkg/driver/mongodb"
	log "github.com/jcsw/address-grpc-service/pkg/system/log"
	properties "github.com/jcsw/address-grpc-service/pkg/system/properties"
	server "github.com/jcsw/address-grpc-service/pkg/system/server"
)

var env string

func main() {
	flag.StringVar(&env, "env", "prod", "app environment")
	flag.Parse()

	log.Info("p=main f=main m=initialize_by_env_[%s]", env)
	properties.Load(env)

	mongodb.Initialize()

	grpcServer := server.GrpcServer{}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		grpcServer.Stop()
	}()

	grpcServer.Start()
}
