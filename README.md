Cassandra Adapter [![Build Status](https://travis-ci.org/casbin/cassandra-adapter.svg?branch=master)](https://travis-ci.org/casbin/cassandra-adapter) [![Coverage Status](https://coveralls.io/repos/github/casbin/cassandra-adapter/badge.svg?branch=master)](https://coveralls.io/github/casbin/cassandra-adapter?branch=master) [![Godoc](https://godoc.org/github.com/casbin/cassandra-adapter?status.svg)](https://godoc.org/github.com/casbin/cassandra-adapter)
====

Cassandra Adapter is the [Apache Cassandra DB](http://cassandra.apache.org/) adapter for [Casbin](https://github.com/casbin/casbin). With this library, Casbin can load policy from Cassandra or save policy to it.

## Installation

    go get github.com/casbin/cassandra-adapter

## Simple Example

```go
package main

import (
	"github.com/casbin/casbin"
	"github.com/casbin/cassandra-adapter"
)

func main() {
	// Initialize a Cassandra adapter and use it in a Casbin enforcer:
	a := cassandra-adapter.NewAdapter("127.0.0.1") // Your Cassandra hosts. 
	e := casbin.NewEnforcer("examples/rbac_model.conf", a)
	
	// Load the policy from DB.
	e.LoadPolicy()
	
	// Check the permission.
	e.Enforce("alice", "data1", "read")
	
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	
	// Save the policy back to DB.
	e.SavePolicy()
}
```

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.
