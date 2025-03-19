# Zadania do nauki `context` w Go

## 1. Tworzenie podstawowego kontekstu
**Zadanie**: Stwórz kontekst przy użyciu `context.Background()`, a następnie wydrukuj jego wartość.  
**Cel**: Zrozumienie podstawowego użycia kontekstu.

## 2. Przekazywanie kontekstu do funkcji
**Zadanie**: Stwórz funkcję, która przyjmuje `context.Context` jako argument. Funkcja powinna drukować komunikat, gdy jest wywołana. Wywołaj ją, przekazując kontekst z `context.Background()`.  
**Cel**: Zrozumienie, jak przekazywać kontekst między funkcjami.

## 3. Anulowanie operacji za pomocą kontekstu
**Zadanie**: Utwórz kontekst z funkcją `context.WithCancel()`. Następnie wywołaj operację, która po kilku sekundach zostanie anulowana przez kontekst.  
**Cel**: Nauka anulowania operacji.

## 4. Limit czasu operacji
**Zadanie**: Stwórz kontekst z limitem czasu (np. 3 sekundy) przy pomocy `context.WithTimeout()`. Symuluj długą operację, która powinna zostać anulowana po upływie tego czasu.  
**Cel**: Poznanie kontekstu z limitem czasu.

## 5. Przekazywanie wartości w kontekście
**Zadanie**: Utwórz kontekst z danymi (np. nazwą użytkownika) przy użyciu `context.WithValue()`. Przekaż ten kontekst do funkcji, która wypisze przechowywaną w nim wartość.  
**Cel**: Zrozumienie, jak przechowywać i przekazywać dane w kontekście.

## 6. Sprawdzanie stanu kontekstu w operacji
**Zadanie**: Stwórz operację, która sprawdza, czy kontekst został anulowany (użyj `ctx.Done()`). Symuluj operację, która po pewnym czasie się zatrzyma.  
**Cel**: Nauka monitorowania stanu kontekstu (np. czy został anulowany).

## 7. Zagnieżdżanie kontekstów
**Zadanie**: Utwórz główny kontekst z `context.Background()`, a następnie stwórz kontekst podrzędny z `context.WithCancel()` i przekaż go do funkcji. Sprawdź, jak anulowanie kontekstu podrzędnego wpływa na główny kontekst.  
**Cel**: Zrozumienie zagnieżdżania kontekstów i ich wzajemnych zależności.

## 8. Kontekst z wieloma operacjami
**Zadanie**: Utwórz dwa konteksty z limitem czasu i przeprowadź dwie operacje, które działają równolegle. Jedna z operacji powinna zakończyć się przed upływem czasu, a druga po.  
**Cel**: Zrozumienie, jak kontekst wpływa na równoległe operacje.

## 9. Timeout w operacjach sieciowych
**Zadanie**: Utwórz funkcję symulującą operację sieciową (np. HTTP request), która akceptuje kontekst z limitem czasu. W przypadku przekroczenia limitu czasu, operacja powinna zostać anulowana.  
**Cel**: Nauka używania kontekstu w prawdziwych operacjach sieciowych.

## 10. Połączenie kontekstu z błędami
**Zadanie**: Zaimplementuj operację, która zwraca błąd, jeśli kontekst zostanie anulowany (użyj `ctx.Err()`). Przykład: symuluj zapytanie do bazy danych, które zwróci błąd, jeśli operacja przekroczy czas oczekiwania.  
**Cel**: Zrozumienie, jak kontekst może przekazywać informacje o błędach w przypadku anulowania.
