Cassandra adapter for casbin [![Godoc](https://godoc.org/github.com/casbin/cassandra_adapter?status.svg)](https://godoc.org/github.com/casbin/cassandra_adapter)
====

**cassandra_adapter** is the [Apache Cassandra DB](http://cassandra.apache.org/) adapter for [casbin](https://github.com/hsluoyz/casbin). With it, casbin can load policy from Cassandra or save policy to it.

## Get started

The usage is very simple:

```golang

import (
	"github.com/hsluoyz/casbin"
	"github.com/casbin/cassandra_adapter"
)

// Initialize a Cassandra adapter and use it in a casbin enforcer:
a := NewAdapter("192.168.41.130") // Your Cassandra hosts. 
e := casbin.NewEnforcer("examples/rbac_model.conf", a)
```

For usage about casbin, please refer to: https://github.com/hsluoyz/casbin
