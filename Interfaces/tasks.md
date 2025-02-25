# Zadania z interfejsami w Golang

## 1. Prosty interfejs
Napisz interfejs `Speaker` z metodą `Speak() string`. Zaimplementuj go dla struktur `Dog` i `Cat`.

## 2. Interfejs dla figur geometrycznych
Stwórz interfejs `Shape` z metodą `Area() float64`. Zaimplementuj go dla `Circle` i `Rectangle`.

## 3. Interfejs do serializacji
Zdefiniuj interfejs `Serializable` z metodą `ToJSON() string`. Zaimplementuj go dla `User` i `Product`.

## 4. Interfejs z wieloma metodami
Napisz interfejs `Vehicle` z metodami `Start()`, `Stop()`. Zaimplementuj go dla `Car` i `Bike`.

## 5. Interfejs zwracający różne typy
Stwórz interfejs `Generator`, który zwraca różne typy wartości (int, string itp.). Zaimplementuj go dla `RandomInt` i `RandomString`.

## 6. Interfejs i funkcja przyjmująca go jako argument
Napisz funkcję `MakeNoise(Speaker)`, która wywołuje `Speak()`. Przetestuj z `Dog` i `Cat`.

## 7. Interfejs do operacji matematycznych
Zdefiniuj interfejs `Calculator` z metodą `Compute(a, b int) int`. Zaimplementuj go dla `Adder` i `Multiplier`.

## 8. Interfejs do odczytu plików
Stwórz interfejs `Reader` z metodą `Read() string`. Zaimplementuj go dla `FileReader` i `NetworkReader`.

## 9. Interfejs do logowania
Zdefiniuj interfejs `Logger` z metodą `Log(message string)`. Zaimplementuj `ConsoleLogger` i `FileLogger`.

## 10. Interfejs dla systemu płatności
Napisz interfejs `Payment` z metodą `Pay(amount float64) string`. Zaimplementuj `CreditCard` i `PayPal`.

## 11. Interfejs dla zadań asynchronicznych
Zdefiniuj interfejs `Task` z metodą `Execute()`. Zaimplementuj `EmailTask` i `BackupTask`.

## 12. Interfejs do transformacji tekstu
Napisz interfejs `Transformer` z metodą `Transform(string) string`. Zaimplementuj `UpperCase` i `Reverse`.

## 13. Interfejs dla systemu powiadomień
Stwórz interfejs `Notifier` z metodą `SendNotification(message string)`. Zaimplementuj `EmailNotifier` i `SMSNotifier`.

## 14. Interfejs do operacji na bazie danych
Zdefiniuj interfejs `Database` z metodami `Insert()`, `Delete()`. Zaimplementuj `MySQL` i `MongoDB`.

## 15. Interfejs dla gier komputerowych
Napisz interfejs `Character` z metodą `Attack() string`. Zaimplementuj `Warrior` i `Mage`.

## 16. Interfejs dla systemu obsługi błędów
Stwórz interfejs `ErrorReporter` z metodą `ReportError(err error)`. Zaimplementuj `ConsoleReporter` i `FileReporter`.

## 17. Interfejs dla odtwarzaczy multimedialnych
Zdefiniuj interfejs `MediaPlayer` z metodami `Play()`, `Pause()`. Zaimplementuj `MP3Player` i `VideoPlayer`.

## 18. Interfejs dla systemu transportowego
Napisz interfejs `Transport` z metodą `Move()`. Zaimplementuj `Car`, `Train`, `Plane`.

## 19. Pusty interfejs (`interface{}`)
Napisz funkcję `PrintAnything(value interface{})`, która przyjmuje dowolną wartość i ją wyświetla.

## 20. Interfejs i dynamiczne rzutowanie (`type assertion`)
Stwórz interfejs `Animal` z metodą `MakeSound()`. Napisz funkcję, która sprawdza, czy dany obiekt jest `Dog`, i jeśli tak, wywołuje metodę `Bark()`.
