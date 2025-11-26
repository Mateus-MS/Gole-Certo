package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/routes"
	_ "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/routes"

	_ "github.com/Mateus-MS/Gole-Certo/dev/backend/routes/components"
	_ "github.com/Mateus-MS/Gole-Certo/dev/backend/routes/pages"
	"github.com/joho/godotenv"

	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/middlewares"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	environment := flag.String("env", "dev", "The environment to run")
	flag.Parse()

	app := app.GetInstance()

	app.Router.Use(middlewares.CorsMiddleware(app.Router.Routes))

	startServer(app.Router, *environment)
}

func startServer(router *app.Router, env string) {
	if env == "dev" {
		println("Starting SERVER in DEV mode")
		err := http.ListenAndServe(":3434", router.Handle())
		if err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}
}
