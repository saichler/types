package common

import "github.com/saichler/types/go/types"

// Add a bool for transaction
type IServicePoints interface {
	//Register a service point at an area and notify all listeners if the vnic is not nil
	RegisterServicePoint(IServicePointHandler, uint16, IVirtualNetworkInterface) error
	// Handle a message and forward to the handler
	Handle(IElements, Action, IVirtualNetworkInterface, IMessage, bool) IElements
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(IElements, IVirtualNetworkInterface, IMessage, bool) IElements
	// Return the service point handler for the service name and area
	ServicePointHandler(string, uint16) (IServicePointHandler, bool)
}

type IServicePointHandler interface {
	Post(IElements, IResources) IElements
	Put(IElements, IResources) IElements
	Patch(IElements, IResources) IElements
	Delete(IElements, IResources) IElements
	GetCopy(IElements, IResources) IElements
	Get(IElements, IResources) IElements
	Failed(IElements, IResources, IMessage) IElements
	ServiceName() string
	ServiceModel() IElements
	EndPoint() string
	Transactional() bool
	ReplicationCount() int
	ReplicationScore() int
}

type IServicePointCacheListener interface {
	PropertyChangeNotification(*types.NotificationSet)
}
