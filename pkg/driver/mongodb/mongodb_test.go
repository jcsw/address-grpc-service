// +build integration

package mongodb_test

import (
	testing "testing"

	assert "github.com/stretchr/testify/assert"

	mongodb "github.com/jcsw/address-grpc-service/pkg/driver/mongodb"
	properties "github.com/jcsw/address-grpc-service/pkg/system/properties"
)

func TestShouldInitializeMongoDBSession(t *testing.T) {

	properties.Values =
		properties.Schema{
			MongodbURI: "mongodb://address_service_adm:address_service_pwd@localhost:27017/admin?connectTimeoutMS=100&socketTimeoutMS=500",
		}

	mongodb.Initialize()
	defer mongodb.Close()

	assert.True(t, mongodb.IsAlive())
}
