package main

import (
	"flag"
	"log"
	"net/http"
	"obas/config"
	"obas/controllers"
	"os"
)

func main() {

	var path = "./src/views/html/"
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	env := &config.Env{
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime),
		InfoLog:  log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		Path:     path,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: env.ErrorLog,
		Handler:  controllers.Controllers(env),
	}

	env.InfoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	error := srv.ListenAndServe()
	env.ErrorLog.Fatal(error)

}
