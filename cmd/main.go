package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"time"
)

const authToken = ""
const token = ""
const channelID = ""

var data struct {
    Token         string `json:"verification_token"`
    First_payment bool   `json:"is_first_subscription_payment"`
    Name          string `json:"from_name"`
    Tier          string `json:"tier_name"`
    Amount        string `json:"amount"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if data.Token != authToken {
		http.Error(w, "Error, incorrect token", http.StatusBadRequest)
		return
	}

	fmt.Println("Received JSON:", data)

	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("JSON received successfully"))

    if data.First_payment != true {
        fmt.Println("not first payment, don't send msg")
        return
    }

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
    



	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Name",
			Value:  data.Name,
			Inline: false,
		},
		{
			Name:   "Amount paid",
			Value:  data.Amount,
			Inline: false,
		},
		{
			Name:   "Tier",
			Value:  data.Tier,
			Inline: false,
		},
	}

	embed := &discordgo.MessageEmbed{
		Title:     "New Subscriber",
		Fields:    fields,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err = dg.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		fmt.Println("error sending Discord message,", err)
		return
	}

}

func main() {
	http.HandleFunc("/kofi", handlePostRequest)

	fmt.Println("Server is listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
