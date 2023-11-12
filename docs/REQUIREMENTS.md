### Тестовое Задание для Бэкенд-Разработчика

#### Задача:
Разработать систему, состоящую из воркера, gRPC сервиса и API.

#### Компоненты:

1. **Воркер:**
   - Читает данные из JSON файла (data.json).
   - Отправляет данные через gRPC в сервис.

2. **gRPC Сервис:**
   - Принимает и сохраняет данные в Elasticsearch.
   - Возвращает данные в виде бесконечного списка (infinite scroll).
   - Поддерживает полнотекстовый поиск по полю title учитывая русскую и руммынсую морфологию.
   - Возвращает агрегированные данные по полю subcategory ( количество документов по каждой подгатегории ).

3. **API:**
   - GraphQL (в крайнем случае REST) для предоставления данных из gRPC сервиса.

#### Требования:

1. **Язык программирования:** Golang или Python.
   - Если используется Python, то обязательно использовать асинхронные библиотеки, например, AsyncIO, FastAPI или aiohttp.

2. **База данных:** Elasticsearch.
3. **Контейнеризация:** Docker (обязательно).

#### Функциональные Требования:

1. **Воркер:**
   - Должен эффективно читать данные из больших JSON файлов, даже на машинах с ограниченными ресурсами.
		- размер файла может быть более 1ТБ а у вас на машине 1ГБ памяти

2. **gRPC Сервис:**
   - Должен корректно принимать и сохранять данные в Elasticsearch.
   - Должен предоставлять возможность просмотра данных в виде бесконечного списка.
   - Должен предоставлять возможность просмотра количества данных по каждой подкатегории.

3. **API:**
   - Должен корректно отображать данные, полученные от gRPC сервиса.

#### Нефункциональные Требования:

1. **Код:**
   - Код должен быть чистым, модульным и хорошо организованным.
   - Должны быть использованы подходящие шаблоны проектирования.
   - Код должен соответствовать принципам SOLID (например, Clean Architecture или DDD).

2. **Тестирование:**
   - Проект должен содержать юнит-тесты.
   - Используйте mock-объекты, если это необходимо.

3. **Документация:**
   - Каждый компонент системы должен быть задокументирован.

4. **Оркестрация:**
   - Возможность запуска всей системы с помощью команды `docker compose up` будет считаться преимуществом.

#### Ограничение по Времени:
- Нам важно увидеть, как вы подходите к решению задач, и мы учитываем, что у вас может не быть возможности полностью завершить проект. 
Частичное выполнение также даст нам понять ваш уровень и подход к решению задач.

### Сроки и Доставка:

1. **Срок:** 10 дней с момента получения задания.
2. **Доставка:** 
   - Ссылка на репозиторий на GitHub с исходным кодом проекта.
   - Инструкция по запуску проекта, включая Docker.

### Оценка:

Проект будет оцениваться по следующим критериям:

- Как вы подходите к решению задачи и разбиению её на подзадачи.
- Соблюдение функциональных и нефункциональных требований.
- Качество кода и его организация.
- Наличие и качество тестов.
- Наличие и качество документации.
- Наличие оркестрации с использованием Docker Compose.

Удачи! Если у вас есть какие-либо вопросы по заданию, не стесняйтесь спрашивать.