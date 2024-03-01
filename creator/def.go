package creator

import (
	"fmt"
	"github.com/go-kid/kioc/util"
	"reflect"
)

type FileCreator interface {
	Create() error
}

type BatchCreator struct {
	files []FileCreator
}

func NewBatchCreator(files ...FileCreator) *BatchCreator {
	return &BatchCreator{files: files}
}

func (p *BatchCreator) Create() error {
	for _, file := range p.files {
		if err := file.Create(); err != nil {
			return fmt.Errorf("%s: %s", reflect.TypeOf(file).String(), err)
		}
	}
	return nil
}

var cmd util.GoCmd
