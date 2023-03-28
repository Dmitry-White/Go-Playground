package main

import (
	"errors"
	"fmt"
	"go-security/output/sensitive-data/dal"
	"log"
	"net/http"
)

type HandleFunc func(resw http.ResponseWriter, user *dal.User)

const (
	SAFE_STRATEGY  = "safe"
	HACKY_STRATEGY = "hacky"
)

func getDataHandler(strategy string) (HandleFunc, error) {
	choice := fmt.Sprintf("Using %s data handler", strategy)
	log.Println(choice)

	switch strategy {
	case SAFE_STRATEGY:
		return handleDataSafe, nil
	case HACKY_STRATEGY:
		return handleDataHacky, nil
	}
	return nil, errors.New("this strategy is not supported")
}
