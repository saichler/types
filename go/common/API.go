package common

import "github.com/saichler/types/go/types"

type IElements interface {
	Elements() []interface{}
	Keys() []interface{}
	Errors() []error
	Element() interface{}
	Query(IResources) IQuery
	Key() interface{}
	Error() error
	Serialize() ([]byte, error)
	Deserialize([]byte, IRegistry) error
}

type IQuery interface {
	RootType() *types.RNode
	Properties() map[string]IProperty
}

type IProperty interface {
	Get(interface{}) (interface{}, error)
	Set(interface{}, interface{}) (interface{}, interface{}, error)
}
