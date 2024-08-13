/*
Generics: create type-agnostic functions and
structs, making code adaptable to
different data types.

# Reduced code duplication

improved type safety

enhanced readbility
*/
package generics

import "fmt"

type Sortable interface {
	~int | ~float64 | ~string
}

func GenericSort[T Sortable](slice []T) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func main() {
	intSlice := []int{5, 1, 2, 9, 8}
	floatSlice := []float64{2.1, 1.3, 5.5, 4.7, 8.9}
	stringSlice := []string{"apple", "orange", "banana", "pear"}

	GenericSort(intSlice)
	GenericSort(floatSlice)
	GenericSort(stringSlice)

	fmt.Println("Sorted intSlice: ", intSlice)
	fmt.Println("Sorted floatSlice: ", floatSlice)
	fmt.Println("Sorted stringSlice: ", stringSlice)
}
