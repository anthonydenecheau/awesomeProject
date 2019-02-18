package main

import (
	"github.com/anthonydenecheau/gopocservice/cmd"
	"log"
)

// Librairies
// https://hackernoon.com/the-myth-about-golang-frameworks-and-external-libraries-93cb4b7da50f
// https://www.getrevue.co/profile/golang/issues/writing-a-go-chat-server-the-myths-about-golang-frameworks-much-more-140766
// https://qiita.com/moz450/items/bdd0eb8dff24caa5174a

// https://github.com/sepulsa/rest_echo
// https://github.com/uchonyy/echo-rest-api
// https://github.com/PacktPublishing/Echo-Essentials/tree/master/chapter8
// go get -u github.com/labstack/echo

// Config
// https://www.netlify.com/blog/2016/09/06/creating-a-microservice-boilerplate-in-go/

// Functions
// https://github.com/s1s1ty/Data-Structures-and-Algorithms
func main() {

	// Start with Viper + Cobra
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
