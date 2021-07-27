package lesson4

type Counter struct {
	cnt int64
}

func (c *Counter) Incriment() {
	c.cnt++

}

func (c *Counter) Result() int64 {

	return c.cnt
}
