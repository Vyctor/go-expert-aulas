package main

import "context"

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "abc123")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) error {
	token := ctx.Value("token").(string)
	println(token)
	return nil
}
