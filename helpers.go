package main

import (
	"log"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func replyToTweet(tweet anaconda.Tweet) {
	var response = "@AdventuresOfFil Hi"
	if response != "" {
		v := url.Values{}
		v.Add("in_reply_to_status_id", strconv.FormatInt(tweet.Id, 10))

		respTweet, err := api.PostTweet(response, v)
		if err != nil {
			log.Println("Error while posting reply", err)
			return
		}

		log.Println("Reply posted : ", respTweet.Text)
	} else {
		log.Println("No response found for tweet : " + tweet.Text)
	}
	return

}
