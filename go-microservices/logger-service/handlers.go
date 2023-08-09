package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func (app *AppConfig) listenRPC() {
	listener, err := net.Listen("tcp", app.ADDR_RPC)
	if err != nil {
		log.Fatalln("RPC Server Error:", err)
	}
	defer listener.Close()

	log.Printf("RPC Server Listener Address: %+v\n", listener.Addr())

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Printf("RPC Server Connection Error: %+v\n", err)
			continue
		}

		log.Printf("RPC Server Connection: %+v\n", connection)

		go rpc.ServeConn(connection)
	}
}

func (app *AppConfig) handleSetup() error {
	rpcErr := rpc.Register(app)
	if rpcErr != nil {
		return rpcErr
	}

	go app.listenRPC()

	return nil
}

func (app *AppConfig) handleWriteLog(resw http.ResponseWriter, req *http.Request) {
	requestPayload := RequestPayload{}

	err := readJSON(resw, req, &requestPayload)
	if err != nil {
		errorJSON(resw, err, http.StatusBadRequest)
		return
	}

	entry, err := app.Services.write(&requestPayload)
	if err != nil {
		errorJSON(resw, errors.New("failed to log"), http.StatusInternalServerError)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Logged entry: %s", entry.InsertedID),
		Data:    entry,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) handleReadLogs(resw http.ResponseWriter, req *http.Request) {
	entries, err := app.Services.all()
	if err != nil {
		errorJSON(resw, errors.New("failed to access logs"), http.StatusInternalServerError)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Logged entries: %d", len(entries)),
		Data:    entries,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) HandleLogRPC(req LogRPCPayload, res *ResponsePayload) error {
	requestPayload := RequestPayload{
		Name: req.Name,
		Data: req.Data,
	}
	entry, err := app.Services.write(&requestPayload)
	if err != nil {
		log.Println("failed to log", err)
		return err
	}

	data := fmt.Sprintf("%v", entry)
	*res = ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Logged entry: %s", entry.InsertedID),
		Data:    data,
	}

	return nil
}
