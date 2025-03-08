package application

import (
	"log"
	"os"

	"mi-notificacion/src/notifications/domain/entities"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var wsConnection *websocket.Conn

func init() {
	
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}
}

func SendNotification(notification entities.Notification) error {
	if wsConnection == nil {
		err := connectToWebSocketServer()
		if err != nil {
			log.Printf("Error al conectar con el WebSocket server: %v", err)
			return err
		}
	}

	err := sendMessageToWebSocketServer(notification.Message)
	if err != nil {
		log.Printf("Error al enviar mensaje al WebSocket server: %v", err)
		return err
	}

	return nil
}

func connectToWebSocketServer() error {
	var err error

	wsURL := os.Getenv("WS_SERVER_URL")
	
	wsConnection, _, err = websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		return err
	}
	log.Println("Conexión WebSocket establecida con éxito")
	return nil
}

func sendMessageToWebSocketServer(message string) error {
	if wsConnection == nil {
		err := connectToWebSocketServer()
		if err != nil {
			return err
		}
	}

	err := wsConnection.WriteJSON(map[string]string{"message": message})
	if err != nil {
		log.Printf("Error enviando mensaje: %v", err)
		wsConnection = nil
		return err
	}

	log.Printf("Mensaje enviado al WebSocket server: %s", message)
	return nil
}
