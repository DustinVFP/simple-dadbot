package main

import (
	"fmt"
	"strings"
	"github.com/andersfylling/disgord"
)

var (
	dadtype string
	conf_Token string
)

func init() {
	// configuration
	
	conf_Token = "NTE2OTQwOTQzMTE1NDE5NjQ4.Duz04g.cGriVLkZNRPGxSmsICG2hRfdz7k"
	dadtype = "dad"

}

func messageDo(session disgord.Session, data *disgord.MessageCreate) {
	// call function to check for im responce
	go imresponce(session, data)
	go kysresoponce(session, data)
	go kmsresoponce(session, data)
	
}

func kysresoponce(session disgord.Session, data *disgord.MessageCreate) {

	msg := data.Message.Content
	
	var doresponce bool
	
	// Check if message contains kys or related 
	
	if strings.Contains(msg, " kys") {
		doresponce = true
	}
	if strings.Contains(msg, " Kys") {
		doresponce = true
	}
	if strings.Contains(msg, " kys ") {
		doresponce = true
	}
	if strings.Contains(msg, " Kys ") {
		doresponce = true
	}
	if strings.Contains(msg, "kys ") {
		doresponce = true
	}
	if strings.Contains(msg, "Kys ") {
		doresponce = true
	}
	if msg == "kys" {
		doresponce = true
	}
	if msg == "Kys" {
		doresponce = true
	}
	
	// If it does then respond
	
	if doresponce {
		responce := fmt.Sprint("thats not very nice, maybe you should follow your own advice before you tell someone to kys")
		data.Message.RespondString(session, responce)
	}
	
}

func kmsresoponce(session disgord.Session, data *disgord.MessageCreate) {

	msg := data.Message.Content
	
	var doresponce bool
	
	// Check if message contains kms or related 
	
	if strings.Contains(msg, " kms") {
		doresponce = true
	}
	if strings.Contains(msg, " Kms") {
		doresponce = true
	}
	if strings.Contains(msg, " kms ") {
		doresponce = true
	}
	if strings.Contains(msg, " Kms ") {
		doresponce = true
	}
	if strings.Contains(msg, "kms ") {
		doresponce = true
	}
	if strings.Contains(msg, "Kms ") {
		doresponce = true
	}
	if msg == "kms" {
		doresponce = true
	}
	if msg == "Kms" {
		doresponce = true
	}
	
	// If it does then respond
	
	if doresponce {
		responce := fmt.Sprint("Please don't kill yourself, not under my roof, its not good for you")
		data.Message.RespondString(session, responce)
	}
	
}

func imresponce(session disgord.Session, data *disgord.MessageCreate) {

	msg := data.Message.Content
	
	var doresponce bool
	var msg2 string
	
	// Check if message starts with i'm or related 
	
	if strings.HasPrefix(msg, "i'm ") {
		doresponce = true
		msg2 = strings.Replace(msg, "i'm ", "", -1)
	}		
	if strings.HasPrefix(msg, "i’m ") {
		doresponce = true
		msg2 = strings.Replace(msg, "i’m ", "", -1)
	}	
	if strings.HasPrefix(msg, "im ") {
		doresponce = true
		msg2 = strings.Replace(msg, "im ", "", -1)
	}	
	if strings.HasPrefix(msg, "i am ") {
		doresponce = true
		msg2 = strings.Replace(msg, "i am ", "", -1)
	}
	if strings.HasPrefix(msg, "I'm ") {
		doresponce = true
		msg2 = strings.Replace(msg, "I'm ", "", -1)
	}	
	if strings.HasPrefix(msg, "I’m ") {
		doresponce = true
		msg2 = strings.Replace(msg, "I’m ", "", -1)
	}	
	if strings.HasPrefix(msg, "Im ") {
		doresponce = true
		msg2 = strings.Replace(msg, "Im ", "", -1)
	}	
	if strings.HasPrefix(msg, "I am ") {
		doresponce = true
		msg2 = strings.Replace(msg, "I am ", "", -1)
	}
	
	// If it does then respond with dad responce
	
	if doresponce {
		responce := fmt.Sprint("hi ", msg2, " i'm ", dadtype)
		data.Message.RespondString(session, responce)
	}
}

func main() {
	// Configure disgords for it to do its shenanigans
	
	session, err := disgord.NewSession(&disgord.Config{
		Token: conf_Token,
		Debug: true, 
	})
	
	// Message handler
	
	session.On(disgord.EventMessageCreate, func(session disgord.Session, data *disgord.MessageCreate) {
		fmt.Println(data.Message.Content)
		user, err := session.GetCurrentUser()
		if err != nil {
			fmt.Println("Error getting current user")
		}
		fmt.Println(user.ID)
		fmt.Println(data.Message.Author)
		if data.Message.Author.ID != user.ID {
			go messageDo(session, data)
			}
	})
	
	// Create Discord Session
	
	err = session.Connect()
	if err != nil {
		fmt.Println("Discord Session error")
		fmt.Println(err.Error())
		panic(err)
	}
	
	// Keep session open until ctl-c is used where it'll then be closed
	
	session.DisconnectOnInterrupt()	
}