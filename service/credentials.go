package service

import (
	"encoding/json"
	"fmt"
	"github.com/molsbee/tweet-streamer/config"
	"github.com/molsbee/tweet-streamer/model"
	"io/ioutil"
)

type Credentials struct {
}

func (cs *Credentials) Save(credentials model.Credentials) {
	data, _ := json.Marshal(credentials)
	ioutil.WriteFile(config.FilePath(), data, 0777)
}

func (cs *Credentials) Get() (*model.Credentials, error) {
	data, err := ioutil.ReadFile(config.FilePath())
	if err != nil {
		return nil, fmt.Errorf("unable to read file to retrieve credentials")
	}

	credentials := &model.Credentials{}
	err = json.Unmarshal(data, credentials)
	if err != nil {
		return nil, fmt.Errorf("unable to convert data in file %s", config.FilePath())
	}

	return credentials, nil
}
