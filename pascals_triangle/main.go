package main

import (
	"errors"
	"flag"
	"fmt"
)

// Triangle is an array of integer arrays
type Triangle [][]int

func initializeTriangle() Triangle {
	triangle := make([][]int, 0)
	triangle = append(triangle, []int{1})
	triangle = append(triangle, []int{1, 1})
	return triangle
}

func (triangle Triangle) addRow() Triangle {
	previous, newRow := triangle[len(triangle)-1], []int{1}

	for i := 0; i < len(previous)-1; i++ {
		newRow = append(newRow, previous[i]+previous[i+1])
	}

	newRow = append(newRow, 1)
	triangle = append(triangle, newRow)
	return triangle
}

func newTriangle(rows int) (Triangle, error) {
	var triangle = initializeTriangle()

	if rows < 0 {
		return make([][]int, 0), errors.New("negative rows are not supported")
	} else if rows == 1 {
		return [][]int{[]int{1}}, nil
	} else if rows == 2 {
		return triangle, nil
	}

	for i := 2; i < rows; i++ {
		triangle = triangle.addRow()
	}

	return triangle, nil
}

func (triangle Triangle) print() {
	for _, val := range triangle {
		fmt.Println(val)
	}
}

func main() {
	row := flag.Int("row", 5, "the row of values of Pascal's Triangle")
	flag.Parse()

	triangle, err := newTriangle(*row)

	if err != nil {
		fmt.Println(err)
	}

	triangle.print()
	fmt.Println("")
	fmt.Printf("The row you want is %v\n", triangle[*row-1])
}
