package main

import (
	"fmt"
	"log"
	"net/http"

	"kovardin.ru/projects/boosty"
	"kovardin.ru/projects/boosty/auth"
	"kovardin.ru/projects/boosty/request"
)

func main() {
	auth, err := auth.New(
		auth.WithFile(".boosty"),
		//auth.WithInfo(auth.Info{}),
		auth.WithInfoUpdateCallback(func(i auth.Info) {
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

	s, err := b.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("current: %+v", s)
}
