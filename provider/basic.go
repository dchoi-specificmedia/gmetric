package provider

import (
	"github.com/dchoi-specificmedia/gmetric/counter"
	"github.com/dchoi-specificmedia/gmetric/stat"
)

type basic struct{}

func (p basic) Keys() []string {
	return []string{
		stat.ErrorKey,
		stat.Pending,
	}
}

//Map maps values into slice index
func (p basic) Map(value interface{}) int {
	if value == nil {
		return -1
	}
	if _, ok := value.(error); ok {
		return 0
	}
	if value == stat.Pending {
		return 1
	}
	return -1
}

//NewBasic creates a basic counter stats provider
func NewBasic() counter.Provider {
	return &basic{}
}
