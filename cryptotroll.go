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

	timelinePullNoRTsNoReplies, _ := api.GetUserTimeline(url.Values{"screen_name": []string{PersonToCrypto}, "include_rts": []string{"false"}, "exclude_replies": []string{"true"}})

	for _, tweets := range timelinePullNoRTsNoReplies {
		//fmt.Println(tweets.Id)
		tweet, err := api.GetTweet(tweets.Id, url.Values{})
		if err != nil {
			log.Println("Error while querying twitter API", err)
			return
		}
		replyToTweet(tweet)
		fmt.Println("replied to" + tweet.FullText)
	}

	// searchResult, _ := api.GetSearch("golang", nil)
	// for _, tweet := range searchResult.Statuses {
	// 	fmt.Println(tweet.Text)
	// }

}
