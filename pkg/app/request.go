package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/jamesluo111/gin-blog/pkg/logging"
)

func MakeErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Value)
	}

	return
}
