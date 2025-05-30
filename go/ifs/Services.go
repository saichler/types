package ifs

import "github.com/saichler/l8types/go/types"

// Add a bool for transaction
type IServices interface {
	// Add a service point type so compiling will pull the code for it
	RegisterServiceHandlerType(IServiceHandler)
	// Activate a service point
	Activate(string, string, uint16, IResources, IServiceCacheListener, ...interface{}) (IServiceHandler, error)
	DeActivate(string, uint16, IResources, IServiceCacheListener) error
	// Handle a message and forward to the handler
	Handle(IElements, Action, IVNic, IMessage) IElements
	TransactionHandle(IElements, Action, IVNic, IMessage) IElements
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(IElements, IVNic, IMessage, bool) IElements
	// Return the service point handler for the service name and area
	ServiceHandler(string, uint16) (IServiceHandler, bool)
	// Register a distributed cache
	RegisterDistributedCache(cache IDistributedCache)
}

type IServiceHandler interface {
	Activate(string, uint16, IResources, IServiceCacheListener, ...interface{}) error
	DeActivate() error
	Post(IElements, IVNic) IElements
	Put(IElements, IVNic) IElements
	Patch(IElements, IVNic) IElements
	Delete(IElements, IVNic) IElements
	GetCopy(IElements, IVNic) IElements
	Get(IElements, IVNic) IElements
	Failed(IElements, IVNic, IMessage) IElements
	TransactionMethod() ITransactionMethod
	WebService() IWebService
}

type IServiceCacheListener interface {
	PropertyChangeNotification(*types.NotificationSet)
}

type IDistributedCache interface {
	Put(string, interface{}, ...bool) (*types.NotificationSet, error)
	Update(string, interface{}, ...bool) (*types.NotificationSet, error)
	Delete(string, ...bool) (*types.NotificationSet, error)
	Get(k string) interface{}
	Collect(f func(interface{}) (bool, interface{})) map[string]interface{}
	ServiceName() string
	ServiceArea() uint16
	Sync()
}

type ITransactionMethod interface {
	Replication() bool
	ReplicationCount() int
	KeyOf(IElements, IResources) string
}
