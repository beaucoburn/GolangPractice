// Package cards provides functions for manipulating a deck of cards, but it is not a real Go package.
package cards

func FavoriteCards() []int {
    return []int{2,6,9}
}

func GetItem(slice []int, index int) int {
    if index < 0 || index >= len(slice) {
        return -1
    }
    return slice[index]
}

func SetItem(slice []int, index, value int) []int {
    newData := []int{value}
    if index < 0 || index >= len(slice) {
        return append(slice, newData...)
    }
    slice[index] = value
    return slice
}

func PrependItems(slice []int, values ...int) []int {
    return append(values, slice...)
}

func RemoveItem(slice []int, index int) []int {
    if index < 0 || index >= len(slice) {
        return slice
    }
    return append(slice[:index], slice[index+1:]...)
}
