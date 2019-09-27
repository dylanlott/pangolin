package collection

import (
	"github.com/dylanlott/pangolindb/pkg/kvstore"
)

// Collection holds the active Collection and its related data
type Collection struct {
	name string
	DB   *kvstore.Database
}
