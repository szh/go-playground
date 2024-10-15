package main

import (
	"math/rand"
	"strconv"
)

type foo struct {
	lastname  string
	firstname string
}

type bar struct {
	foos map[string][]foo
}

type zig struct {
	zag bar
}

func (z zig) DoSomething() {
	z.addStuff()
	z.printStuff()
	z.clearStuff(false)
}

func (z zig) DoSomethingFullCleanse() {
	z.addStuff()
	z.printStuff()
	z.clearStuff(true)
}

func (z zig) addStuff() {
	// Generate a random number
	num := rand.Int()
	z.zag.foos[strconv.Itoa(num)] = []foo{
		{
			lastname:  "Doe",
			firstname: "John",
		},
	}
}

func (z zig) printStuff() {
	for k, v := range z.zag.foos {
		println(k)
		for _, f := range v {
			println(f.lastname)
			println(f.firstname)
		}
	}
}

func (z zig) clearStuff(fullCleanse bool) {
	if fullCleanse {
		for f := range z.zag.foos {
			delete(z.zag.foos, f)
		}
		return
	}
	// z.zag.foos = map[string][]foo{}
	z.zag.foos = make(map[string][]foo)
}

func main() {
	obj := zig{
		zag: bar{
			foos: map[string][]foo{},
		},
	}

	obj.DoSomething()
	obj.DoSomething()
}
