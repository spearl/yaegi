package interp

// Do not edit! File generated by ../test/gen_example.sh

import "github.com/containous/gi/export"

func Example_a1() {
	src := `
package main

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	println(a[1]) // 2
	for i, v := range a {
		println(v)
		if i == 3 {
			break
		}
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2
	// 1
	// 2
	// 3
	// 4
}

func Example_a2() {
	src := `
package main

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	a[1] = 5
	println(a[1]) // 2
	for i, v := range a {
		println(v)
		if i == 3 {
			break
		}
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 5
	// 1
	// 5
	// 3
	// 4
}

func Example_and() {
	src := `
package main

func main() {
	a, b := 1, 2

	if f1() && f2() {
		println(a, b)
	}
}

func f1() bool {
	println("f1")
	//return true
	return 0 == 0
}

func f2() bool {
	println("f2")
	//return false
	return 1 == 0
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// f1
	// f2
}

func Example_assign() {
	src := `
package main

func main() {
	a, b := 1, 2 // Multiple assign
	println(a, b)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 1 2
}

func Example_chan0() {
	src := `
package main

type Channel chan string

func send(c Channel) { c <- "ping" }

func main() {
	channel := make(Channel)
	go send(channel)
	msg := <-channel
	println(msg)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// ping
}

func Example_chan1() {
	src := `
package main

func send(c chan<- string) { c <- "ping" }

func main() {
	channel := make(chan string)
	go send(channel)
	msg := <-channel
	println(msg)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// ping
}

func Example_chan2() {
	src := `
package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// ping
}

func Example_const0() {
	src := `
package main

//const a = 1
const (
	a = iota
	b
)

func main() {
	println(a, b)
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_cont() {
	src := `
package main

func main() {
	for i := 0; i < 10; i++ {
		if i < 5 {
			continue
		}
		println(i)
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 5
	// 6
	// 7
	// 8
	// 9
}

func Example_cont0() {
	src := `
package main

func main() {
	i := 0
	for {
		if i > 10 {
			break
		}
		i++
		if i < 5 {
			continue
		}
		println(i)
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
}

func Example_cont1() {
	src := `
package main

func main() {
	i := 0
	for {
		if i > 10 {
			break
		}
		if i < 5 {
			i++
			continue
		}
		println(i)
		i++
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}

func Example_export0() {
	src := `
package main

func Test() {
	println("Hello from test")
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_fib() {
	src := `
package main

// Compute fibonacci numbers, no memoization
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	println(fib(35))
	//println(fib(10))
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_fib0() {
	src := `
package main

// Compute fibonacci numbers, no memoization
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	println(fib(4))
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 3
}

func Example_for0() {
	src := `
package main

func main() {
	i := 0
	//for ;i >= 0; i++ {
	for {
		if i > 5 {
			break
		}
		println(i)
		i++
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
}

func Example_for1() {
	src := `
package main

func main() {
	i := 0
	for i < 10 {
		if i > 4 {
			break
		}
		println(i)
		i++
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func Example_for2() {
	src := `
package main

func main() {
	for i := 2; ; i++ {
		println(i)
		if i > 3 {
			break
		}
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2
	// 3
	// 4
}

func Example_fun() {
	src := `
package main

func f(i int) int { return i + 15 }

func main() {
	println(f(4))
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 19
}

func Example_fun2() {
	src := `
package main

type Coord struct{ x, y int }

func f(c Coord) int { return c.x + c.y }

func main() {
	c := Coord{3, 4}
	println(f(c))
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 7
}

func Example_fun3() {
	src := `
package main

type Coord struct{ x, y int }

func f(i, j int, c Coord) int { return i*c.x + j*c.y }

func main() {
	c := Coord{3, 4}
	println(f(2, 3, c))
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 18
}

func Example_goroutine() {
	src := `
package main

func f() {
	println("in goroutine f")
}

func main() {
	go f()
	sleep(100)
	println("in main")
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_heap() {
	src := `
// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_if() {
	src := `
package main

func main() {
	if a := f(); a > 0 {
		println(a)
	}
}

func f() int { return 1 }
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 1
}

func Example_import0() {
	src := `
package main

import "fmt"

func main() {
	fmt.Println("Hello", 42)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// Hello 42
}

func Example_import1() {
	src := `
package main

import f "fmt"

func main() {
	f.Println("Hello", 42)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// Hello 42
}

func Example_import2() {
	src := `
package main

import . "fmt"

func main() {
	Println("Hello", 42)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// Hello 42
}

func Example_init0() {
	src := `
package main

func init() {
	println("Hello from init 1")
}

func init() {
	println("Hello from init 2")
}

func main() {
	println("Hello from main")
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// Hello from init 1
	// Hello from init 2
	// Hello from main
}

func Example_l2() {
	src := `
package main

func main() {
	for a := 0; a < 20000; a++ {
		if (a & 0x8ff) == 0x800 {
			println(a)
		}
	}
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_l3() {
	src := `
package main

func myprint(i int) { println(i) }

func main() {
	for a := 0; a < 20000000; a++ {
		if a&0x8ffff == 0x80000 {
			println(a)
			//myprint(a)
		}
	}
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_l4() {
	src := `
package main

func main()       { println(f(5)) }
func f(i int) int { return i + 1 }
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 6
}

func Example_map() {
	src := `
package main

type Dict map[string]string

func main() {
	dict := make(Dict)
	dict["truc"] = "machin"
	println(dict["truc"])
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// machin
}

func Example_map2() {
	src := `
package main

func main() {
	dict := make(map[string]string)
	dict["truc"] = "machin"
	println(dict["truc"])
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// machin
}

func Example_map3() {
	src := `
package main

func main() {
	dict := map[string]string{}
	dict["truc"] = "machin"
	println(dict["truc"])
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// machin
}

func Example_map4() {
	src := `
package main

func main() {
	dict := map[string]string{"bidule": "machin", "truc": "bidule"}
	dict["hello"] = "bonjour"
	println(dict["bidule"])
	println(dict["hello"])
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// machin
	// bonjour
}

func Example_math0() {
	src := `
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Cos(math.Pi))
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// -1
}

func Example_method() {
	src := `
package main

type Coord struct {
	x, y int
}

func (c Coord) dist() int { return c.x*c.x + c.y*c.y }

func main() {
	o := Coord{3, 4}
	println(o.dist())
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 25
}

func Example_method2() {
	src := `
package main

type Coord struct {
	x, y int
}

func (c Coord) dist() int { return c.x*c.x + c.y*c.y }

type Point struct {
	Coord
	z int
}

func main() {
	o := Point{Coord{3, 4}, 5}
	println(o.dist())
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 25
}

func Example_method3() {
	src := `
package main

type Coord struct {
	x, y int
}

func (c Coord) dist() int { return c.x*c.x + c.y*c.y }

type Point struct {
	Coord
	z int
}

func main() {
	o := Point{Coord{3, 4}, 5}
	println(o.Coord.dist())
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 25
}

func Example_method4() {
	src := `
package main

type Coord struct {
	x, y int
}

func (c Coord) dist() int { return c.x*c.x + c.y*c.y }

type Point struct {
	Coord
	z int
}

type Tpoint struct {
	t int
	Point
}

func main() {
	o := Tpoint{0, Point{Coord{3, 4}, 5}}
	println(o.dist())
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 25
}

func Example_ptr0() {
	src := `
package main

type myint int

func main() {
	var a myint = 2
	var b *myint = &a
	println(*b)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2
}

func Example_ptr1() {
	src := `
package main

func main() {
	var a int = 2
	b := &a
	println(*b)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2
}

func Example_ptr2() {
	src := `
package main

func f(i *int) {
	*i = *i + 3
}

func main() {
	var a int = 2
	f(&a)
	println(a)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 5
}

func Example_ptr3() {
	src := `
package main

func f(i *int) {
	*i++
}

func main() {
	var a int = 2
	f(&a)
	println(a)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 3
}

func Example_ret1() {
	src := `
package main

func f(i int) (o int) { o = i + 1; return }

func main() { println(f(4)) }`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_ret2() {
	src := `
package main

func main() {
	a, b := r2()
	println(a, b)
}

func r2() (int, int) { return 1, 2 }
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 1 2
}

func Example_ret3() {
	src := `
package main

import "fmt"

func main() {
	fmt.Println(r2())
}

func r2() (int, int) { return 1, 2 }
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 1 2
}

func Example_run0() {
	src := `
package main

func f() (int, int) { return 2, 3 }

func main() {
	println(f())
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2 3
}

func Example_run1() {
	src := `
package main

func f() (int, int) { return 2, 3 }

func g(i, j int) int { return i + j }

func main() {
	println(g(f()))
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 5
}

func Example_run10() {
	src := `
package main

func main() {
	func() { println("hello") }()
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// hello
}

func Example_run4() {
	src := `
package main

type fn func(int)

func f1(i int) { println("f1", i) }

func test(f fn, v int) { f(v) }

func main() { test(f1, 21) }
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// f1 21
}

func Example_run5() {
	src := `
package main

type fn func(int)

func test(f fn, v int) { f(v) }

func main() {
	f1 := func(i int) { println("f1", i) }
	test(f1, 21)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// f1 21
}

func Example_run6() {
	src := `
package main

type fn func(int)

func test(f fn, v int) { f(v) }

func main() {
	test(func(i int) { println("f1", i) }, 21)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// f1 21
}

func Example_run7() {
	src := `
package main

type fn func(int)

func test(f fn, v int) { f(v) }

func main() {
	a := 3
	test(func(i int) { println("f1", i, a) }, 21)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// f1 21 3
}

func Example_run8() {
	src := `
package main

func main() {
	a := 3
	f := func(i int) { println("f1", i, a) }
	f(21)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// f1 21 3
}

func Example_run9() {
	src := `
package main

func main() {
	a := 3
	f := func(i int) int { println("f1", i, a); return i + 1 }
	b := f(21)
	println(b)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// f1 21 3
	// 22
}

func Example_scope0() {
	src := `
package main

var a int = 1

func main() {
	println(a)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 1
}

func Example_scope1() {
	src := `
package main

func f(a int) int {
	return 2*a + 1
}

var b int = f(3)

func main() {
	println(b)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 7
}

func Example_scope2() {
	src := `
package main

var a int = 1

func f() { println(a) }

func main() {
	println(a)
	a := 2
	println(a)
	f()
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 1
	// 2
	// 1
}

func Example_scope3() {
	src := `
package main

func main() {
	a := 1
	if a := 2; a > 0 {
		println(a)
	}
	println(a)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2
	// 1
}

func Example_scope4() {
	src := `
package main

func main() {
	a := 1
	if a := 2; a > 0 {
		println(a)
	}
	{
		a := 3
		println(a)
	}
	println(a)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2
	// 3
	// 1
}

func Example_server() {
	src := `
package main

import (
	"fmt"
	"net/http"
)

var v string = "v1.0"

func main() {
	a := "hello "
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website! ", a, v)
	})

	http.ListenAndServe(":8080", nil)
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_sieve() {
	src := `
// A concurrent prime sieve

package main

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.

	for i := 0; i < 10; i++ {
		prime := <-ch
		println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2
	// 3
	// 5
	// 7
	// 11
	// 13
	// 17
	// 19
	// 23
	// 29
}

func Example_sleep() {
	src := `
package main

func main() {
	println("sleep")
	sleep(1000)
	println("bye")
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_str() {
	src := `
package main

func main() {
	println("hello world")
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// hello world
}

func Example_struct() {
	src := `
package main

type T struct {
	f int
	g int
}

func main() {
	a := T{7, 8}
	println(a.f, a.g)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 7 8
}

func Example_struct0() {
	src := `
package main

type T struct {
	f int
	g int
}

func main() {
	a := T{}
	println(a.f, a.g)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 0 0
}

func Example_struct0a() {
	src := `
package main

type T struct {
	f int
}

func main() {
	a := T{}
	println(a.f)
	a.f = 8
	println(a.f)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 0
	// 8
}

func Example_struct1() {
	src := `
package main

type T struct {
	f int
	g struct {
		h int
	}
}

func main() {
	a := T{}
	a.g.h = 3 + 2
	println("a.g.h", a.g.h)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// a.g.h 5
}

func Example_struct2() {
	src := `
package main

type T struct {
	f int
	g int
}

func main() {
	a := T{g: 8, f: 7}
	println(a.f, a.g)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 7 8
}

func Example_struct3() {
	src := `
package main

type T struct {
	f int
	g int
	h struct {
		k int
	}
}

func f(i int) int { return i + 3 }

func main() {
	a := T{}
	a.h.k = f(4)
	println(a.h.k)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 7
}

func Example_struct4() {
	src := `
package main

type T3 struct {
	k int
}

type T2 struct {
	h int
	T3
}

type T struct {
	f int
	g int
	T2
}

func f(i int) int { return i * i }

func main() {
	a := T{5, 7, T2{f(8), T3{9}}}
	println(a.f, a.g, a.h, a.k)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 5 7 64 9
}

func Example_struct5() {
	src := `
package main

type T struct {
	f int
	g int
}

func f(i int) int { return i * i }

func main() {
	a := T{7, f(4)}
	println(a.f, a.g)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 7 16
}

func Example_struct6() {
	src := `
package main

type T struct {
	f, g int
}

func main() {
	a := T{7, 8}
	println(a.f, a.g)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 7 8
}

func Example_switch() {
	src := `
package main

func main() {
	a := 3
	switch a {
	case 0:
		println(200)
	default:
		println(a)
	case 3:
		println(100)
	}
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 100
}

func Example_time0() {
	src := `
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_time1() {
	src := `
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	m := t.Minute()
	fmt.Println(t, m)
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_time2() {
	src := `
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	h, m, s := t.Clock()
	fmt.Println(h, m, s)
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_time3() {
	src := `
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Clock())
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_time4() {
	src := `
package main

import (
	"fmt"
	"time"
)

func main() {
	var m time.Month
	m = 9
	fmt.Println(m)
}`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

}

func Example_type0() {
	src := `
package main

type newInt int

func main() {
	var a newInt
	println(a)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 0
}

func Example_var() {
	src := `
package main

func main() {
	var a, b, c int
	println(a, b, c)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 0 0 0
}

func Example_var2() {
	src := `
package main

func main() {
	var a int = 2
	println(a)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2
}

func Example_var3() {
	src := `
package main

func main() {
	var a, b int = 2, 3
	println(a, b)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2 3
}

func Example_var4() {
	src := `
package main

func main() {
	var a, b = 2, 3
	println(a, b)
}
`
	i := NewInterpreter(Opt{Entry: "main"})
	i.ImportBin(export.Pkg)
	i.Eval(src)

	// Output:
	// 2 3
}
