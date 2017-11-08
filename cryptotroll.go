package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

var api *anaconda.TwitterApi

func main() {
	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api = anaconda.NewTwitterApi(Token, TokenSecret)

	userObj, err := api.GetUsersShow(PersonToCrypto, nil)
	if err != nil {
		log.Println("Error while querying twitter API", err)
		return
	}

	twitterStream := api.PublicStreamFilter(url.Values{"follow": []string{userObj.IdStr}})

	fmt.Println("Stream started, let the stalking commence")

	//twitterStream := api.PublicStreamSample(nil)
	for {
		x := <-twitterStream.C
		switch tweet := x.(type) {
		case anaconda.Tweet:
			response, err := replyToTweet(tweet)
			if err == nil {
				fmt.Println("Replied to following Tweet: " + tweet.User.Name + " " + tweet.FullText)
				fmt.Println("Response: " + response)
			}
			//for now return after one reply until I fix nested replies infinite looping #OOPSMYACCOUNTGOTLOCKED
			return
		case anaconda.StatusDeletionNotice:
			// pass
		default:
			fmt.Printf("unknown type(%T) : %v \n", x, x)
		}
	}

}
