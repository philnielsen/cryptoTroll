package main

import "os"

// ConsumerKey - Twitter API Consumer Key
var ConsumerKey = os.Getenv("ConsumerKey")

// ConsumerSecret - Twitter API Consumer Secret
var ConsumerSecret = os.Getenv("ConsumerSecret")

// Token - Twitter API Access Token
var Token = os.Getenv("Token")

// TokenSecret - Twitter API Access Token Secret
var TokenSecret = os.Getenv("TokenSecret")

// PersonToCrypto - Person's twitter handle we are going to mess with
var PersonToCrypto = os.Getenv("PersonToCrypto")
