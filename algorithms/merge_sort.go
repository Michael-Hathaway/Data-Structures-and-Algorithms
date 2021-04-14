// Mergesort algorithm implemented in go

package main

import "fmt"

type Sortable interface {
	SortableValue() int
}

func MergeSort(array []Sortable) []Sortable {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2
	left := MergeSort(array[:mid])
	right := MergeSort(array[mid:])

	i, j, k := 0, 0, 0
	newArray := make([]Sortable, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i].SortableValue() < right[j].SortableValue() {
			newArray[k] = left[i]
			i, k = i+1, k+1
		} else {
			newArray[k] = right[j]
			j, k = j+1, k+1
		}
	}

	for i < len(left) {
		newArray[k] = left[i]
		i, k = i+1, k+1
	}

	for j < len(right) {
		newArray[k] = right[j]
		j, k = j+1, k+1
	}

	return newArray
}

// Person implements the Sortable interface
type Person struct {
	Name string
	Age  int
}

func (p Person) SortableValue() int {
	return p.Age
}

func CreatePerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func main() {
	michael := CreatePerson("Michael", 23)
	jerry := CreatePerson("Jerry", 45)
	alex := CreatePerson("Alex", 15)
	sarah := CreatePerson("Sarah", 16)
	people := []Sortable{michael, jerry, alex, sarah}

	people = MergeSort(people)

	fmt.Println(people)
}
