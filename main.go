package main

import (
	"fmt"
	"google.golang.org/grpc/metadata"
)

func main() {
	var _ metadata.MD
	fmt.Println("Yo v2")
}
