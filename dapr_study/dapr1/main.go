package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()
	data := []byte("hello")
	store := "statestore" // defined in the component YAML

	// save state with the key key1, default options: strong, last-write
	if err := client.SaveState(ctx, store, "key1", data, nil); err != nil {
		panic(err)
	}

	// get state for key key1
	item, err := client.GetState(ctx, store, "key1", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data [key:%s etag:%s]: %s", item.Key, item.Etag, string(item.Value))

	// delete state for key key1
	if err := client.DeleteState(ctx, store, "key1", nil); err != nil {
		panic(err)
	}
}
