package application

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/a3d21/goddd/domain/account"
)

// TODO: app层

func NewHttpApp(
	accountService account.Service,
) *HttpApp {
	return &HttpApp{
		accountService: accountService,
	}
}

type HttpApp struct {
	accountService account.Service
}

func (a *HttpApp) Run() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"stat": "UP"})
	})

	// 开户
	mux.HandleFunc("/v1/open", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		u, err := a.accountService.Open(context.TODO(), name)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"user": string(u)})
	})

	// 查余额
	// curl http://127.0.0.1:8000/v1/balance?name=xxx
	mux.HandleFunc("/v1/balance", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		u := account.User(name)
		balance, err := a.accountService.Balance(context.TODO(), u)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{"user": balance})
	})

	// 因为只做Demo演示，接口不补全了

	log.Println("listen on :8080")
	return http.ListenAndServe(":8080", mux)
}
