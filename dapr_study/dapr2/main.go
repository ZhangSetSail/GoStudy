package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"time"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	ctx := context.Background()
	store := "statestore"
	item1 := &dapr.SetStateItem{
		Key: "key1",
		Etag: &dapr.ETag{
			Value: "1",
		},
		Metadata: map[string]string{
			"created-on": time.Now().UTC().String(),
		},
		Value: []byte("hello"),
		Options: &dapr.StateOptions{
			Concurrency: dapr.StateConcurrencyLastWrite,
			Consistency: dapr.StateConsistencyStrong,
		},
	}

	item2 := &dapr.SetStateItem{
		Key: "key2",
		Metadata: map[string]string{
			"created-on": time.Now().UTC().String(),
		},
		Value: []byte("hello again"),
	}

	item3 := &dapr.SetStateItem{
		Key: "key3",
		Etag: &dapr.ETag{
			Value: "1",
		},
		Value: []byte("hello again"),
	}

	if err := client.SaveBulkState(ctx, store, item1, item2, item3); err != nil {
		panic(err)
	}

	keys := []string{"key1", "key2", "key3"}
	items, err := client.GetBulkState(ctx, store, keys, nil, 100)
	for {
		time.Sleep(10 * time.Second)
		for _, item := range items {
			fmt.Println(item.Etag)
			fmt.Println(item.Error)
			fmt.Println(item.Metadata)
			fmt.Println(string(item.Value))
		}
	}
}
