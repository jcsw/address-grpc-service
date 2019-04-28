package properties

import (
	ioutil "io/ioutil"
	os "os"

	yaml "gopkg.in/yaml.v2"

	log "github.com/jcsw/address-grpc-service/pkg/system/log"
)

// Schema define the properties values
type Schema struct {
	ServerPort int    `yaml:"server_port"`
	MongodbURI string `yaml:"mongodb_uri"`
}

// Values the loaded properties values
var Values Schema

// Load load properties in Values
func Load(env string) {

	pwd, _ := os.Getwd()
	fileProperties, err := ioutil.ReadFile(pwd + "/properties/" + env + ".yaml")
	if err != nil {
		log.Fatal("p=properties f=Load \n%v", err)
	}

	err = yaml.UnmarshalStrict(fileProperties, &Values)
	if err != nil {
		log.Fatal("p=properties f=Load \n%v", err)
	}

	log.Info("p=properties f=Load m=properties_loaded")
}
