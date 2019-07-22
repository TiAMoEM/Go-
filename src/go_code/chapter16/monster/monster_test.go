package monster

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() error, hope is %v, fact is %v", true, res)
	}
	t.Log("monster.Store() success")
}

func TestRestore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() error, hope is %v, fact is %v", true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() error, hope is %v, fact is %v", "红孩儿", monster.Name)
	}
	t.Log("monster.Store() success")
}
