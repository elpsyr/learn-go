package runner

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("cannot finish tasks within the timeout")
	ErrInterrupt = errors.New("received interrupt from OS")
)

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

func New(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t),
		tasks:     make([]func(int), 0),
	}

}
func (r *Runner) AddTask(task ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

//工作线程观察有没有收到停止信号
func (r *Runner) run() error {
	for id, task := range r.tasks {
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterrupt
		default:
			task(id)
		}
	}
	return nil

}

func (r *Runner) Start() error {
	//relay interrupt from OS
	signal.Notify(r.interrupt, os.Interrupt)
	//run the task
	go func() {
		r.complete <- r.run()
	}()
	fmt.Println("waiting...")

	//主线程观察complete&timeout通道
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}

}
