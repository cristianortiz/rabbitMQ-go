
# Ejemplo practico de la implementación de rabbitMQ en golang

implmenta un Message broker con rabbitMQ, tambien un servicio publisher con un enpoint usando el framework Fiber y un servicio suscriptor que consume
los mensajes que publica el servicio publisher, usando docker para levantar toda la aplicacion

## QuickStart
Levantar los contenedores del Message broker, el publisher con Fiber y el consumer con docker compose

```
 make run
```
Hacer un request al endpoint del publisher

```
 curl --request GET  --url 'http://localhost:3000/send?msg=publicandoMensajes'
 ```
 abrir el dashboard de rabbitMQ para ver la actividad de la queue en http://localhost:15672/ con usuario y pass => guest
 
 ![image](https://user-images.githubusercontent.com/1621639/172944031-32564d88-5fae-45b7-8f4e-22871159582c.png)
 
 tambien se puede ver en consola la publicacion y consumo de los mensajes
 

![image](https://user-images.githubusercontent.com/1621639/172945007-a6288877-2696-475c-9170-ce8cb1f6ced9.png)
