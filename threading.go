package main

import (
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/core/threading"
)

type Cat struct {
	X            int
	Y            int
	PlusResult   int
	DivideResult int
}

func (c *Cat) Plus(v int) {
	c.PlusResult = c.X + c.Y + v
}

func (c *Cat) Divide() {
	c.DivideResult = c.X / c.Y
}

func main() {
	ticker := time.NewTicker(time.Second * 15)

	a := &Cat{}
	fmt.Println(a)

	a.X = 6
	a.Y = 2
	fmt.Println(a)

	b := 7

	// 串行
	//a.Plus(b)
	//a.Divide()
	//fmt.Println(a)

	// 并行
	g := threading.NewRoutineGroup()
	g.Run(func() {
		a.Plus(b)
	})
	g.RunSafe(func() {
		a.Divide()
	})
	g.Wait()
	fmt.Println(a)

	_ = stat.CpuUsage()
	for {
		select {
		case <-ticker.C:
			fmt.Println("sleep a while")
		}
	}
}
