package main

import (
	"context"
	"log"
	"time"

	pb "github.com/XavierCuases/gRPC-Go/greet" // Ruta del paquete generado
	"google.golang.org/grpc"
)

func main() {
	// Conectar al servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// Enviar una solicitud al servidor
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Greet(ctx, &pb.GreetRequest{Name: "Mundo"})
	if err != nil {
		log.Fatalf("Error al llamar al método Greet: %v", err)
	}

	log.Printf("Respuesta del servidor: %s", res.GetMessage())
}
