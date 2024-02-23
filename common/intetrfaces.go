package common

// Closable is the interface for objects that can release its resources.
//
// xray:api:beta
type Closable interface {
	// Close release all resources used by this object, including goroutines.
	Close() error
}

// Runnable is the interface for objects that can start to work and stop on demand.
type Runnable interface {
	// Start starts the runnable object. Upon the method returning nil, the object begins to function properly.
	Start() error

	Closable
}
