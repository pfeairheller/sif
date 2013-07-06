package gomethius

import (

	"encoding/json"
	"os"
	"io"
)


type Configuration struct {
	ConfigHost string
	Application map[string]string
}

func NewConfiguration(fileName string) (c *Configuration) {
	c = new(Configuration)

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	dec := json.NewDecoder(file)	
	err = dec.Decode(c)

	if err != nil && err != io.EOF {
		panic(err)
	}
	
	return
}
