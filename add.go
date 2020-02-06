package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) (int, error) {
	if x > 10 || y > 10 {
		return 0, fmt.Errorf("%d, or %d was greater than 10", x, y)
	}
	return x + y, nil
}

type score struct {
	home	int
	away	int
}

func (s score) scoreDiff() string {
	if s.home > s.away {
		return fmt.Sprintf("home wins by %d", s.home - s.away)
	} else {
		return fmt.Sprintf("away wins by %d", s.away - s.home)
	}
}

func main() {
	a, err := add(rand.Intn(10), rand.Intn(10))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Output from add(): %d\n", a)
	}

	var s score
	s = score{
		home: 31,
		away: 20,
	}
	fmt.Printf("Output from scoreDiff(): %s\n", s.scoreDiff())
}