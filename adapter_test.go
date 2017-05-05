package cassandra_adapter

import (
	"testing"
	"github.com/hsluoyz/casbin"
	"log"
	"github.com/hsluoyz/casbin/util"
)

func testGetPolicy(t *testing.T, e *casbin.Enforcer, res [][]string) {
	myRes := e.GetPolicy()
	log.Print("Policy: ", myRes)

	if !util.Array2DEquals(res, myRes) {
		t.Error("Policy: ", myRes, ", supposed to be ", res)
	}
}

func TestAdapter(t *testing.T) {
	e := casbin.NewEnforcer("examples/rbac_model.conf", "examples/rbac_policy.csv")

	a := NewAdapter("192.168.41.130")
	a.SavePolicy(e.GetModel())

	e.ClearPolicy()
	testGetPolicy(t, e, [][]string{})

	a.LoadPolicy(e.GetModel())
	testGetPolicy(t, e, [][]string{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"data2_admin", "data2", "read"}, {"data2_admin", "data2", "write"}})

	a = NewAdapter("192.168.41.130")
	e = casbin.NewEnforcer("examples/rbac_model.conf", a)
	testGetPolicy(t, e, [][]string{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"data2_admin", "data2", "read"}, {"data2_admin", "data2", "write"}})
}
