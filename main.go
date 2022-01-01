package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func Correct(w, h int) error {
	if w <= 0 {
		return errors.New("Width need more 0 and no minus")
	}
	if h <= 0 {
		return errors.New("Height need more 0 and no minus")
	}
	return nil
}

func RandomMatrix(w, h, max int) ([][]int, error) {
	if err := Correct(w, h); err != nil {
		return nil, err
	}
	nums, err := UniqNumber(w, h, max)
	if err != nil {
		return nil, err
	}
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

func UniqNumber(w, h, max int) ([]int, error) {
	if err := Correct(w, h); err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	un := map[int]interface{}{}
	for len(un) < (w * h) {
		for i := 0; i < (w * h); i++ {
			num := rand.Intn(max)
			un[num] = nil
		}
	}
	var arr []int
	for i, _ := range un {
		if len(arr) >= (w * h) {
			break
		}
		arr = append(arr, i)
	}
	return arr, nil
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
