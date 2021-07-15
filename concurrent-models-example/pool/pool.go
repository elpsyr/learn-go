package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	ErrPoolClosed = errors.New("pool has been closed")
)

type Pool struct {
	factory   func() (io.Closer, error) //创建资源
	resources chan io.Closer            //资源通道
	mtx       sync.Mutex                //一把锁
	closed    bool                      //房门
}

//提供资源和资源管道大小
func New(factory func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("invalid size for the resources pool")
	}
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, size),
		closed:    false,
	}, nil

}

func (p *Pool) AcquireResource() (io.Closer, error) {
	select {
	case resource, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("acquire resource from the pool")
		return resource, nil
	default:
		fmt.Println("acquire new resource")
		return p.factory()
	}

}

func (p *Pool) ReleaseResource(resource io.Closer) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	if p.closed {
		resource.Close()
		return
	}

	select {
	case p.resources <- resource:
		fmt.Println("release resource back to the pool")
	default:
		fmt.Println("pool filled")
		resource.Close()
	}

}

func (p *Pool) Close() {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for resource := range p.resources {
		resource.Close()
		fmt.Println("closing resource ")
	}
}
