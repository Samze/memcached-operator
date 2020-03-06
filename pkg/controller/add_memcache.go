package controller

import (
	"github.com/samze/memcache-operator/pkg/controller/memcache"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, memcache.Add)
}
