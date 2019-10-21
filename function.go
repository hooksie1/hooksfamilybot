// Package p contains an HTTP Cloud Function.
package p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Message struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

type Response struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type Family struct {
	Name     string
	Pronoun  string
	Pronoun2 string
}

var token = os.Getenv("TOKEN")
var port = os.Getenv("PORT")
var url = os.Getenv("URL")

func getMember() Family {
	members := []Family{
		{"Nathanael", "his", "he"},
		{"Emily", "her", "she"},
		{"Dad", "his", "he"},
		{"Mom", "her", "she"},
		{"Bre", "her", "she"},
		{"John", "his", "he"},
		{"Bentley", "her", "she"},
		{"Buddy", "his", "he"},
		{"Scarlett", "her", "she"},
	}

	member := members[rand.Intn(len(members))]

	return member

}

func sendMessage(m Message) {
	log.Println("setting up message")
	var response Response

	response.ChatID = m.Message.Chat.ID

	botURL := "https://api.telegram.org/bot" + token + "/sendMessage"

	member := getMember()

	message := fmt.Sprintf("%s's a weirdo just look at %s beardo, %s's on drugs and %s eats bugs.", member.Name, member.Pronoun, member.Pronoun2, member.Pronoun2)

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

func Bot(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	log.Println("started")

	var message Message

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &message)

	log.Println(string(body))
	log.Println(message)

	if message.Message.Text == "/weirdo" || message.Message.Text == "/weirdo@hooksfamilybot" {
		sendMessage(message)
	}

	if message.Message.Text == "/excuse" || message.Message.Text == "/excuse@hooksfamilybot" {
		sendExcuse(message)
	}

	if message.Message.Text == "/quote" || message.Message.Text == "/quote@hooksfamilybot" {
		sendQuote(message)
	}

}
