package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

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

func (app *AppConfig) listenRPC() error {
	log.Printf("RPC Server listening on %d", app.PORT_RPC)

	listener, err := net.Listen("tpc", app.ADDR_RPC)
	if err != nil {
		return nil
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(connection)
	}
}

func (app *AppConfig) handleSetup() error {
	rpcErr := rpc.Register(&app)
	if rpcErr != nil {
		log.Panic(rpcErr)
	}

	go app.listenRPC()

	return nil
}

func (app *AppConfig) handleLogRPC(req RequestPayloadRPC, res *string) error {
	entry, err := app.Services.writeRPC(&req)
	if err != nil {
		log.Println("failed to log", err)
		return err
	}

	data := fmt.Sprintf("%v", entry)
	*res = "Processed log via RPC" + data

	return nil
}
