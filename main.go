package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Tree struct {
	Open  bool
	Close bool
	Child []RootNode
}

type RootNode struct {
	Key   string
	Value Value
}

type Value interface{}

type Object struct {
	Type     string
	Children []Property
}

type Property struct {
	Type  string
	Key   string
	Value Value
}

type Array struct {
	Type     string
	Children []Value
}

type Literal struct {
	Type  string
	Value string
}

func main() {
	flag.Parse()

	f, _ := os.Open(flag.Args()[0])
	defer f.Close()

	buf := bufio.NewReader(f)
	loop := true

	root := Tree{Open: false}

	for loop {
		r, _, err := buf.ReadRune()

		if err != nil {
			fmt.Println("read rune error")
			loop = false
		}

		fmt.Println(r)
		switch r {
		case 123: // bracket open
			if !root.Open {
				root.Open = true
				fmt.Println("open")
			}
			break
		case 10: // \n
			break
		case 32: // space
			break
		case 34: // "

		}
	}
}
