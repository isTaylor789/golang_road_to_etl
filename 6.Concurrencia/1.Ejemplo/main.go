package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"http://localhost:1234?duration=3s",
	"http://localhost:1234?duration=1s",
	"http://localhost:1234?duration=5s",
}

func main() {
	// Inicia un servidor simulado en un goroutine
	go startMockServer()

	time.Sleep(1 * time.Second) // Espera un momento para que el servidor arranque

	// Ejecuta el ejemplo en modo secuencial y concurrente
	fmt.Println("Fetch Sequential:")
	fetchSequential(urls)

	fmt.Println("\nFetch Concurrent:")
	fetchConcurrent(urls)
}

// fetchSequential procesa las URLs secuencialmente
func fetchSequential(urls []string) {
	for _, url := range urls {
		fetch(url)
	}
}

// fetchConcurrent procesa las URLs concurrentemente
func fetchConcurrent(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fetch(u)
		}(url)
	}

	wg.Wait()
}

// fetch realiza una solicitud HEAD a una URL y maneja errores
func fetch(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Printf("Error fetching URL: %s, error: %v", url, err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Fetched URL: %s, Status Code: %d", url, resp.StatusCode)
}

// startMockServer inicia un servidor HTTP que simula respuestas lentas
func startMockServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		duration := r.URL.Query().Get("duration")
		if duration == "" {
			http.Error(w, "Missing 'duration' parameter", http.StatusBadRequest)
			return
		}

		d, err := time.ParseDuration(duration)
		if err != nil {
			http.Error(w, "Invalid 'duration' parameter", http.StatusBadRequest)
			return
		}

		time.Sleep(d) // Simula un procesamiento lento
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Processed in %s\n", duration)
	})

	log.Println("Mock server running on http://localhost:1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
