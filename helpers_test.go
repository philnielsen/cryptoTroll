package main

import (
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

func TestResponseBuildShort(t *testing.T) {
	testResponse := "@" + PersonToCrypto + " CRYPTO WOULD SOLVE THIS: lorem ipsum"
	test := anaconda.Tweet{}
	test.FullText = "lorem ipsum"
	response := buildResponseToTweet(test)
	if response != testResponse {
		t.Errorf("Short Tweet Response was incorrect, got: %s, want: %s", response, testResponse)
	}

}
func TestResponseBuildLong(t *testing.T) {
	testResponse := "@" + PersonToCrypto + " CRYPTO WOULD SOLVE THIS: lorem ipsum"
	test := anaconda.Tweet{}
	test.FullText = "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum socii"
	response := buildResponseToTweet(test)
	if response == testResponse {
		t.Errorf("Short Tweet Response was incorrect, got: %s, want: %s", response, testResponse)
	}
}
