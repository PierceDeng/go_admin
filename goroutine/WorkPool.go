package goroutine

import "sync"

type WorksPool struct {
	TaskList chan func()
	Wg       *sync.WaitGroup
	QuitChan chan struct{}
}

func NewWorksPool(size int) *WorksPool {

	pool := &WorksPool{
		TaskList: make(chan func()),
		Wg:       &sync.WaitGroup{},
		QuitChan: make(chan struct{}),
	}

	for i := 0; i < size; i++ {
		pool.Wg.Add(1)
		go pool.Start()
	}
	return pool

}

func (pool *WorksPool) Start() {
	defer pool.Wg.Done()
	for {
		select {
		case task := <-pool.TaskList:
			task()
		case <-pool.QuitChan:
			return
		}
	}

}

func (pool *WorksPool) Submit(task func()) {
	pool.TaskList <- task
}

func (pool *WorksPool) Stop() {

	close(pool.TaskList)
	close(pool.QuitChan)

	pool.Wg.Wait()

}
