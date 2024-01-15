package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/amonaco/goapi/libs/auth"
	"github.com/amonaco/goapi/libs/config"
	"github.com/amonaco/goapi/libs/logger"

	// "github.com/amonaco/goapi/libs/logger"

	"github.com/amonaco/goapi/libs/util"

	// "github.com/amonaco/goapi/libs/nats"
	// "github.com/amonaco/goapi/libs/websocket"

	authService "github.com/amonaco/goapi/apps/auth"
	"github.com/amonaco/goapi/apps/web"
)

func main() {

	// Config
	config.Read("config/config.yml")
	conf := config.Get()

	// Database
	// database.Start(conf.Postgres)

	// Redis
	// cache.Start()

	// S3 / Minio
	// storage.Setup()

	// Logger
	logger.Init()

	// Worker
	// worker.Init(notification.Handler)

	// Chi
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Auth-Token"},
		ExposedHeaders: []string{},
		// AllowCredentials: true,
		MaxAge: 300,
	})

	r.Use(cors.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// logger.AddMiddleware(r)

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Auth routes
	r.Route("/v1/auth", func(r chi.Router) {
		r.Use(util.RouteToPath)
		r.Handle("/rpc/AuthService/*", authService.New())
	})

	// Web app routes
	r.Route("/v1/web", func(r chi.Router) {
		r.Use(auth.Middleware)
		r.Use(util.RouteToPath)
		// r.Use(auth.Authorize(auth.ConsoleRole))
		// r.Handle("/rpc/CompanyService/*", web.Company())
		r.Handle("/rpc/UserService/*", web.User())
	})

	// Not used at the moment
	// ws := websocket.NewServer()
	// ws.SetConnectionListener(handleWebsocket)
	// r.Mount("/ws", ws)
	// nats.Connect()

	log.Printf("Environment: %v", conf.Environment)
	log.Printf("Listening on %v", conf.Listen)
	log.Fatal(http.ListenAndServe(conf.Listen, r))
}

/*
func handleWebsocket(wsclient *websocket.Client, r *http.Request) *websocket.Error {
	var mode string
	if m, ok := r.URL.Query()["mode"]; ok && len(m) > 0 && len(m[0]) > 0 {
		mode = m[0]
	} else {
		return websocket.NewError("mode needs to be specified", 401)
	}

	var room string
	if r, ok := r.URL.Query()["room"]; ok && len(r) > 0 && len(r[0]) > 0 {
		room = r[0]
	} else {
		return websocket.NewError("room needs to be specified", 401)
	}

	if mode == "master" {
		go wsMasterConn(room, wsclient)
	} else if mode == "slave" {
		go wsSlaveConn(room, wsclient)
	} else {
		return websocket.NewError("invalid mode value", 400)
	}

	return nil
}

func wsMasterConn(room string, ws *websocket.Client) {
	for {
		select {
		case msg := <-ws.Recv():

            // Pulishing to nats
			nats.Publish(room, msg)
			ws.Send([]byte("{\"type\": \"ack\"}"))

		case <-ws.Done():
			return
		}
	}
}

func wsSlaveConn(room string, ws *websocket.Client) {

    // Subscribe to nats
	ch, sub, err := nats.Subscribe(room)

	if err != nil {
		ws.Disconnect()
		return
	}

	for {
		select {
		case msg := <-ch:
			ws.Send(msg.Data)
		case <-ws.Done():
			sub.Unsubscribe()
			return
		}
	}
}
*/
