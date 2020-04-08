package main

import (
	"fmt"
	"testing"
)

//功能测试
func TestHello(t *testing.T)  {
	var name string
	greeting, err := hello(name)
	if err == nil {
		t.Errorf("This error is nil, but it should not be. (name=%q)", name)
	}

	if greeting != "" {
		t.Errorf("Nonempty greeting, but it should not be. (name=%q)", name)
	}

	name = "Robert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("This error is not nil, but it should be. (name=%q)", name)
	}
	if greeting == "" {
		t.Errorf("Empty gretty, but it should not be. (name=%q)", name)
	}

	expected := fmt.Sprintf("hello, %s\n", name)
	if greeting != expected {
		t.Errorf("The actual greeting %q is not excepted. (name=%q)", greeting, name)
	}
	t.Logf("The expceted greeting is %q.\n", expected)
}

func TestIntroduce(t *testing.T)  {
	intro := introduce()
	expected := "Welcome to my Golang column."
	if intro != expected {
		t.Errorf("The actual introduce %q is bot the expceted.", intro)
	}
	t.Logf("The expceted introduce is %q.\n", intro)
}
