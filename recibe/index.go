package main

import (
	"encoding/json"
	"log"
	"recibed/entities"
	"recibed/src"

	"github.com/joho/godotenv"
)

func main() {
  // Cargar las variables de entorno
  godotenv.Load()
  rabbit := src.NewRabbitMQ()
  
  // Tratamiento de un mensaje
  msgs := rabbit.GetMessages()
  var forever chan struct{}

  go func() {
    for d := range msgs {
        var notify entities.Notify
        err := json.Unmarshal(d.Body, &notify)
        if err != nil {
            log.Printf("Error al decodificar el mensaje: %s", err)
            continue
        }
        log.Printf(" [x] Prestamo recibido: id_reader=%d, return_date=%s", notify.Id_reader, notify.Return_date)
        
        if notify.Return_date != "null" {
          src.FetchAPI(notify, "lend")
        } else {
		      src.FetchAPI(notify, "return")
        } 
    }
}()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}


