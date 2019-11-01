package datasource

import (
	"errors"
	"github.com/kataras/iris/_examples/mvc/login/datamodels"
)

type Engine uint32

const   (
	Memory Engine = iota
	Bolt
	MySQL
)

func LoadUsers(engine Engine) (map[int64]datamodels.User, error) {
	if engine != Memory{
		return nil, errors.New("for the shake of simplicity we're using a simple map as the data source")
	}

	return make(map[int64]datamodels.User),nil
}