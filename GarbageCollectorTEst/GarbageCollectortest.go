package main

import (
	"fmt"
	"runtime"
	"time"
)

// Funkcja wyświetlająca statystyki pamięci
func printMemoryStats(message string) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("\n%s\n", message)
	fmt.Printf("Zaalokowane bajty: %v\n", memStats.Alloc)
	fmt.Printf("Liczba obiektów na stercie (HeapObjects): %v\n", memStats.HeapObjects)
}

func allocateMemory() {
	// Tworzymy dużą tablicę, która zajmuje dużo pamięci
	var memoryHog = make([]int, 1e6) // 1 milion elementów
	fmt.Println("Zaalokowano pamięć:", len(memoryHog), "elementów")

	// Symulujemy pracę programu przez chwilę
	time.Sleep(2 * time.Second)

	// Po zakończeniu funkcji, tablica przestaje być dostępna
	// i staje się kandydatem do usunięcia przez GC
}

func main() {
	fmt.Println("Start programu")

	// Wyświetlamy początkowe statystyki pamięci
	printMemoryStats("Statystyki pamięci przed alokacją:")

	// Alokujemy pamięć kilka razy
	for i := 0; i < 5; i++ {
		allocateMemory()
		fmt.Println("Zakończono iterację:", i+1)
	}

	// Wyświetlamy statystyki pamięci po alokacji
	printMemoryStats("Statystyki pamięci po alokacji:")

	// Ręcznie wymuszamy uruchomienie Garbage Collectora
	fmt.Println("\nRęczne wywołanie GC...")
	runtime.GC()

	// Czekamy chwilę, aby GC zdążył zrobić swoje
	time.Sleep(2 * time.Second)

	// Wyświetlamy statystyki pamięci po działaniu GC
	printMemoryStats("Statystyki pamięci po działaniu GC:")

	fmt.Println("Koniec programu")
}
