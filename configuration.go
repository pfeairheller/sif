package gomethius

import (

	"encoding/json"
	"os"
	"io"
)


type Configuration struct {
	ConfigHost string
	Application interface {}
}

func NewConfiguration(fileName string, appConfig interface{}) (c *Configuration) {
	c = new(Configuration)
	c.Application = appConfig
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)	
	err = dec.Decode(c)

	if err != nil && err != io.EOF {
		panic(err)
	}
	
	return
}
