package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tranthienphucbreak/parsleytest/cmd/webapp/config"
	"github.com/tranthienphucbreak/parsleytest/cmd/webapp/config/sqlite"
	"github.com/tranthienphucbreak/parsleytest/cmd/webapp/routes"
	"github.com/tranthienphucbreak/parsleytest/internal/patient"
)

func main() {
	r := chi.NewRouter()
	patientService := Initialize()

	routes := routes.NewServicesHandler(patientService)
	routes.Set(r)

	config := config.ReadConfig()

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
	if err != nil {
		panic(err)
	}
}

func Initialize() *patient.PatientService {
	return &patient.PatientService{
		DBService: sqlite.NewSqliteService(),
	}
}
