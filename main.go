package voluryashop

import (
	"log"
	"net/http"

	"github.com/Iagobarros211256/voluryashop/config"
	"github.com/Iagobarros211256/voluryashop/services"
)

func main() {
	// 1. Carregar configuração com Viper
	cfg := config.LoadConfig()

	// 2. Inicializar repositórios (DB fake aqui só pra ilustrar)
	productRepo := repositories.NewProductRepository()
	orderRepo := repositories.NewOrderRepository()

	// 3. Inicializar serviços
	orderService := services.NewOrderService(orderRepo, productRepo)

	// 4. Inicializar handlers
	orderHandler := handlers.NewOrderHandler(orderService)

	// 5. Definir rotas
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", orderHandler.CreateOrder)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// 6. Rodar servidor na porta da config
	addr := ":" + cfg.ServerPort
	log.Printf("🌐 Servidor rodando em %s [env=%s]", addr, cfg.Env)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
