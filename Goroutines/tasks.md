# 📌 Lista zadań: Współbieżność w Go

Poniżej znajduje się lista zadań, które krok po kroku pomogą Ci nauczyć się obsługi współbieżności w języku Go.

---

## **1. Goroutines – Pierwsze kroki**
✅ *Cel:* Uruchomić pierwszą goroutine i zobaczyć, jak działa równoległe wykonywanie kodu.

🔹 **Zadanie:**  
- Napisz funkcję, która wypisuje liczby od 1 do 10 z krótką pauzą między kolejnymi liczbami.  
- Uruchom tę funkcję zarówno w głównym wątku (`main`), jak i w goroutine (`go funkcja()`).  
- Zobacz, co się stanie i spróbuj rozwiązać problem przedwczesnego zakończenia programu.

---

## **2. Goroutines – Uruchamianie wielu goroutines**  
✅ *Cel:* Zrozumieć, jak działa jednoczesne uruchamianie wielu goroutines.

🔹 **Zadanie:**  
- Napisz funkcję, która przyjmuje liczbę i drukuje jej kwadrat.  
- Uruchom tę funkcję dla liczb od 1 do 10, każdą w osobnej goroutine.  
- Użyj `sync.WaitGroup`, aby upewnić się, że główny wątek poczeka na zakończenie wszystkich goroutines.

---

## **3. Kanały – Przesyłanie danych między goroutines**  
✅ *Cel:* Nauczyć się komunikacji między goroutines za pomocą kanałów.

🔹 **Zadanie:**  
- Stwórz kanał `chan int`.  
- Uruchom goroutine, która wysyła do kanału liczby od 1 do 5.  
- W głównym wątku odbieraj te liczby i wypisuj je na ekran.

---

## **4. Kanały buforowane**  
✅ *Cel:* Poznać różnicę między kanałami buforowanymi a niebuforowanymi.

🔹 **Zadanie:**  
- Stwórz kanał buforowany o pojemności 3.  
- Spróbuj wysłać do niego 5 wartości i zobacz, co się stanie.  
- Odbierz wartości i przetestuj różne pojemności bufora.

---

## **5. Select – Obsługa wielu kanałów**  
✅ *Cel:* Zrozumieć, jak `select` może odbierać dane z wielu kanałów.

🔹 **Zadanie:**  
- Stwórz dwa kanały (`chan string`).  
- Uruchom dwie goroutines – jedna co sekundę wysyła `"ping"`, druga co 1,5 sekundy `"pong"`.  
- W głównym wątku użyj `select`, aby odbierać wartości i wypisywać je na ekran.

---

## **6. Mutex – Synchronizacja dostępu do wspólnego zasobu**  
✅ *Cel:* Nauczyć się używać `sync.Mutex` do ochrony współdzielonych zasobów.

🔹 **Zadanie:**  
- Stwórz licznik (`counter int`).  
- Uruchom 10 goroutines, każda zwiększa licznik 1000 razy.  
- Użyj `sync.Mutex`, aby zapobiec problemom wyścigu (`race condition`).  
- Porównaj wynik bez użycia `Mutex` i z jego użyciem.

---

## **7. Once – Wykonanie kodu tylko raz**  
✅ *Cel:* Poznać `sync.Once` i jego zastosowanie.

🔹 **Zadanie:**  
- Stwórz funkcję, która wypisuje `"Inicjalizacja!"`.  
- Uruchom 10 goroutines, każda próbuje wywołać tę funkcję.  
- Użyj `sync.Once`, aby funkcja wykonała się tylko raz.

---

## **8. Context – Anulowanie goroutines**  
✅ *Cel:* Nauczyć się kontrolować czas życia goroutines za pomocą `context.Context`.

🔹 **Zadanie:**  
- Napisz goroutine, która co sekundę wypisuje `"Pracuję..."`.  
- W głównym wątku ustaw `context.WithCancel()`.  
- Po 5 sekundach anuluj goroutine i wypisz `"Zatrzymano!"`.

---

## **9. Tworzenie Workera – Wzorzec worker pool**  
✅ *Cel:* Implementacja wzorca worker pool do efektywnej pracy z wieloma zadaniami.

🔹 **Zadanie:**  
- Stwórz 5 workerów jako goroutines.  
- Przekazuj do nich zadania przez kanał.  
- Każdy worker przetwarza liczbę (np. oblicza jej kwadrat) i zwraca wynik przez inny kanał.  
- W głównym wątku odbieraj i wypisuj wyniki.

---

## **10. Benchmarking współbieżności w Go**  
✅ *Cel:* Zmierzyć wpływ współbieżności na wydajność programu.

🔹 **Zadanie:**  
- Napisz program, który wykonuje obliczenia sekwencyjnie i równocześnie (z goroutines).  
- Użyj `time.Now()` lub `testing.Benchmark()`, aby zmierzyć różnice w czasie wykonania.

---

To lista, która systematycznie wprowadzi Cię w świat współbieżności w Go. Powodzenia! 🚀