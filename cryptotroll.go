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

	timelinePullNoRTsNoReplies, err := api.GetUserTimeline(url.Values{"screen_name": []string{PersonToCrypto}, "include_rts": []string{"false"}, "exclude_replies": []string{"true"}})
	if err != nil {
		log.Println("Error while querying twitter API", err)
		return
	}

	for _, tweets := range timelinePullNoRTsNoReplies {
		tweet, err := api.GetTweet(tweets.Id, url.Values{})
		if err != nil {
			log.Println("Error while querying twitter API", err)
			return
		}
		response, err := replyToTweet(tweet)
		if err == nil {
			fmt.Println("Replied to following Tweet: " + tweet.User.Name + " " + tweet.FullText)
			fmt.Println("Response: " + response)
			return
		}
	}

}
