package p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Quote struct {
	ID      string
	Content string
	Author  string
}

func getQuote() Quote {
	var quote Quote

	req, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Println(err)
	}

	return quote
}

func sendQuote(m Message) {
	log.Println("setting up message")
	var response Response

	response.ChatID = m.Message.Chat.ID

	botURL := "https://api.telegram.org/bot" + token + "/sendMessage"

	quote := getQuote()

	message := fmt.Sprintf("%s\n - %s", quote.Content, quote.Author)

	response.Text = message

	var body []byte

	body, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))

	req, err := http.Post(botURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

}
