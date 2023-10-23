# boosty

Библиотека для работы с приватным API boosty

## Использование

Установка не совсем стандартная. Нужно использовать отдельный домен для go get

```shell
go get kovardin.ru/projects/boosty
```

Пакет будет устанавливать из оригинального репозитория https://gitflic.ru/project/getapp/boosty

Для инициализации необходимо указать блог и токен. Токен можно забрать из браузера 

```go
b := boosty.New("getapp", "5d4b7d8701xxxxxxxxxxxxx")
```
