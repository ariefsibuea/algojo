package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("CountMentionsPerUser", RunTestCountMentionsPerUser)
}

/*
 * Problem 			: Count Mentions Per User
 * Topics           : Array, Math, Sorting, Simulation
 * Level            : Medium
 * URL              : https://leetcode.com/problems/count-mentions-per-user
 * Description      : <Description>
 * Examples         : <Examples>
 */

const (
	EventMessage = "MESSAGE"

	MentionAll  = "ALL"
	MentionHere = "HERE"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	sort.Slice(events, func(i, j int) bool {
		timei, _ := strconv.Atoi(events[i][1])
		timej, _ := strconv.Atoi(events[j][1])
		if timei != timej {
			return timei < timej
		}
		return events[i][0] != EventMessage && events[j][0] == EventMessage
	})

	var mentions = make([]int, numberOfUsers)
	var nextOnlineTime = make(map[int]int)

	for _, e := range events {
		eventType := e[0]
		eventTimestamp, _ := strconv.Atoi(e[1])

		if eventType == EventMessage {
			switch e[2] {
			case MentionAll:
				for i := 0; i < numberOfUsers; i++ {
					mentions[i] += 1
				}
			case MentionHere:
				for i := 0; i < numberOfUsers; i++ {
					ot, ok := nextOnlineTime[i]
					if ok && ot > eventTimestamp {
						continue
					}
					delete(nextOnlineTime, i)
					mentions[i] += 1
				}
			default:
				ids := strings.Split(e[2], " ")
				for _, id := range ids {
					mentionedID, _ := strconv.Atoi(id[2:])
					mentions[mentionedID] += 1
				}
			}
		} else {
			offlineID, _ := strconv.Atoi(e[2])
			nextOnlineTime[offlineID] = eventTimestamp + 60
		}
	}

	return mentions
}

func RunTestCountMentionsPerUser() {
	testCases := map[string]struct {
		numberOfUsers int
		events        [][]string
		expect        []int
	}{
		"case-1": {
			numberOfUsers: 2,
			events: [][]string{
				{"MESSAGE", "10", "id1 id0"},
				{"OFFLINE", "11", "0"},
				{"MESSAGE", "71", "HERE"},
			},
			expect: []int{2, 2},
		},
		"case-2": {
			numberOfUsers: 2,
			events: [][]string{
				{"MESSAGE", "10", "id1 id0"},
				{"OFFLINE", "11", "0"},
				{"MESSAGE", "12", "ALL"},
			},
			expect: []int{2, 2},
		},
		"case-3": {
			numberOfUsers: 2,
			events: [][]string{
				{"OFFLINE", "10", "0"},
				{"MESSAGE", "12", "HERE"},
			},
			expect: []int{0, 1},
		},
		"case-4": {
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "5", "HERE"},
				{"OFFLINE", "10", "0"},
				{"MESSAGE", "15", "HERE"},
				{"OFFLINE", "18", "2"},
				{"MESSAGE", "20", "HERE"},
			},
			expect: []int{1, 3, 2},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := countMentions(testCase.numberOfUsers, testCase.events)
		format.PrintInput(map[string]interface{}{"numberOfUsers": testCase.numberOfUsers, "events": testCase.events})

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
