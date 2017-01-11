package main

import "flag"

var webHook, botToken, mongoHost string

func init() {
	flag.StringVar(&webHook, "h", "web-hook", "telegram web hook endpoint")
	flag.StringVar(&botToken, "t", "token", "telegram bot access token")
	flag.StringVar(&mongoHost, "m", "localhost", "mongo db hostname")
	flag.Parse()
}

func main() {

}
