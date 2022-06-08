#### Selekcja Naturalna API

## Permissions:
```
0000000000000001 READ_SELF        Pozwala na pozyskanie danych o sobie.
0000000000000010 READ_USERS       Odczytywanie wszystkich urzytkowników.
0000000000000100 MODIFY_USERS     Modyfikacja danych użytkownika (w tym płatności).
0000000000001000 MANAGE_USERS     Możliwość dodawania i usuwania użytkowników oraz modyfikacji ich typu.
0000000000010000 READ_RECEIPTS    Umożliwia odczytywanie wszystkich paragonów.
0000000000100000 MODIFY_RECEIPTS  Umożliwia dodanie oraz modyfikowanie swoich paragonów.
0000000001000000 MANAGE_RECEIPTS  Umożliwia modyfikowanie wszystich paragonów lub ich usuwanie.
0000000010000000 READ_DOCS        Pozwala na odczytywanie wszystich uchwał i dokumentów.
0000000100000000 MODIFY_DOCS      Daje możliwość dodania dokumentu.
0000001000000000 MANAGE_DOCS      Daje możliwość usuwania i modyfikacji parametrów dokumentów.
0000010000000000 MESSAGE_DISCORD  |
0000100000000000 MESSAGE_EMAIL    +- Pozwala wysłąć wiadmomość do użytkownika na za pomocą danego komunikatora
0001000000000000 MESSAGE_FACEBOOK |
0010000000000000 MESSAGE_PHONE    Umożliwia odczytanie danych o telefonie.
0100000000000000 ADMIN            Umożliwia wszystko powyższe. Możliwość przyznawania uprawnień czasowych użytkownikom.
```

## User types:
```
0 UNDEFINED Brak dostępu do wszystkich elementów. Wysyła zgłoszenie do administracji.
1 VIEWER    Pozwala przeglądać dane zgodnie z permisjami z perfiksem READ_ za wyjątkiem READ_SELF
2 DEFAULT   Użytkownik wyświetlany w bazie. Posiada przypisane płatności.
3 DELETED   Uprawnienia VIEWER, tyle że widzi siebie i jest wliczany do wpłaconych pieniędzy, ale to czego nie zapłacił nie jest brakiem
```