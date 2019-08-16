package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func main() {
	varInitExample()
}

func varInitExample() {
	var i int
	fmt.Println(i)

	var y int = 5
	fmt.Println(y)

	a := 10
	fmt.Println(a)
}

func conditionExample() {
	x := 7
	if x < 8 {
		fmt.Println("less than 8")
	} else if x > 8 {
		fmt.Println("greater than 8")
	} else {
		fmt.Println("nothing")
	}
}

func arrayExample() {
	arr := [5]int{}
	fmt.Println(arr)
	arr[0] = 5
	fmt.Println(arr)
}

func slicesExample() {
	slices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slices)
	slices[0] = 5
	fmt.Println(slices)
}

func mapExample() {
	m := make(map[string]int)
	m["hello"] = 1
	m["world"] = 2
	fmt.Println(m)

	delete(m, "world")
	fmt.Println(m)
}

func forLoopExample() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func whileLoopExample() {
	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
	}
}

func loopThroughSlices() {
	sl := []string{"h", "e", "l", "l", "o"}
	for index, value := range sl {
		fmt.Println("index", index, "value", value)
	}
}

func loopThroughMap() {
	m := map[string]int{"h": 1, "e": 2, "l": 3}
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}
}

func sumExample(x int, y int) int {
	return x + y
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("negative number is not allowed")
	}
	return math.Sqrt(num), nil
}

func sqrtWithNamedReturnValues(num float64) (result float64, err error) {
	if num < 0 {
		result = 0
		err = errors.New("negative number is not allowed")
		return
	}
	result = math.Sqrt(num)
	err = nil
	return
}

func pointerExample() {
	i := 7
	fmt.Println(i)

	fmt.Println(&i)

	fmt.Println(*&i)
}

func advancedPointerExample() {
	i := 7
	inc(i)
	fmt.Println(i)

	incPointer(&i)
	fmt.Println(i)
}

func inc(i int) {
	i++
}

func incPointer(i *int) {
	*i++
}

type Person struct {
	name string
	age  int
}

func structExample() {
	person := Person{name: "nhat", age: 23}
	fmt.Println(person)
	fmt.Println(person.name)
}

func structWithMethod() {
	p := Person{name: "nhat", age: 23}
	fmt.Println(p)
	p.changeNameWithoutPointer("hello")
	fmt.Println(p)
	p.changeName("codelink")
	fmt.Println(p)
}

func (p *Person) changeName(newName string) {
	p.name = newName
}

func (p Person) changeNameWithoutPointer(newName string) {
	p.name = newName
}

func goRoutineExample() {
	// main is one goroutine
	// go count("sheep") // now we have 2 go routines
	countInfinity("sheep")
	countInfinity("fish")

	// go count("fish") // now we have 3 go routines, but nothing happens because the main have no line of code to be executed --> program terminates
}

func countInfinity(thing string) {
	for {
		fmt.Println(thing)
		time.Sleep(500 * time.Millisecond)
	}
}

func goroutinesCommunicationExample() {
	// channel is a pipe, send or receive a message, chan have its own type
	// sending and receiving blocking operation
	c := make(chan string)

	go countWithChan("sheep", c)

	// Case 1: 1 message output
	msgReceived := <-c // blocking, wait for message to be sent
	fmt.Println(msgReceived)

	// Case 2: 5 messages output, but deadlock, because msgReceived block the system and infinitely wait for channel to sent message
	// for {
	// 	msgReceived := <-c
	// 	fmt.Println(msgReceived)
	// }

	// Case 3: Leverage the range to output all messages then stop
	// for msgReceived := range c {
	// 	fmt.Println(msgReceived)
	// }

}

func countWithChan(msgSent string, c chan string) {
	for i := 0; i < 6; i++ {
		c <- msgSent // blocking, wait for message to be ready to receive
		time.Sleep(500 * time.Millisecond)
	}

	// close channel, only close in sender, if close in receiver --> deadlock, because channel is closed but sender keep sending
	// close(c)
}

func channelWithBufferedExample() {
	// Init code will be deadlock
	c := make(chan string)

	c <- "hello"

	msg := <-c
	fmt.Println(msg)

	// Solution

	// c := make(chan string, 2)
	// c <- "hello"

	// msg := <-c
	// fmt.Println(msg)

}

func channelWithSelectStatement() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}

	// Solution

	// for {
	// 	select { // execute first non-blocking channel
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case ms2 := <-c2:
	// 		fmt.Println(ms2)
	// 	}
	// }
}

type Animal interface {
	introduce()
}

type Dog struct {
	name string
}

func (d *Dog) introduce() {
	// if d == nil {
	// 	fmt.Println("<nil>")
	// 	return
	// }
	fmt.Println("My name is", d.name)
}

type Cat struct {
	name string
}

func (c *Cat) introduce() {
	// if c == nil {
	// 	fmt.Println("<nil>")
	// 	return
	// }
	fmt.Println("My name is", c.name)
}

func interfaceExample() {
	var animal Animal = &Cat{name: "Mèo"}
	animal.introduce()
	animal = &Dog{name: "Chó"}
	animal.introduce()
}

func interfaceValuesExample() {
	var animal Animal = &Cat{name: "C"}
	describe(animal)
	animal = &Dog{name: "D"}
	describe(animal)
}

func describe(t interface{}) {
	fmt.Printf("(Type: %T, Value: %v)\n", t, t)
}

func interfaceWithUnderlyingNilValue() {
	var animal Animal
	var dog *Dog
	animal = dog
	describe(animal)
	animal.introduce()
}

func interfaceWithNilValue() {
	var animal Animal
	describe(animal)
	animal.introduce()
}

func emptyInterfaceExample() {
	var emptyInterface interface{}
	emptyInterface = 1
	describe(emptyInterface)
	emptyInterface = "hello"
	describe(emptyInterface)
	i := 5
	describe(i)
}

type MyError struct {
	errString string
}

func (mE *MyError) Error() string {
	return mE.errString
}

func sqrtWithError(f float64) (float64, error) {
	if f < 0 {
		return 0, &MyError{errString: "input number is negative"}
	}
	return math.Sqrt(f), nil
}

func errorExample() {
	i := 4.0
	if value, err := sqrtWithError(i); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
