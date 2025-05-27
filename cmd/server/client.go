package main

import (
	tt "fullGen/tintin"
)

func CreateClient() *tt.Client {
	client, err := tt.NewClient("http://localhost:8080")
	if err != nil {
		panic(err.Error())
	}
	return client
}
