package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/samber/mo"
)

var anError = fmt.Errorf("error")

func main() {
	result := mayFail()
	// Error condition
	if result.IsError() {
		_, r := result.Get()
		fmt.Println("Oh no,", r.Error())
		return
	}
	// Success
	r, _ := result.Get()
	fmt.Println("Good job ", r)
}

func mayFail() mo.Result[string] {
	// Pseudo-random effectual nonsense
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i := r1.Intn(10)
	m := i % 2

	if m == 0 {
		// Success
		return mo.Ok("we have randomly succeeded")
	}
	// Failure
	return mo.Err[string](
		fmt.Errorf("totally predictable failure scenario"),
	)
}
