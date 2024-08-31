package main

import "fmt"

// Fonction de tri par insertion
func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1

        // Déplace les éléments de arr[0..i-1] qui sont plus grands que la clé
        // d'une position vers la droite pour faire de la place pour la clé
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
}

func main() {
    arr := []int{12, 11, 13, 5, 6}
    fmt.Println("Tableau avant tri :", arr)

    insertionSort(arr)

    fmt.Println("Tableau après tri :", arr)
}
