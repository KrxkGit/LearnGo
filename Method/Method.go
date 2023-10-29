package Method

import "fmt"

type Number struct {
	sum int
}

func (n Number) Add(i int) Number {
	n.sum += i
	//res = n
	return n
}

func (np *Number) AddbyPtr(i int) {
	np.sum += i
}

func Test() {
	a := Number{1}
	b := Number{1}

	a.Add(1)
	fmt.Printf("a: %d\n", a.sum)
	b.AddbyPtr(1)
	fmt.Printf("b: %d\n", b.sum)

	fmt.Println()

	c := a
	c.Add(1)
	fmt.Printf("a: %d\n", a.sum)
	fmt.Printf("c: %d\n", c.sum)

	e := a.Add(1)
	fmt.Printf("e: %d\n", e.sum)

	fmt.Println()

	d := b
	d.AddbyPtr(1)
	fmt.Printf("b: %d\n", b.sum)
	fmt.Printf("d: %d\n", d.sum)

	fmt.Println()
	f := &b
	f.AddbyPtr(1)
	fmt.Printf("b: %d\n", b.sum)
	fmt.Printf("f: %d\n", f.sum)

}
