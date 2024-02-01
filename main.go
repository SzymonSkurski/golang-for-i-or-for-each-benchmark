package main

import "fmt"

type colorRGB struct {
	red   uint8
	green uint8
	blue  uint8
}

func (c colorRGB) clone() colorRGB {
	return colorRGB{
		red:   c.red,
		green: c.green,
		blue:  c.blue,
	}
}

func main() {
	myColors := getMyColors()

	fmt.Println(cloneForI(myColors))
	fmt.Println(cloneForEach(myColors))
}

func cloneForI(s map[int]colorRGB) map[int]colorRGB {
	if s == nil {
		return nil
	}

	c := make(map[int]colorRGB, len(s))

	for i := 0; i < len(s); i++ {
		c[i] = s[i].clone()
	}

	return c
}

func cloneForEach(s map[int]colorRGB) map[int]colorRGB {
	if s == nil {
		return nil
	}

	c := make(map[int]colorRGB, len(s))

	for key, val := range s {
		c[key] = val.clone()
	}

	return c
}

func getMyColors() map[int]colorRGB {
	return map[int]colorRGB{
		1: {
			red:   255,
			green: 0,
			blue:  0,
		},
		2: {
			red:   0,
			green: 255,
			blue:  0,
		},
		3: {
			red:   0,
			green: 0,
			blue:  255,
		},
		4: {
			red:   0,
			green: 0,
			blue:  0,
		},
		5: {
			red:   255,
			green: 255,
			blue:  255,
		},
	}
}
