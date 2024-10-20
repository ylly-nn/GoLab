package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var wg sync.WaitGroup
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// TCP-сервер для обработки текстовых сообщений
func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("TCP сервер получил сообщение: %s", message)
	conn.Write([]byte("TCP сервер: сообщение получено\n"))
}

func startTCPServer() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Ошибка запуска TCP сервера:", err)
	}
	defer ln.Close()

	fmt.Println("TCP сервер запущен на порту 8081")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Ошибка соединения TCP:", err)
			continue
		}
		wg.Add(1)
		go func(conn net.Conn) {
			defer wg.Done()
			handleTCPConnection(conn)
		}(conn)
	}
}

// HTTP-сервер для обработки GET и POST запросов
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		fmt.Printf("HTTP сервер получил данные: %v\n", data)
		fmt.Fprintln(w, "Данные приняты")
	} else {
		http.Error(w, "Только POST метод поддерживается", http.StatusMethodNotAllowed)
	}
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func startHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/data", dataHandler)

	loggedMux := logger(mux)

	fmt.Println("HTTP сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}

// WebSocket-сервер для чата
func handleWSConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Ошибка чтения сообщения WS: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleWSMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Ошибка отправки сообщения WS: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func startWebSocketServer() {
	http.HandleFunc("/ws", handleWSConnections)

	go handleWSMessages()

	fmt.Println("WebSocket сервер запущен на порту 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	// Обработка сигналов для graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Запуск TCP-сервера
	go startTCPServer()

	// Запуск HTTP-сервера
	go startHTTPServer()

	// Запуск WebSocket-сервера
	go startWebSocketServer()

	// Ожидание сигнала завершения
	<-stop
	fmt.Println("\nЗавершаем работу серверов...")

	// Завершаем все горутины
	wg.Wait()
	fmt.Println("Серверы завершили работу")
}
