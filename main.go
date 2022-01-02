package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func isCorrect(w, h, max int) error {
	if w <= 0 {
		return errors.New("Width needs to be more than 0")
	}
	if h <= 0 {
		return errors.New("Height needs to be more than 0")
	}
	if max < (w * h) {
		return errors.New("Max needs more w*h for uniq numbers")
	}
	return nil
}

func RandomMatrix(w, h, max int) ([][]int, error) {
	if err := isCorrect(w, h, max); err != nil {
		return nil, err
	}
	nums := UniqNumber(w, h, max)
	var arr [][]int
	for i := 0; i < w; i++ {
		arr = append(arr, []int{})
		for j := 0; j < h; j++ {
			arr[i] = append(arr[i], nums[0])
			nums = nums[1:]
		}
	}
	return arr, nil
}

func UniqNumber(w, h, max int) []int {
	rand.Seed(time.Now().UnixNano())
	un := map[int]interface{}{}
	for len(un) < (w * h) {
		for i := 0; i < ((w * h) - len(un)); i++ {
			num := rand.Intn(max)
			un[num] = nil
		}
	}
	arr := make([]int, 0, len(un))
	for i, _ := range un {
		arr = append(arr, i)
	}
	return arr
}

var w, h, max int

func init() {
	flag.IntVar(&w, "w", 1, "Width matrix")
	flag.IntVar(&h, "h", 1, "Hight matrix")
	flag.IntVar(&max, "max", 10, "Max range matrix")
}

func main() {
	flag.Parse()
	arr, err := RandomMatrix(w, h, max)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\n", arr[i])
	}
}
