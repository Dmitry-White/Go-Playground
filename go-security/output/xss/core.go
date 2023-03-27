package main

import (
	"errors"
	"fmt"
	"go-security/output/xss/dal"
	"log"
	"net/http"
)

type HandleFunc func(resw http.ResponseWriter, ms []dal.Message, user string) error

const (
	SAFE_STRATEGY  = "safe"
	HACKY_STRATEGY = "hacky"
)

func getRenderFunc(strategy string) (HandleFunc, error) {
	choice := fmt.Sprintf("Using %s render", strategy)
	log.Println(choice)

	switch strategy {
	case SAFE_STRATEGY:
		return renderSafe, nil
	case HACKY_STRATEGY:
		return renderHacky, nil
	}
	return nil, errors.New("this strategy is not supported")
}
