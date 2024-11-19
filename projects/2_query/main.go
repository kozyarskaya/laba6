package main

import (
	"fmt"
	"io"
	"net/http" // пакет для поддержки HTTP протокола
	"os"      
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	itog := "Hello," + name + "!"
	w.Write([]byte(itog))
}

func main() {
	http.HandleFunc("/api/user", handler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
