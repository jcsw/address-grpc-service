package mongodb

import (
	context "context"
	atomic "sync/atomic"
	time "time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	mongoOption "go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/jcsw/address-grpc-service/pkg/system/log"
	properties "github.com/jcsw/address-grpc-service/pkg/system/properties"
)

var (
	session *mongo.Client
	healthy int32
)

// Initialize initiliaze the mongodb connection
func Initialize() {
	session = connect()
	go monitor()
}

// IsAlive return mongoDB session status
func IsAlive() bool {
	return atomic.LoadInt32(&healthy) == 1
}

// GetSession Return a mongodb session
func GetSession() *mongo.Client {
	return session
}

// Close close the mongodb session
func Close() {
	if session != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		session.Disconnect(ctx)
		log.Info("p=mongodb f=Close m=mongodb_client_closed")
	}
}

func connect() *mongo.Client {

	client, err := mongo.NewClient(mongoOption.Client().ApplyURI(properties.Values.MongodbURI))

	if err != nil {
		log.Error("p=mongodb f=connect m=could_not_create_session err=%v", err)
		return nil
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Error("p=mongodb f=connect m=could_not_connect_at_mongodb err=%v", err)
		return nil
	}

	if err := client.Ping(nil, nil); err != nil {
		log.Error("p=mongodb f=connect m=could_not_ping_at_mongodb err=%v", err)
	}

	databases, _ := client.ListDatabases(nil, nil, nil)
	log.Info("p=mongodb f=connect m=session_created databases=%+v", databases)
	setMongoDBStatusUp()

	return client
}

func monitor() {
	for {

		if session == nil || session.Ping(nil, nil) != nil {
			setMongoDBStatusDown()
			log.Warn("p=mongodb f=monitor m=session_is_not_active_trying_to_reconnect")
			session = connect()
		} else {
			setMongoDBStatusUp()
			log.Info("p=mongodb f=monitor m=session_it's_alive'")
		}

		time.Sleep(30 * time.Second)
	}
}

func setMongoDBStatusUp() {
	atomic.StoreInt32(&healthy, 1)
}

func setMongoDBStatusDown() {
	atomic.StoreInt32(&healthy, 0)
}
