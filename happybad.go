package main

import "fmt"

import "github.com/zeromicro/go-zero/core/threading"

func Happy() {
	fmt.Println("Happy")
}

func Bad() {
	fmt.Println("Bad")
}

func main() {
	fa := make([]func(), 0, 2)
	fa = append(fa, Happy)
	fa = append(fa, Bad)

	g := threading.NewRoutineGroup()
	for _, v := range fa {
		rv := v
		g.RunSafe(func() {
			rv()
		})
	}
	g.Wait()
}
