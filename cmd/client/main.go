package main

import (
	"awsproj/internal/hashes/hashesDelivery"
	"awsproj/internal/strings/stringsRepository"
	"awsproj/internal/strings/stringsUseCase"
	"bufio"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

var (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := hashesDelivery.NewHasherClient(conn)

	stringsR := stringsRepository.NewStringsRepository(c)
	stringsUC := stringsUseCase.NewStringsUseCase(stringsR)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter string: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)

		fmt.Println(stringsUC.HashString(text))
	}
}