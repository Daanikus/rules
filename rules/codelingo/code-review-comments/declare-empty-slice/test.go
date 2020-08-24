package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	personSlice := []Person{} //Issue
	strSlice := []string{}    //Issue
	intSlice1 := []int{}      //Issue
	intSlice2 := []int{1, 2}

	i1 := []interface{}{}   //ISSUE
	i2 := [][]interface{}{} //ISSUE

	s1 := []struct{}{}   //ISSUE
	s2 := [][]struct{}{} //ISSUE

	personSlice = append(personSlice, Person{FirstName: "John", LastName: "Snow", Age: 45})
	strSlice = append(strSlice, "test")
	intSlice1 = append(intSlice1, 1)

	fmt.Println(personSlice, strSlice, intSlice1, intSlice2)

	personSlice = []Person{}
	strSlice = []string{}
	intSlice1 = []int{}
	intSlice2 = []int{}

	fmt.Println(personSlice, strSlice, intSlice1, intSlice2)
	fmt.Println(i1, i2)
	fmt.Println(s1, s2)
}
