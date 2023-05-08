package main

import (
	"context"
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

const (
	TIMEOUT = 10 * time.Millisecond
)

type Bid struct {
	url   string
	price float32
}

var DEFAULT_BID = Bid{
	url:   "http://dmitrywhite.com",
	price: 0.1,
}

func getBidOutcome(url string) bool {
	h := fnv.New32()
	h.Write([]byte(url))

	seed := int64(h.Sum32())
	source := rand.NewSource(seed)
	random := rand.New(source)

	return random.Int()%2 == 0
}

func bestBid(ch chan Bid, url string) {
	outcome := getBidOutcome(url)

	if outcome {
		time.Sleep(2 * time.Millisecond)

		bid := Bid{
			url:   url,
			price: 0.05,
		}
		ch <- bid
	} else {
		time.Sleep(TIMEOUT + 20*time.Millisecond)

		bid := Bid{
			url:   url,
			price: 0.018,
		}
		ch <- bid
	}
}

func bidOn(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1)

	go bestBid(ch, url)

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		log.Printf("Bid for %q timed out, returning default", url)
		return DEFAULT_BID
	}
}

func realTimeBidding() interface{} {
	bids := []Bid{}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	successfulBid := bidOn(ctx, "http://google.com")
	bids = append(bids, successfulBid)

	ctx, cancel = context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	timeoutBid := bidOn(ctx, "http://example.com")
	bids = append(bids, timeoutBid)

	return bids
}
