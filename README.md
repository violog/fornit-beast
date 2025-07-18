# Beast+Docker

Руководство по запуску Beast кроссплатформенно с помощью Docker с возможностью отладки

Docker — хороший инструмент для работы с приложениями вне зависимости от ОС разработчика, т.е. кроссплатформенной разработки. Таким образом, можно избежать многочисленных проблем переноса локального проекта между машинами. Подробнее можно почитать в интернете, здесь только небольшая инструкция.

## Условия для работы
- Golang
- Docker (с плагином Compose)
- Консоль для ввода команд

В этом репозитории взять файлы `Dockerfile` и `docker-compose.yml`. Поместить их в корень проекта, как в репозитории, т.е. в одно место с папками BEAST_GO, BEAST_PHP. Все описанные команды необходимо запускать отсюда, это важно для Docker.

Есть 2 способа запуска.

## Без отладки

Первый раз достаточно запустить `docker compose up -d`. Это скачает Docker-образ сервера PHP для пульта, локально соберёт образ бота из BEAST_GO и запустит оба приложения. Файлы сохранений в BEAST_PHP будут изменяться и локально. Бот готов к работе.

Остановка: `docker compose down -v`
Остановка с флагом `-v` удаляет контейнеры и разделы, но не удаляет файлы памяти бота.

Когда код Beast изменяется, нужно останавливать проект командой down выше, а запускать в дальнейшем с дополнительным флагом для повторной сборки Beast:
`docker compose up -d --build`

## С отладкой в IDE

После каждого запуска нужно останавливать контейнер Beast, чтобы не было конфликта портов:
`docker compose stop go`. Альтернатива — закомментировать всю секцию `go-server` в `docker-compose.yml`. Оба варианта предполагают запуск PHP в контейнере, а Beast локально.