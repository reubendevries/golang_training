package main

import "testing"

func TestFoo(t *testing.T) {
	t.Error("intentional error 1")
}

func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {
		t.Errorf("expecing bar, go %s", result)
	}
}

func TestQuiz(t *testing.T) {
	t.Error("intentional error 2")
}
