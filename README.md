#   YADRO ИМПУЛЬС тестовое задание "club_manager"


## Описание

В репозитории представлены все исходные файлы и файлы для тестирования программы( 6 тестов ).

Программа написана на языке Golang стандарта 1.22 используя только стандартную библиотеку
(https://pkg.go.dev/std).

## Инструкции по запуску

```bash
git clone https://github.com/your_repository.git
cd your_repository/src
go run ./ClubInit.go <путь к тестовому файлу>

```
## Инструкции по запуску в контейнере
Нужно перейти в корневую папку проекта и далее ввести команды 
```bash
 docker build -t club_manager .  
 docker run -it --name <имя_контейнера> club_manager
 $ ./main data/test<1-6>.txt

```
