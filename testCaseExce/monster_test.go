package testCaseExce

import "testing"

func TestMonster_Store(t *testing.T) {
	//先创建一个monster
	monster := &monster{"红孩儿", 200, "火龙戏珠"}
	if res := monster.Store(); !res {
		t.Errorf("monster.Store error")
	}
	t.Logf("monster.Store success")
}

func TestMonster_Restore(t *testing.T) {
	//先创建一个monster
	monster := &monster{}
	if res := monster.Restore(); !res {
		t.Errorf("monster.Restore error")
	}
	//进一步判断
	if monster.Name != "红孩儿" {
		t.Errorf("monster.Nnme not 红孩儿, error")
	}
	t.Logf("monster.Restore success")
}
