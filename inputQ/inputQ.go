package inputq


type counter struct {
	inputQ chan interface{}
	count int
}

func (c *counter) Initialize() {
	c.inputQ = make(chan interface{}, 0)
	c.count = 0
}

func (c *counter) Increment() {

}

func processInputQ() {

}