package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	//declarar una cola que puede ser para suscribir(consumir) o publicar (producir)
	//campos minimos
	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1", //nombre de la cola
		true,            // durable
		false,           //auto delete
		false,           // exclusiva
		false,           // sin tiempo de espera
		nil,             //argumentos, sin args por ahora
	)
	if err != nil {
		log.Println("Error en creacion de queue", err)
		panic(err)

	}

	//web server, instancia fiber
	app := fiber.New()

	//agregar middleware
	app.Use(
		logger.New(), //logger simple
	)

	//agregar routes
	app.Get("send/", func(c *fiber.Ctx) error {
		//crear un mensaje para publicar
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("msg")),
		}

		//publicar el mensaje en la cola
		if err := channelRabbitMQ.Publish(
			"", //exchange
			"QueueService1",
			false,   //obligatorio enrutar el mensaje al menos a una queue
			false,   // enviar el mensaje a un consumidor de forma inmediata
			message, //mensaje a publicar
		); err != nil {
			log.Fatal(err)
			return err
		}
		return nil

	})
	//iniciar el servidor fiber
	log.Fatal(app.Listen(":3000"))

}
