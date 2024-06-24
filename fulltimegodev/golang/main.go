package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)

	go getComments(10, respch, wg)
	go getLikes(10, respch, wg)
	go getFriends(10, respch, wg)
	wg.Add(3)
	wg.Wait()
	close(respch)

	userProfile := &UserProfile{}

	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []string:
			userProfile.Comments = msg
		case []int:
			userProfile.Friends = msg
		}
	}

	return userProfile, nil
}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"Hey, that was great",
		"Yeah buddy",
		"Oh, I didnt know that",
	}
	respch <- Response{
		data: comments,
		err:  nil,
	}
	wg.Done()
}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	respch <- Response{
		data: 33,
		err:  nil,
	}
	wg.Done()
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	friendsIds := []int{11, 34, 854, 455}

	respch <- Response{
		data: friendsIds,
		err:  nil,
	}
	wg.Done()
}

type Response struct {
	data any
	err  error
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userProfile)
	fmt.Println("fetching the user profile took ", time.Since(start))
}
