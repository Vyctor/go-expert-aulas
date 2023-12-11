package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled!. Timeout Reached")
		return nil
	case <-time.After(2 * time.Second):
		fmt.Println("Hotel Booked!")
		return nil
	}
}
