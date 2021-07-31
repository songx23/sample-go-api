package main

import "info.songxue/go/sample/api/internal/rest"

func main() {
	s := rest.New(8080)

	s.Run()
}
