# Zadania ze wskaźnikami w Go

## Podstawowe operacje

1. **Inicjalizacja wskaźnika**  
   Zadeklaruj wskaźnik do zmiennej typu `int`, przypisz mu adres innej zmiennej i wyświetl jego wartość oraz adres.

2. **Dereferencja wskaźnika**  
   Utwórz zmienną typu `float64`, przypisz jej wskaźnik i zmodyfikuj wartość poprzez dereferencję.

3. **Sprawdzenie wskaźnika `nil`**  
   Zainicjalizuj wskaźnik bez przypisania wartości i sprawdź, czy jest `nil`.

4. **Użycie operatora `new`**  
   Użyj `new` do alokacji pamięci dla zmiennej typu `string` i przypisz do niej wartość.

5. **Zamiana wartości przez wskaźnik**  
   Napisz funkcję, która zamienia wartość zmiennej typu `int` przekazanej jako wskaźnik.

## Wskaźniki i funkcje

6. **Przekazanie wskaźnika do funkcji**  
   Napisz funkcję, która zwiększa liczbę o 10, przyjmując wskaźnik na `int`.

7. **Zamiana dwóch zmiennych**  
   Napisz funkcję `swap`, która zamienia wartości dwóch liczb całkowitych przy użyciu wskaźników.

8. **Zwiększanie wartości wskaźnika w funkcji**  
   Utwórz funkcję, która inkrementuje wartość zmiennej przekazanej przez wskaźnik.

9. **Funkcja zwracająca wskaźnik**  
   Napisz funkcję, która zwraca wskaźnik do nowo utworzonej zmiennej typu `bool`.

10. **Zwracanie większej liczby przez wskaźnik**  
    Napisz funkcję, która przyjmuje dwa wskaźniki na liczby całkowite i zwraca wskaźnik na większą z nich.

## Wskaźniki i tablice

11. **Tablica i wskaźniki**  
    Utwórz tablicę 5 liczb całkowitych, zainicjalizuj wskaźnik na jej pierwszy element i iteruj po elementach za pomocą wskaźnika.

12. **Modyfikacja elementów tablicy przez wskaźniki**  
    Napisz funkcję, która przyjmuje wskaźnik do tablicy i zwiększa każdy jej element o 1.

13. **Wskaźnik do `slice`**  
    Utwórz `slice` liczb całkowitych i przekaż go do funkcji jako wskaźnik, aby zmodyfikować jego zawartość.

14. **Wskaźnik na pierwszy element tablicy**  
    Napisz program, który przechowuje adres pierwszego elementu tablicy i używa wskaźnika do iteracji po elementach.

15. **Wskaźnik do wskaźnika**  
    Napisz program, który używa wskaźnika do wskaźnika (`**int`) do modyfikacji wartości zmiennej.

## Struktury i wskaźniki

16. **Wskaźnik na strukturę**  
    Utwórz strukturę `Person` z polami `Name` i `Age`, utwórz jej instancję i użyj wskaźnika do modyfikacji pól.

17. **Przekazywanie struktury przez wskaźnik do funkcji**  
    Napisz funkcję, która modyfikuje pole `Age` struktury `Person`, przyjmując wskaźnik jako argument.

18. **Wskaźnik w strukturze**  
    Dodaj do struktury `Book` pole `*Author`, gdzie `Author` to inna struktura, i przypisz jej wartość.

19. **Metody odbiorcze ze wskaźnikami**  
    Utwórz metodę odbiorczą dla struktury `Counter`, która zwiększa `Value` o 1 i używa wskaźnika do modyfikacji.

20. **Dynamiczna alokacja struktury**  
    Użyj `new`, aby dynamicznie utworzyć strukturę `Car` i przypisać jej wartości.
