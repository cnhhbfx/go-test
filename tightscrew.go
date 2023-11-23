package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zeromicro/go-zero/core/threading"
)

type Screw struct {
	input  int
	output int
}

func (t *Screw) Tight(i int) {
	t.output = t.input*t.input + i
//	time.Sleep(time.Second)
}

func NewScrew() *Screw {
	s := &Screw{}
	s.input = rand.Intn(1000)
	return s
}

func main() {
	fmt.Println("prepare")
	count := 10000000
	is := make([]*Screw, 0, count)
	for i := 0; i < cap(is); i++ {
		is = append(is, NewScrew())
	}
	fmt.Println(count)
	for i := 0; i < 3; i++ {
		fmt.Printf("%v, ", is[i])
	}
	fmt.Printf("...\n")

	c := make(chan bool)
	threading.GoSafe(func() {
		fmt.Println("start")
		st := time.Now()
		g := threading.NewRoutineGroup()
		for i, v := range is {
			ri := i
			rv := v
			g.RunSafe(func() {
				rv.Tight(ri)
			})
		}
		g.Wait()
		el := time.Since(st)
		fmt.Printf("end\n%v\n", el)
		c <- true
	})

	<-c
	for i := 0; i < 3; i++ {
		fmt.Printf("%v, ", is[i])
	}
	fmt.Printf("...\n")
}
