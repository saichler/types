package ifs

type SerializerMode int

const (
	BINARY SerializerMode = 1
	JSON   SerializerMode = 2
	STRING SerializerMode = 3
)

type ISerializer interface {
	Mode() SerializerMode
	Marshal(interface{}, IResources) ([]byte, error)
	Unmarshal([]byte, IResources) (interface{}, error)
}
