package main

import (
	"log"
	"os"

	"github.com/armon/go-socks5"
	"github.com/spf13/viper"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	users := viper.Get("users").([]interface{})
	creds := socks5.StaticCredentials{}
	for _, u := range users {
		user := u.(map[string]interface{})
		log.Printf("user: %s", user)
		creds[user["user"].(string)] = user["pass"].(string)
	}

	auth := socks5.UserPassAuthenticator{
		Credentials: creds,
	}

	config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{auth},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}

	srv, err := socks5.New(config)
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("listen on 0.0.0.0:" + viper.GetString("port"))
	log.Fatal(srv.ListenAndServe("tcp", ":"+viper.GetString("port")))
}
