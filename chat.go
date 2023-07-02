package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", serveFile)
	app.Get("/ws", websocketHandler)

	log.Fatal(app.Listen(":3000"))
}

func serveFile(c *fiber.Ctx) error {
	return c.SendFile("chat/chat.html")
}

func websocketHandler(c *fiber.Ctx) error {
	return websocket.New(func(conn *websocket.Conn) {
		var (
			msg string
			err error
		)
		for {
			if err = conn.ReadJSON(&msg); err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("recv: %s", msg)

			if err = conn.WriteJSON(msg); err != nil {
				log.Println("write:", err)
				return
			}
		}
	})(c)
}
