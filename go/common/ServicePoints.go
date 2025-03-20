package common

import (
	"github.com/saichler/types/go/types"
	"google.golang.org/protobuf/proto"
)

// Add a bool for transaction
type IServicePoints interface {
	RegisterServicePoint(string, int32, IServicePointHandler) error
	Handle(proto.Message, types.Action, IVirtualNetworkInterface, *types.Message, bool) (proto.Message, error)
	Notify(proto.Message, IVirtualNetworkInterface, *types.Message, bool) (proto.Message, error)
	ServicePointHandler(string) (IServicePointHandler, bool)
}

type IServicePointHandler interface {
	Post(proto.Message, IResources) (proto.Message, error)
	Put(proto.Message, IResources) (proto.Message, error)
	Patch(proto.Message, IResources) (proto.Message, error)
	Delete(proto.Message, IResources) (proto.Message, error)
	GetCopy(proto.Message, IResources) (proto.Message, error)
	Get(proto.Message, IResources) (proto.Message, error)
	Failed(proto.Message, IResources, *types.Message) (proto.Message, error)
	EndPoint() string
	Multicast() string
	Transactional() bool
	ReplicationCount() int
	ReplicationScore() int
}
