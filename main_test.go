package main

import "testing"

func TestMain(t *testing.T) {
	// Test the addStuff function
	z := zig{
		zag: bar{
			foos: make(map[string][]foo),
		},
	}
	z.addStuff()
	if len(z.zag.foos) != 1 {
		t.Error("Expected 1 foo, got", len(z.zag.foos))
	}

	// Test the printStuff function
	z.printStuff()

	// Test the clearStuff function
	z.clearStuff(false)
	if len(z.zag.foos) != 0 {
		t.Error("Expected 0 foo, got", len(z.zag.foos))
	}

	// Test the DoSomething function
	z.DoSomething()
	if len(z.zag.foos) != 1 {
		t.Error("Expected 1 foo, got", len(z.zag.foos))
	}

	// Test the DoSomethingFullCleanse function
	z.DoSomethingFullCleanse()
	if len(z.zag.foos) != 0 {
		t.Error("Expected 0 foo, got", len(z.zag.foos))
	}
}
