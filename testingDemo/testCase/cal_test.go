package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	//调用
	if res := addUpper(10); res != 55 {
		//fmt.Printf("addUpper error,期望值=%v,实际值=%v\n", 55, res)
		t.Fatalf("addUpper error,期望值=%v,实际值=%v\n", 55, res)
	}
	t.Logf("addUpper 执行正确...")
}

func TestGetSub(t *testing.T) {
	//调用
	if res := getSub(10,3); res != 7 {
		t.Fatalf("getSub error,期望值=%v,实际值=%v\n", 7, res)
	}
	t.Logf("getSub 执行正确...")
}
