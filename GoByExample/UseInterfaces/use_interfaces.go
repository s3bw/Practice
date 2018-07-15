package main

import "fmt"

// From:
// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

/* Introduction to interfaces:

So what is an interface? An interface is two things: it is a set of
	methods, but it is also a type. Let's focus on the method set
	aspect of interfaces first.

Typically, we're introduced to interfaces with some contrived example.
	Let's go with the contrived example of writing some application
	where you're defining Animal datatypes, because that's a totally
	realistic situation that happens all the time. The Animal type
	will be an interface, and we'll define an Animal as being
	anything that can speak. This is a core concept in Go's type
	system; instead of designing our abstractions in teams of what kind
	of data our types can hold, we design our abstractions in terms of
	what actions our type can execute.
*/

// We start by defining our Animal interface:
type Animal interface {
	Speak() string
}

/* pretty simple:
we define an Animal as being any type that has a method named Speak.
	The Speak method takes no arguments and returns a string. Any type
	that defines this method is said to satisfy the Animal interface.
	There is no implements keyword in Go; whether or not a type satisfies
	an interface is determined automatically. Let's create a couple of types
	that satisfy this interface.
*/
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design Patterns!"
}

/* We now have four different types of animals:
A dog, a cat, a Llama, and a Java programmer. In our main() function, we can
	create a slice of Animals, and put one of each type into that slice, and
	see what each animal says. Let's do that now:
*/

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

/* Great, now you know how to use interfaces, and I don't need to talk about
them any more, right? Well, no, not really. Let's look at a few things that
	aren't very obvious to the buddy gopher.

The interface{} type

The interface{} type, the empty interface, is that source of much confusion
	the interface{} type is the interface that has no methods. Since there is
	no implements keyword, all types implement at least zero methods, and
	satisfying an interface is done automatically, all types satisfy the empty
	interface. That means that if you write a function that takes an
	interface{} value as a parameter, you can supply that function with any
	value. So, this function

func DoSomething(v interface{}){
	// ...
}

Will accept any parameter whatsoever.
*/
