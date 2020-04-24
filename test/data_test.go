package test

import (
	"fmt"
	"testing"
)

const (
	r = 1 << iota
	w
	e
)
const (
	Monday = 2 * iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func TestWeek(t *testing.T) {
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(Sunday)

	fmt.Println(r)
	fmt.Println(w)
	fmt.Println(e)
}
func TestString(t *testing.T) {
	var str string
	fmt.Println(str)
}
func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1])
	arr1 := [3]int{}
	t.Log(arr1[1])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for _, v := range arr3 {
		t.Log(v)
	}
}
func TestSliceShareMemory(t *testing.T) {
	var year [12]int
	for i := 1; i <= 12; i++ {
		year[i-1] = i
	}
	sli := year[6:11]
	t.Log(&year[6], &sli[0])
	sli = append(append(sli, 18), 24)
	t.Log(year)
	t.Log(sli)
	t.Log(&year[6], &sli[0])

	//slice := year[3:6]
	//t.Log(len(slice),cap(slice))
	//t.Log(year)
	//t.Log(slice)
	//sli := year[4:7]
	//sli[0]=3014
	//t.Log(year)
	//t.Log(slice)
	//t.Log(sli)
}
