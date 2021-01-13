package client

import (
	clients "sigs.k8s.io/controller-runtime/pkg/client"
)

// Client interface from controller-runtime
type Client = clients.Client

// Injector allows injection of logger
type Injector interface {
	WithLogger(client Client)
}
