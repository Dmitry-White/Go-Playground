package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

// Message is a message from user
type Message struct {
	Content string

	once      sync.Once
	signature string // cached signature
}

// Signature returns the digital signature of the message
func (m *Message) Sign() string {
	m.once.Do(m.calcSig)
	return m.signature
}

func (m *Message) calcSig() {
	log.Printf("Calculating new signature...")

	h := sha1.New()
	io.Copy(h, strings.NewReader(m.Content))
	signature := h.Sum(nil)

	m.signature = fmt.Sprintf("%x", signature)
}
