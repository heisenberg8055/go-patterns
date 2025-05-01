package strategy

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(e EvictionAlgo) *Cache {
	mp := make(map[string]string)
	return &Cache{storage: mp, evictionAlgo: e, capacity: 0, maxCapacity: 0}
}

func (c *Cache) SetEvictionPolicy(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.Evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
}

func (c *Cache) Evict() {
	c.evictionAlgo.Evict(c)
	c.capacity--
}
