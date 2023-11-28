# boosty

Библиотека для работы с приватным API boosty

## Использование

Установка не совсем стандартная. Нужно использовать отдельный домен для go get

```shell
go get kovardin.ru/projects/boosty
```

Пакет будет устанавливать из оригинального репозитория https://gitflic.ru/project/getapp/boosty

Для инициализации необходимо указать блог и токен. Токен можно забрать из браузера

```golang
auth, err := auth.New(
	auth.WithFile(".boosty"),
    // auth.WithInfo(auth.Info{}),
    auth.WithInfoUpdateCallback(func (i auth.Info) {
        log.Printf("info update: %+v\n", i)
    }),
)
if err != nil {
    log.Fatal(err)
}

request, err := request.New(
    //request.WithUrl("https://api.boosty.to"),
    request.WithClient(&http.Client{}),
    request.WithAuth(auth),
)
if err != nil {
    log.Fatal(err)
}

b, err := boosty.New("getapp", boosty.WithRequest(request))
if err != nil {
    log.Fatal(err)
}
```
Канал с новостями [@kodikapusta](https://t.me/kodikapusta)
