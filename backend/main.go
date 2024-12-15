package main

import (
	"fmt"
	"net/http"

	filehandler "backend/handlers"

	"github.com/gorilla/mux"
)

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		next.ServeHTTP(w, req)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", filehandler.FileUpload).Methods("POST")

	srv := &http.Server{
		Addr:    ":8090",
		Handler: middleWare(r),
	}

	fmt.Println("Server is running on port 8090")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server error: %v\n", err)
	}

	// Start the server in a goroutine
	// go func() {
	// 	fmt.Println("Server is running on port 8090")
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 			fmt.Printf("Server error: %v\n", err)
	// 	}
	// }()

	// 	// Graceful shutdown
	// 	stop := make(chan os.Signal, 1)
	// 	signal.Notify(stop, os.Interrupt)
	// 	<-stop

	// 	fmt.Println("Shutting down server...")
	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()
	// 	if err := srv.Shutdown(ctx); err != nil {
	// 			fmt.Printf("Shutdown error: %v\n", err)
	// 	}
	// 	fmt.Println("Server stopped")
}
