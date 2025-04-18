**Zadanie 1 Docker**

✅ 3.0 obraz ubuntu z Pythonem w wersji 3.10 [Link do obrazu](https://hub.docker.com/repository/docker/michalharasim/zad1/tags/latest/sha256-9547de1a949110c025e0ed0a745632f699f3e4ae064b0be8885b009031903b16)

✅ 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem [Link do obrazu](https://hub.docker.com/repository/docker/michalharasim/zad1/tags/latest/sha256-9547de1a949110c025e0ed0a745632f699f3e4ae064b0be8885b009031903b16)

❌ 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC
SQLite w ramach projektu na Gradle (build.gradle)

❌ 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji
przez CMD oraz gradle

❌ 5.0 dodać konfigurację docker-compose



**Zadanie 2 Należy stworzyć aplikację na frameworku Play w Scali 3**

✅ 3.0 Należy stworzyć kontroler do Produktów [Link do kodu](https://github.com/michalharasim/e-biznes-uj-class/tree/main/zad2)

✅ 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane
pobierane z listy [Link do kodu](https://github.com/michalharasim/e-biznes-uj-class/tree/main/zad2)

❌ 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
zgodnie z CRUD

❌ 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
skrypt uruchamiający aplikację via ngrok

❌ 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD


Nagranie: https://github.com/user-attachments/assets/050ef2ac-15a3-4d29-bf37-e44a3a5fc051




**Zadanie 3 Kotlin**

✅ 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor,
która pozwala na przesyłanie wiadomości na platformę Discord [Link do kodu](https://github.com/michalharasim/e-biznes-uj-class/tree/main/zad3/src/main/kotlin)

❌ 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z
platformy Discord skierowane do aplikacji (bota)

❌ 4.0 Zwróci listę kategorii na określone żądanie użytkownika

❌ 4.5 Zwróci listę produktów wg żądanej kategorii

❌ 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger,
Webex

Nagranie: https://github.com/user-attachments/assets/0c60b656-9491-4517-b8c3-86bc87065875

**Zadanie 4 GO**

✅ 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD [Link do kodu](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad4/routes/routes.go)

✅ 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy) [Link do kodu](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad4/models/product.go)

❌ 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint

❌ 4.5 Należy stworzyć model kategorii i dodać relację między kategorią, a produktem

❌ 5.0 pogrupować zapytania w gorm’owe scope'y


https://github.com/user-attachments/assets/5adf6d0f-d98a-4ec7-a123-83a8c3a6119b



