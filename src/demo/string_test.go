package main

import "testing"

func TestConcatStr(t *testing.T) {
	result := ConcatStr([]string{"1", "2", "3"})
	t.Log(result)
}

func BenchmarkConcatStr(t *testing.B) {
	params := []string{"1", "2", "3"}
	for i := 0; i < t.N; i++ {
		result := ConcatStr(params)
		t.Log(result)
	}
}

func TestConcatStr1(t *testing.T) {
	result := ConcatStr1([]interface{}{"1", "2", "3"}...)
	t.Log(result)
}

func TestConcatStrs(t *testing.T) {
	t.Run("A=1", func(t *testing.T) {
		t.Log("A=1")
	})
	t.Run("A=2", func(t *testing.T) {
		t.Log("A=2")
	})
	t.Run("B=1", func(t *testing.T) {
		t.Log("B=1")
	})
}
