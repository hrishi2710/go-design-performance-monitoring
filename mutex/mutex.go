package mutex

import "sync"

var (
	count int
	mutex sync.Mutex
)

func Increment(incrementBy int) int{
	mutex.Lock()
	count += incrementBy
	mutex.Unlock()
	return count
}