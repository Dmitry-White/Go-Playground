package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type HandleFunc func(decoder *json.Decoder) (*PaymentBasic, error)

const (
	SAFE_STRATEGY  = "safe"
	BASIC_STRATEGY = "basic"
	HACKY_STRATEGY = "hacky"
)

func getHandleFunc(strategy string) (HandleFunc, error) {
	choice := fmt.Sprintf("Using %s handler", strategy)
	log.Println(choice)

	switch strategy {
	case SAFE_STRATEGY:
		return handleSafe, nil
	case BASIC_STRATEGY:
		return handleBasic, nil
	case HACKY_STRATEGY:
		return handleHacky, nil
	}
	return nil, errors.New("this strategy is not supported")
}
