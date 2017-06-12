Cassandra Adapter for Casbin [![Godoc](https://godoc.org/github.com/casbin/cassandra_adapter?status.svg)](https://godoc.org/github.com/casbin/cassandra_adapter)
====

Cassandra Adapter is the [Apache Cassandra DB](http://cassandra.apache.org/) adapter for [Casbin](https://github.com/casbin/casbin). With this library, Casbin can load policy from Cassandra or save policy to it.

## Installation

    go get github.com/casbin/cassandra_adapter

## Simple Example

```go
package main

import (
	"github.com/casbin/casbin"
	"github.com/casbin/cassandra_adapter"
)

func main() {
	// Initialize a Cassandra adapter and use it in a Casbin enforcer:
	a := cassandra_adapter.NewAdapter("192.168.41.130") // Your Cassandra hosts. 
	e := casbin.NewEnforcer("examples/rbac_model.conf", a)
	
	e.Enforce("alice", "data1", "read")
}
```

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.
