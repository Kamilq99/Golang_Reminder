# Ćwiczenia ze strukturami w Golang

## 1. Struktura Osoby
**Zadanie:**  
Stwórz strukturę `Person` z polami `Name` (string) i `Age` (int). Utwórz kilka instancji i wypisz je na ekran.

**Rozszerzenie:**  
Napisz funkcję, która zwróci osobę o największym wieku z listy.

---

## 2. Metody dla struktury
**Zadanie:**  
Dodaj metodę `Greet()` do struktury `Person`, która wypisze powitanie, np. _"Cześć, mam na imię Jan!"_.

**Rozszerzenie:**  
Dodaj metodę `IsAdult()`, która zwróci `true`, jeśli osoba ma 18 lat lub więcej.

---

## 3. Struktura Samochodu
**Zadanie:**  
Stwórz strukturę `Car` z polami `Brand`, `Model` i `Year`.

**Rozszerzenie:**  
Dodaj metodę `Age()`, która zwróci wiek samochodu na podstawie aktualnego roku.

---

## 4. Wskaźniki w strukturach
**Zadanie:**  
Stwórz strukturę `BankAccount` z polami `Owner` (string) i `Balance` (float64).  
Napisz metodę `Deposit(amount float64)`, która zwiększy saldo konta.

**Rozszerzenie:**  
Dodaj metodę `Withdraw(amount float64)`, która pozwoli na wypłatę pieniędzy (nie pozwól na zejście poniżej zera!).

---

## 5. Struktury zagnieżdżone
**Zadanie:**  
Stwórz strukturę `Address` z polami `City` i `Street`, a następnie użyj jej w strukturze `Person`, aby każda osoba miała adres.

**Rozszerzenie:**  
Napisz funkcję `ChangeAddress(p *Person, city string, street string)`, która zmieni adres osoby.

---

## 6. Lista studentów
**Zadanie:**  
Stwórz strukturę `Student` z polami `Name` i `Grades` (slice int).  
Dodaj metodę `AverageGrade()`, która zwróci średnią ocen studenta.

**Rozszerzenie:**  
Napisz funkcję, która zwróci listę studentów z oceną powyżej 4.5.

---

## 7. Książka i biblioteka
**Zadanie:**  
Stwórz strukturę `Book` z polami `Title`, `Author` i `Year`.  
Następnie stwórz strukturę `Library`, która zawiera listę książek.

**Rozszerzenie:**  
Dodaj metodę `AddBook(b Book)`, która doda książkę do biblioteki oraz `FindBooksByAuthor(author string)`, która zwróci książki danego autora.

---

## 8. System zamówień
**Zadanie:**  
Stwórz strukturę `Order` z polami `ID`, `CustomerName` i `TotalPrice`.

**Rozszerzenie:**  
Stwórz listę zamówień i napisz funkcję, która znajdzie zamówienie o największej wartości.

---

## 9. Struktura Zwierzęcia
**Zadanie:**  
Stwórz strukturę `Animal` z polami `Name`, `Species` i `Sound`.  
Dodaj metodę `MakeSound()`, która wypisze dźwięk zwierzęcia.

**Rozszerzenie:**  
Stwórz slice zwierząt i napisz funkcję, która wyświetli dźwięki wszystkich zwierząt w pętli.

---

## 10. Rejestracja użytkownika
**Zadanie:**  
Stwórz strukturę `User` z polami `Username`, `Email` i `Password`.  
Dodaj metodę `Validate()`, która sprawdzi, czy email zawiera `@`, a hasło ma przynajmniej 8 znaków.

**Rozszerzenie:**  
Dodaj metodę `HashPassword()`, która zamieni hasło na ukrytą wersję (np. `******`).
