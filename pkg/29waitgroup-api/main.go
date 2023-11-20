package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func runMockAPI() {
	http.HandleFunc("/endpoint1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("endpoint1 OK"))
	})
	http.HandleFunc("/endpoint2", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("endpoint2 OK"))
	})
	http.HandleFunc("/endpoint3", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("endpoint3 OK"))
	})
	http.ListenAndServe(":8080", nil)
}

func fetchAPI(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	// この関数が終了したら必ずレスポンスを閉じる
	defer resp.Body.Close()

	//ReadAllでResponse Bodyを読み切る
	// (keepAliveできずにコネクションが再利用されずに終了してしまうのを防ぐため)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// ここでレスポンスを処理
	results <- fmt.Sprintf("%s: %s", url, string(b))
}

func main() {
	go runMockAPI()

	apiURLs := []string{
		"http://localhost:8080/endpoint1",
		"http://localhost:8080/endpoint2",
		"http://localhost:8080/endpoint3",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(apiURLs))

	for _, url := range apiURLs {
		wg.Add(1)
		go fetchAPI(url, &wg, results)
	}

	// すべてのゴルーチンが完了するのを待つ
	wg.Wait()
	close(results)

	// 結果を収集
	for result := range results {
		fmt.Println(result)
	}
}
