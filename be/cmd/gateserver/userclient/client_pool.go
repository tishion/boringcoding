package userclient

import (
	"errors"
	"sync"

	// local packages
	"mmc.com/landingtask/be/internal/common"
)

// The pre-defined errors
var (
	ErrInvalidConfig = errors.New("invalid pool config")
	ErrPoolClosed    = errors.New("pool closed")
)

// Represents the client factory
type ClientFactory func() (*UserClient, error)

// Represents the client pool
type ClientPool struct {
	sync.Mutex
	pool    chan *UserClient
	maxOpen int
	numOpen int
	minOpen int
	closed  bool
	factory ClientFactory
}

// Creates client pool
func NewClientPool(minOpen int, maxOpen int, factory ClientFactory) *ClientPool {
	if maxOpen <= 0 || minOpen > maxOpen {
		common.LogF("Invalid config for client pool")
		// terminate
	}

	p := &ClientPool{
		maxOpen: maxOpen,
		minOpen: minOpen,
		factory: factory,
		pool:    make(chan *UserClient, maxOpen),
	}

	for i := 0; i < minOpen; i++ {
		client, err := factory()
		if err != nil {
			continue
		}
		p.numOpen++
		p.pool <- client
	}

	return p
}

// Acquires a client from the pool
func (p *ClientPool) Acquire() (*UserClient, error) {
	if p.closed {
		return nil, ErrPoolClosed
	}

	client, err := p.getOrCreate()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Gets a client from the pool, if no available creates new one and returns.
func (p *ClientPool) getOrCreate() (*UserClient, error) {
	select {
	case client := <-p.pool:
		return client, nil
	default:
	}

	p.Lock()
	defer p.Unlock()

	if p.numOpen >= p.maxOpen {
		client := <-p.pool
		return client, nil
	}

	client, err := p.factory()
	if err != nil {
		return nil, err
	}

	p.numOpen++
	return client, nil
}

// Release the client namely return it back to the pool for next comsuming
func (p *ClientPool) Release(client *UserClient) error {
	if p.closed {
		return ErrPoolClosed
	}
	p.pool <- client
	return nil
}

// Closes the client
func (p *ClientPool) Close(client *UserClient) error {
	p.Lock()
	defer p.Unlock()

	client.Close()
	p.numOpen--

	return nil
}

// Shuts the client pool down
func (p *ClientPool) Shutdown() error {
	if p.closed {
		return ErrPoolClosed
	}

	p.Lock()
	defer p.Unlock()

	close(p.pool)
	for client := range p.pool {
		client.Close()
		p.numOpen--
	}

	p.closed = true
	return nil
}
