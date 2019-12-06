package cache

import "github.com/monasuite/monautil"

// CacheableBlock is a wrapper around the btcutil.Block type which provides a
// Size method used by the cache to target certain memory usage.
type CacheableBlock struct {
	*monautil.Block
}

// Size returns size of this block in bytes.
func (c *CacheableBlock) Size() (uint64, error) {
	return uint64(c.Block.MsgBlock().SerializeSize()), nil
}
