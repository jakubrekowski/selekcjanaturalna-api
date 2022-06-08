#### Selekcja Nautralna API

## Endpoints:
```http
prefx: /api/v0/

POST   /auth/login/discord  Endpoint logowania do systemu za pomocą Discord
[//]: # (GET  /users/@me Zwraca obiekt altualnie zalogowanego użytkownika)
GET    /users               Zwraca obiekty wszystkich użytkowników
GET    /user/:id            Zwraca obiekt zadanego użytkownika
POST   /user                Tworzy obiekt użytkownika
PUT    /user/:id            Uaktualnia obiekt użytkownika
DELETE /user/:id            *Usuwa obiekt użytkownika
GET    /receipts            Zwraca wszystkie obiekty paragonów
GET    /receipt/:id         Zwraca obiekt zadaniego paragonu
POST   /receipt             Tworzy obiekt paragonu
PUT    /receipt/:id         Uaktualnia obiekt paragonu
DELETE /receipt/:id         Usuwa obiekt paragonu
GET    /documents           Zwraca obiekty wszystkich dokumentów
GET    /document/:id        Zwraca obiekt zadanego dokumentu
POST   /document            Tworzy obiekt dokumentu
PUT    /document/:id        Akutalizuje obiekt dokumentu
DELETE /document/:id        Usuwa obiekt dokumentu

```