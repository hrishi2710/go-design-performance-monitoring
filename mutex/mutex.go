package mutex

import "sync"

var (
	count int
	mutex sync.Mutex
)

func Increment() int{
	mutex.Lock()
	count++
	mutex.Unlock()
	return count
}