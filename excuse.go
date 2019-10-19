package p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func sendExcuse(m Message) {
	excuses := []string{
		"I accidentally ate cat food and fear I might die",
		"bats have become tangled in my hair and I cannot get them out",
		"the sobriety device will not allow me to start my car",
		"I had issues getting out of the drive way and now I'm too tired",
		"last night did not go well and I do not believe I can perform today",
		"I made it to the parking lot but had an accident in my pants and had to turn around",
		"I have fallen and I cannot get up",
		"hiccups have overcome me and I cannot function",
		"back snapped, legs tingling",
		"my legs fell asleep on the toilet, and I fell standing up",
		"I accidentally drank mouthwash instead of Gatorade",
		"my cat unplugged my alarm clock",
		"the fortune teller told me I might die if I left the house today",
		"someone has glued my doors and windows shut",
		"I need to go to the hospital to get a candy extracted from my nose",
		"I have accidentally locked myself inside of the house",
		"I have accidentally put petroleum jelly in my eyes",
		"my children changed all of the clocks in the house",
		"I am currently stuck under my bed",
		"all of my underwear are in the washer",
		"my dog has swallowed my keys and I have to wait for them to come out",
		"I am not sure how the solar eclipse will affect me so it's better if I stay home",
		"I became ill from trying too hard at work yesterday",
	}

	i := rand.Intn(len(excuses))
	excuse := excuses[i]

	var response Response

	response.ChatID = m.Message.Chat.ID

	botURL := "https://api.telegram.org/bot" + token + "/sendMessage"

	message := fmt.Sprintf("I cannot make it to work today because %s", excuse)

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
