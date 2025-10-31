package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

type Tweet struct {
	UserId    int
	TweetId   int
	Timestamp int
}

type Twitter struct {
	UserTweets map[int][]Tweet
	Following  map[int]map[int]bool
	Timestamp  int
}

func DesignTwitterConstructor() Twitter {
	return Twitter{
		UserTweets: map[int][]Tweet{},
		Following:  map[int]map[int]bool{},
		Timestamp:  0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := Tweet{
		UserId:    userId,
		TweetId:   tweetId,
		Timestamp: this.Timestamp,
	}
	this.Timestamp++

	this.UserTweets[userId] = append(this.UserTweets[userId], tweet)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	allTweets := []Tweet{}
	if userTweets, ok := this.UserTweets[userId]; ok {
		allTweets = append(allTweets, userTweets...)
	}

	if followees, ok := this.Following[userId]; ok {
		for followeeId, isFollowing := range followees {
			if isFollowing {
				if followeeTweets, ok := this.UserTweets[followeeId]; ok {
					allTweets = append(allTweets, followeeTweets...)
				}
			}
		}
	}

	sort.Slice(allTweets, func(i, j int) bool {
		return allTweets[i].Timestamp > allTweets[j].Timestamp
	})

	result := make([]int, 0, 10)
	for i := 0; i < 10 && i < len(allTweets); i++ {
		result = append(result, allTweets[i].TweetId)
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.Following[followerId]; !ok {
		this.Following[followerId] = make(map[int]bool)
	}
	this.Following[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.Following[followerId]; ok {
		delete(this.Following[followerId], followeeId)

		if len(this.Following[followerId]) == 0 {
			delete(this.Following, followerId)
		}
	}
}

func RunTestDesignTwitter() {
	obj := DesignTwitterConstructor()
	obj.PostTweet(1, 5)

	res := obj.GetNewsFeed(1)
	if !cmp.EqualSlices([]int{5}, res) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", []int{5}, res)
		os.Exit(1)
	}

	obj.Follow(1, 2)
	obj.PostTweet(2, 6)

	res = obj.GetNewsFeed(1)
	if !cmp.EqualSlices([]int{6, 5}, res) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", []int{6, 5}, res)
		os.Exit(1)
	}

	obj.Unfollow(1, 2)

	res = obj.GetNewsFeed(1)
	if !cmp.EqualSlices([]int{5}, res) {
		fmt.Printf("=== FAILED: expect = %v - got = %v\n", []int{5}, res)
		os.Exit(1)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
