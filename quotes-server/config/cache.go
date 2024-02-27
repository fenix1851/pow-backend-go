package config

type Cache struct {
	LeadingZerosCountCache *LeadingZerosCountCache
}

type LeadingZerosCountCache struct {
	// LeadingZerosCount - the number of leading zeros for the PoW
	LeadingZerosCount int64
	// LastUpdate - the last time the leading zeros count was updated
	LastUpdate int64
}

// getter and setter for the cache
func (c *Cache) GetLeadingZerosCount() (int64, int64, error) {
	return c.LeadingZerosCountCache.LeadingZerosCount, 0, nil
}

func (c *Cache) SetLeadingZerosCount(leadingZerosCount int64, lastUpdate int64) error {
	c.LeadingZerosCountCache.LeadingZerosCount = leadingZerosCount
	return nil
}

func (c *Cache) GetLastUpdate() (int64, error) {
	return c.LeadingZerosCountCache.LastUpdate, nil
}

func (c *Cache) SetLastUpdate(lastUpdate int64) error {
	c.LeadingZerosCountCache.LastUpdate = lastUpdate
	return nil
}
