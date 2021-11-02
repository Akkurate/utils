package misc

const (
	// DefaultLimit is the default concurrency limit
	DefaultLimit = 100
)

// ConcurrencyLimiter object
type ConcurrencyLimiter struct {
	limit         int
	tickets       chan int
	numInProgress int32
}

// NewConcurrencyLimiter allocates a new ConcurrencyLimiter
func NewConcurrencyLimiter(limit int) *ConcurrencyLimiter {
	if limit <= 0 {
		limit = DefaultLimit
	}
	// allocate a limiter instance
	c := &ConcurrencyLimiter{
		limit:   limit,
		tickets: make(chan int, limit),
	}
	// allocate the tickets:
	for i := 0; i < c.limit; i++ {
		c.tickets <- i
	}
	return c
}

// Execute adds a function to the execution queue.
// if num of go routines allocated by this instance is < limit
// launch a new go routine to execute job
// else wait until a go routine becomes available
func (c *ConcurrencyLimiter) Execute(job func()) {
	ticket := <-c.tickets
	go func() {

		// run the job
		job()
		c.tickets <- ticket
	}()

}

// ExecuteIndex ExecuteIndex
func (c *ConcurrencyLimiter) ExecuteIndex(index int, job func(numb int)) {
	ticket := <-c.tickets
	go func(o int) {

		// run the job
		job(o)
		c.tickets <- ticket
	}(index)

}

func (c *ConcurrencyLimiter) Loop(n int, job func(o int)) {
	for i := 0; i < n; n++ {
		ticket := <-c.tickets
		go func(num int) {

			// run the job
			job(num)
			c.tickets <- ticket
		}(i)
	}

}

// Wait will block all the previously Executed jobs completed running.
//
// IMPORTANT: calling the Wait function while keep calling Execute leads to
//            un-desired race conditions
func (c *ConcurrencyLimiter) Wait() {
	for i := 0; i < c.limit; i++ {
		_ = <-c.tickets
	}
}
