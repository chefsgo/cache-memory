package session_buntdb

import (
	cb "github.com/chefsgo/cache-buntdb"
	"github.com/chefsgo/chef"
)

func init() {
	chef.Register("memory", cb.Driver(":memory:"))
}
