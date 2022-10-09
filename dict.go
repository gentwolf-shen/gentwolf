package gentwolf

import (
	"encoding/json"
	"io/ioutil"
)

type dict struct {
	items map[string]string
}

var Dict *dict

func init() {
	Dict = &dict{}
}

func (this *dict) Init(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &this.items)
	return err
}

func (this *dict) Get(key string) string {
	return this.items[key]
}

func (this *dict) Set(key, value string) {
	this.items[key] = value
}
