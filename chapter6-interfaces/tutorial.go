package main

import "fmt"

// type S
type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

// type R
type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	// switch t := p.(type) {
	// case *S:
	// case *R:
	// default:

	// }
	fmt.Println(p.Get())
	p.Put(10)
	fmt.Println(p.Get())
}

func g(something interface{}) int {
	return something.(I).Get()
}

// SORT
type Sorter interface {
	Len() int           // len() as a method.
	Less(i, j int) bool // p[j] < p[i] as a method.
	Swap(i, j int)      // p[i], p[j] = p[j], p[i] as a method.
}

type Xi []int
type Xs []string

// implement Sorter interface for []int
func (p Xi) Len() int           { return len(p) }
func (p Xi) Less(i, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// implement Sorter interface for []string
func (p Xs) Len() int           { return len(p) }
func (p Xs) Less(i, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}

		}
	}
}

func main() {
	var x = new(S)
	f(x)

	fmt.Println(g(x))

	// var i int
	// fmt.Println(g(i))

	fmt.Println("==== SORTER =====")
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}

	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
}
