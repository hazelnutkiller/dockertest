package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "i love %s!", r.URL.Path[1:])
}

func pinger(port string) error { //回傳error interface
	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//確認resp server 狀態
	if resp.StatusCode != 200 {
		return fmt.Errorf("server returned not -200 status code")
	}
	return nil
}

func main() {

	//使用.env檔直接改port號
	//godotenv.Load()
	//viper.SetDefault("port", os.Getenv("PORT"))
	//  var getPort =
	//viper.GetString("port")

	//使用flag package
	var port string
	//使用-port
	flag.StringVar(&port, "port", "8080", "server port")
	//使用-p
	flag.StringVar(&port, "p", "8080", "server port")
	flag.Parse()
	//使用os包
	//if p, ok := os.LookupEnv("PORT"); ok {
	//	port = p
	//}

	//healthy check驗證服務一直都是存在
	var ping bool
	flag.BoolVar(&ping, "PORT", false, "check server live")

	if ping { //if ping 為 true
		if err := pinger(port); err != nil {
			log.Printf("ping server error: %v\n", err)
		}
		//下ping時不會去執行http service，所以直接return
		return
	}

	http.HandleFunc("/", handler)
	log.Println("http server run on " + port + " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
