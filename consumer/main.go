package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {

	//se define rabbitMQ server URL
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	//crear una conexion rabbitMQ
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	//abrir un channel en la instancia rabbitMQ con la que se ha establecido conexion
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	//suscribirse a una cola para recibir mensajes
	messages, err := channelRabbitMQ.Consume(
		"QueueService1",
		"",    //consumidor
		true,  //auto-ack
		false, //exclusive
		false, // no local
		false, // sin esperas
		nil,   // argumentos, nil por ahora
	)
	if err != nil {
		log.Println(err)
	}

	//mensaje de bienvenida del suscriptor (consumidor)
	log.Println("Conexion con RabbitMQ exitosa")
	log.Println("Listo para recibir mensajes...")

	//crear un go channel para recibir mensajes en un loop infinito
	alwaysListen := make(chan bool)

	go func() {
		for message := range messages {
			//mostrar los mensajes recibidos en la consola
			log.Printf("--> Recibido: %s\n", message.Body)
		}
	}()

	<-alwaysListen

}
