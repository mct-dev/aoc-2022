package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ByTotal []int

func (a ByTotal) Len() int           { return len(a) }
func (a ByTotal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTotal) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	bytes, err := os.ReadFile("1/input.txt")
	if err != nil {
		panic(err)
	}

	st := string(bytes)

	chunks := strings.Split(st, "\n\n")

	largestTotal := -1
	largestIndex := -1
	totals := make([]int, len(chunks))

	for index, chunk := range chunks {
		nums := strings.Split(chunk, "\n")
		total := 0
		for _, num := range nums {
			if len(num) == 0 {
				continue
			}

			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			total += val
		}

		if total > largestTotal {
			largestTotal = total
			largestIndex = index
		}

		totals = append(totals, total)
	}

	sort.Sort(ByTotal(totals))

	fmt.Println("largest index:", largestIndex)
	fmt.Println("total:", largestTotal)
	fmt.Println("top 3:", totals[len(totals)-3:])
	fmt.Println("top 3 sum:", totals[len(totals)-3]+totals[len(totals)-2]+totals[len(totals)-1])
}
