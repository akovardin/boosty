package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

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

	fmt.Printf("current: %+v\n\n", s)

	v := url.Values{}
	v.Add("offset", "0")
	v.Add("limit", "100")

	ss, err := b.Subscribers(v)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range ss.Data {
		fmt.Printf("subscriber: %+v\n\n", s)
	}
}
