package builder

import (
	"sync"

	flatbuffers "github.com/google/flatbuffers/go"
)

const builderInitSize = 1024

// Pool - pool with builders.
type Pool struct {
	mu     sync.Mutex
	pool   chan *flatbuffers.Builder
	maxCap int
}

// NewBuilderPool - create new pool with max capacity (maxCap)
func NewBuilderPool(maxCap int) *Pool {
	return &Pool{
		pool:   make(chan *flatbuffers.Builder, maxCap),
		maxCap: maxCap,
	}
}

// Get - return builder or create new if it is empty
func (p *Pool) Get() *flatbuffers.Builder {
	p.mu.Lock()
	defer p.mu.Unlock()

	select {
	case builder := <-p.pool:
		return builder
	default:
		return flatbuffers.NewBuilder(builderInitSize)
	}
}

// Put return builder to the pool
func (p *Pool) Put(builder *flatbuffers.Builder) {
	p.mu.Lock()
	defer p.mu.Unlock()

	builder.Reset()

	select {
	case p.pool <- builder:
		// return to the pool
	default:
		// ignore
	}
}
