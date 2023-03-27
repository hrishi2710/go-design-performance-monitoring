package inputq


var counter = 0;
var inputQ = make(chan interface{});

type increment struct {
	incrementBy int
	responseChan chan int
}

func Increment(incrementBy int) int {
	responseChan := make(chan int, 1);
	inputQ <-increment{incrementBy: incrementBy, responseChan: responseChan}
	return <-responseChan;
}

func ProcessInputQ() {
	for {
		item := <-inputQ;
		v := item.(increment);
		inc(v.incrementBy, v.responseChan)
	}
}

func inc(incrementBy int, responseChan chan int){
	counter += incrementBy;
	responseChan <- counter;
}