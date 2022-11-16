package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "i love %s!", r.URL.Path[1:])
}

func main() {

	//使用.env檔直接改port號
	godotenv.Load()
	viper.SetDefault("port", os.Getenv("PORT"))
	//  var getPort =
	viper.GetString("port")

	//使用flag package
	var port string
	//使用-port
	flag.StringVar(&port, "port", "8080", "server port")
	//使用-p
	flag.StringVar(&port, "p", "8080", "server port")
	flag.Parse()
	//使用os包
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	http.HandleFunc("/", handler)
	log.Println("http server run on " + port + " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
