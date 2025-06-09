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


**Zadanie 5 React**

✅ 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz
Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w
Produktach powinniśmy pobierać dane o produktach z aplikacji
serwerowej; [Link do kodu]([https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad4/routes/routes.go](https://github.com/michalharasim/e-biznes-uj-class/tree/main/zad5))

❌ 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing

❌ 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za
pomocą React hooks

❌ 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz
kliencką na dockerze via docker-compose

❌ 5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS

https://github.com/user-attachments/assets/e6d32505-e2fd-4dea-a212-a1568127aac9

**Zadanie 6 Testy**

✅ 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium
(Kotlin, Python, Java, JS, Go, Scala) [Link do kodu](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad6/frontend/src/tests/test_all.py)

❌ 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50
asercji

❌ 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego
projektu z minimum 50 asercjami

❌ 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z
minimum jednym scenariuszem negatywnym per endpoint

❌ 5.0 Należy uruchomić testy funkcjonalne na Browserstacku

https://github.com/user-attachments/assets/bd63329d-c742-446a-b961-42eb427de91b

**Zadanie 7 Sonar**

Uwaga: Zadanie zostało wykonane dla przedmiotu E-biznes: [LINK](https://github.com/michalharasim/e-biznes-uj-class/tree/main/zad6)

✅ 3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w hookach gita. [Link do commita](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad6/pre-commit)

✅ 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod aplikacji serwerowej) Brak zmian

✅ 4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod aplikacji serwerowej) Brak zmian

✅ 4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa w kodzie w Sonarze (kod aplikacji serwerowej) Brak zmian

✅ 5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie aplikacji klienckiej [Link do commita](https://github.com/michalharasim/e-biznes-uj-class/commit/3fc9e3e1f5fd106d62654244e92e00c3263f5801)
[Link do commita](https://github.com/michalharasim/e-biznes-uj-class/commit/a313b1f254a1ca528ac7971fee1324923d2f8766)
[Link do commita](https://github.com/michalharasim/e-biznes-uj-class/commit/70f35c206cd8d6cb1a21cbc38a82aabbec534e3c)

BACKEND

[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=ebiznes_backend)](https://sonarcloud.io/summary/new_code?id=ebiznes_backend)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_backend&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ebiznes_backend)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_backend&metric=bugs)](https://sonarcloud.io/summary/new_code?id=ebiznes_backend)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_backend&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=ebiznes_backend)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_backend&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=ebiznes_backend)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_backend&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=ebiznes_backend)

FRONTEND

[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=ebiznes_frontend)](https://sonarcloud.io/summary/new_code?id=ebiznes_frontend)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_frontend&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ebiznes_frontend)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_frontend&metric=bugs)](https://sonarcloud.io/summary/new_code?id=ebiznes_frontend)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_frontend&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=ebiznes_frontend)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_frontend&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=ebiznes_frontend)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=ebiznes_frontend&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=ebiznes_frontend)

**Zadanie 8 Auth**

✅ 3.0 logowanie przez aplikację serwerową (bez Oauth2) [Link do backendu](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad8/backend/server.go)

✅ 3.5 rejestracja przez aplikację serwerową (bez Oauth2) [Link do frontendu](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad8/frontend/src/components/AuthForm.jsx)

❌ 4.0 logowanie via Google OAuth2

❌ 4.5 logowanie via Facebook lub Github OAuth2

❌ 5.0 zapisywanie danych logowania OAuth2 po stronie serwera

https://github.com/user-attachments/assets/d0d7fb92-efe8-4a8b-ab8a-9401211be2b3


**Zadanie 9 ChatGPT bot**

✅ 3.0 należy stworzyć po stronie serwerowej osobny serwis do łącznia z chatGPT do usługi chat [Link do commita](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad9/chat.py)

❌ 3.5 należy stworzyć interfejs frontowy dla użytkownika, który komunikuje się z serwisem; odpowiedzi powinny być wysyałen do frontendowego interfejsu

❌ 4.0 stworzyć listę 5 różnych otwarć oraz zamknięć rozmowy

❌ 4.5 filtrowanie po zagadnieniach związanych ze sklepem (np. ograniczenie się jedynie do ubrań oraz samego sklepu) do GPT

❌ 5.0 filtrowanie odpowiedzi po sentymencie

https://github.com/user-attachments/assets/0f394fbe-e7f5-41f1-8385-42b53dc799cd

**Zadanie 10 Chmura/CI**

Zadanie zrobione na podstawie serwisów z zadania 5.

Adres do frontendu: https://frontappnew.azurewebsites.net/

Adres do backendu: https://zad10backendapp-e4eebcg9d4drc3hg.westeurope-01.azurewebsites.net/products

✅ 3.0 Należy stworzyć odpowiednie instancje po stronie chmury na
dockerze [Link do Dockerfile 1](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad10/frontend/Dockerfile)

[Link do Dockerfile 2](https://github.com/michalharasim/e-biznes-uj-class/blob/main/zad10/backend/Dockerfile)

❌ 3.5 Stworzyć odpowiedni pipeline w Github Actions do budowania aplikacji (np. via fatjar)

❌ 4.0 Dodać notyfikację mailową o zbudowaniu aplikacji

❌ 4.5 Dodać krok z deploymentem aplikacji serwerowej oraz klienckiej na chmurę

❌ 5.0 Dodać uruchomienie regresyjnych testów automatycznych (funkcjonalnych) jako krok w Actions

https://github.com/user-attachments/assets/402c331a-ca6b-43d0-9d5d-e82b25f978fd
