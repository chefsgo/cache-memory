package session_buntdb

import (
	"github.com/chefsgo/cache"
	cb "github.com/chefsgo/cache-buntdb"
)

func init() {
	cache.Register("memory", cb.Driver(":memory:"))
}
