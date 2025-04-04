package common

// Add a bool for transaction
type IServicePoints interface {
	// Register A Service Point, handler + service area
	RegisterServicePoint(IServicePointHandler, uint16) error
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
