package main

import (
	"embed"
	"katsan"
	iamServiceHandlers "katsan/iam/presentation/handlers"
	orderServiceHandlers "katsan/order/infrastructure/presentation/handlers"
	traceabilityServiceHandlers "katsan/traceability/presentation/handlers"

	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

//go:embed static/*
var static embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("could not start server: PORT environment variable is required")
	}

	katsan.NewDB(katsan.WithPostgres("postgres", "dbname=katsantestdb user=katsantestuser password=1453Katsan@ host=localhost port=5432 sslmode=disable"))

	r := http.NewServeMux()

	r.Handle("GET /static/", http.FileServerFS(static))

	r.HandleFunc("GET /login", iamServiceHandlers.HandleGetLoginPage)

	r.HandleFunc("GET /oms/accounts", orderServiceHandlers.HandleGetAccountsPage)
	r.HandleFunc("GET /oms/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/oms/accounts/"+r.PathValue("id")+"/account-information", http.StatusSeeOther)
	})
	r.HandleFunc("GET /oms/accounts/{id}/account-information", orderServiceHandlers.HandleGetAccountDetailsAccountInformationPage)
	r.HandleFunc("GET /oms/accounts/{id}/debt-aging-report", orderServiceHandlers.HandleGetAccountDetailsDebtAgingReportPage)

	r.HandleFunc("GET /oms/orders", orderServiceHandlers.HandleGetOrdersPage)
	r.HandleFunc("GET /oms/orders/new", orderServiceHandlers.HandleGetNewOrderPage)

	r.HandleFunc("GET /tm/shipments", traceabilityServiceHandlers.HandleGetShipmentsPage)

	log.Println("starting server on :" + port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
