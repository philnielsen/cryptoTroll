package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func buildResponseToTweet(tweet anaconda.Tweet) string {
	var buffer bytes.Buffer
	buffer.WriteString("@" + PersonToCrypto + " ")

	//Check Length of tweet to make sure I can add in the entertaining text + full tweet text, otherwise Bark at the offender
	if len(tweet.FullText) <= 114 {
		buffer.WriteString("CRYPTO WOULD SOLVE THIS: " + tweet.FullText)
	} else {
		log.Println("Tweet too long, respond with *BARK* CRYPTO *BARK*")
		buffer.WriteString("*BARK* CRYPTO *BARK*")
	}

	//Unfortunatly, most tweets won't support this length of troll, might add back in the future.
	// buffer.WriteString("@" + PersonToCrypto + " ")
	// splitTweetText := strings.Fields(tweet.FullText)

	// log.Println(len(tweet.FullText))

	// for _, tweetWord := range splitTweetText {
	// 	buffer.WriteString(tweetWord + " " + "Crypto" + " ")
	// }

	response := buffer.String()

	log.Println(response)
	return response
}

func replyToTweet(tweet anaconda.Tweet) (string, error) {
	response := buildResponseToTweet(tweet)
	dryRun := false
	if response != "" && dryRun != true {
		v := url.Values{}
		v.Add("in_reply_to_status_id", strconv.FormatInt(tweet.Id, 10))

		respTweet, err := api.PostTweet(response, v)
		if err != nil {
			log.Println("Error while posting reply", err)
			return "ERROR", err
		}

		log.Println("Reply posted : ", respTweet.Text)
	} else if dryRun == true {
		log.Println("DryRun Response: " + response)
	} else {
		log.Println("No response built for tweet : " + tweet.Text)
	}
	return response, nil

}

//this function pulls all tweets from a user that aren't retweets or replies to another tweet
func pullTimeline(user anaconda.User) {
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
