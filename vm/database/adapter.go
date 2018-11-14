package database

type database interface {
	Get(key string) (value string)
	Put(key, value string)
	Has(key string) bool
	Del(key string)
}

const (
	// StateTable name
	StateTable = "state"
)

type chainbaseAdapter struct {
	cb  IMultiValue
}

func (c *chainbaseAdapter) Get(key string) (value string) {
	var err error
	value, err = c.cb.Get(StateTable, key)
	if err != nil {
		panic(err)
	}
	if value == "" {
		return NilPrefix
	}
	return
}

func (c *chainbaseAdapter) Put(key, value string) {
	err := c.cb.Put(StateTable, key, value)
	if err != nil {
		panic(err)
	}
}

func (c *chainbaseAdapter) Has(key string) bool {
	ok, err := c.cb.Has(StateTable, key)
	if err != nil {
		panic(err)
	}
	return ok
}

func (c *chainbaseAdapter) Del(key string) {
	err := c.cb.Del(StateTable, key)
	if err != nil {
		panic(err)
	}
}

func newChainbaseAdapter(cb IMultiValue) *chainbaseAdapter {
	return &chainbaseAdapter{cb}
}
