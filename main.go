package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//number 1
	fmt.Println("Swap", swapNumbers())
	fmt.Println("=====================")

	//number 2
	var list = []int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6}
	findNumbers(list)
	fmt.Println("=====================")

	//number 3
	mergeString("saya", "kamu")
	fmt.Println("=====================")

	//number 4
	morgan("JACK", "DANIEL")
	fmt.Println("=====================")

	//number 5
	var a = [][]int{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}
	arrayManipulation(len(a), a)
}

func swapNumbers() []int {
	a, b := 3, 5
	b, a = a, b
	return []int{a, b}
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func findNumbers(list []int) {
	m := make(map[int]int)
	for _, v := range list {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}

	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	fmt.Println("number", "-->", "total")
	for _, v := range pl {
		fmt.Println(v.Key, "-->", v.Value)
	}
}

func mergeString(a string, b string) {
	arrayA := strings.Split(a, "")
	arrayB := strings.Split(b, "")

	var first []string
	var second []string

	if len(arrayB) > len(arrayA) {
		first = arrayA
		second = arrayB
	} else {
		first = arrayA
		second = arrayB
	}

	result := ""
	for i, a := range first {
		result += a

		if i < len(second) {
			result += second[i]
		}
		if i == len(first)-1 && len(second) > len(first) {
			for j := i + 1; j < len(second); j++ {
				result += second[j]
			}
		}

	}
	fmt.Println(result)
}

func morgan(a string, b string) {
	result := []string{}
	needed := len(a) + len(b)
	found := 0
	a += "z"
	b += "z"
	for found < needed {
		if a < b {
			newCount := 1
			for a[newCount] == a[0] {
				newCount += 1
			}
			result = append(result, a[:newCount])
			found += newCount
			a = a[newCount:]
		} else {
			newCount := 1
			for b[newCount] == b[0] {
				newCount += 1
			}
			result = append(result, b[:newCount])
			found += newCount
			b = b[newCount:]
		}
		if len(a) == 1 {
			result = append(result, b[:len(b)-1])
			break
		} else if len(b) == 1 {
			result = append(result, a[:len(a)-1])
			break
		}
	}
	fmt.Println(strings.Join(result[:], ""))
}

func arrayManipulation(n int, queries [][]int) {
	outputArray := make([]int, n+4)
	for i := 0; i < len(queries); i++ {
		a := queries[i][0]
		b := queries[i][1]
		k := queries[i][2]
		outputArray[a] += k
		outputArray[b+1] -= k
	}
	getMax(outputArray)
}

func getMax(inputArray []int) int {
	max := 0
	sum := 0
	for i := 0; i < len(inputArray); i++ {
		sum += inputArray[i]
		max = maxMath(max, sum)
	}
	fmt.Println(max)
	return max
}

func maxMath(a, b int) int {
	if a > b {
		return a
	}
	return b
}
