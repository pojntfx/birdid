package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

type connection struct {
	from  string
	to    string
	tweet twitter.Tweet
}

func main() {
	clientID := flag.String("client-id", "", "Twitter API client ID (can also be set using the CLIENT_ID env variable)")
	clientSecret := flag.String("client-secret", "", "Twitter API client secret (can also be set using the CLIENT_SECRET env variable)")

	candidateOne := flag.String("candidate-one", "", "First candidate's Twitter handle")
	candidateTwo := flag.String("candidate-two", "", "Second candidate's Twitter handle")

	limit := flag.Int("limit", 1000, "Maximum amount of events to look back in timeline")

	verbose := flag.Bool("verbose", false, "Enable verbose logging")

	flag.Parse()

	if *clientID == "" {
		*clientID = os.Getenv("CLIENT_ID")
	}

	if *clientSecret == "" {
		*clientSecret = os.Getenv("CLIENT_SECRET")
	}

	if *candidateOne == "" {
		panic("empty first candidate's Twitter handle")
	}

	if *candidateTwo == "" {
		panic("empty second candidate's Twitter handle")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &clientcredentials.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := config.Client(ctx)

	client := twitter.NewClient(httpClient)

	tweets := []connection{}
	for candidateOne, candidateTwo := range map[string]string{*candidateOne: *candidateTwo, *candidateTwo: *candidateOne} {
		maxID := int64(math.MaxInt64 - 1)
		sinceID := int64(0)
		curr := 0

		var earliestCandidate twitter.Tweet

		for maxID > sinceID {
			if curr >= *limit {
				break
			}

			results, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
				ScreenName: candidateOne,
				SinceID:    sinceID,
				MaxID:      maxID,
			})
			if err != nil {
				panic(err)
			}

			curr += len(results)

			newMaxID := maxID
			for _, tweet := range results {
				if tweet.ID < newMaxID {
					newMaxID = tweet.ID
				}

				if *verbose {
					log.Println("Evaluating tweet candidate for connection", candidateOne, "to", candidateTwo+":", tweet.ID, tweet.CreatedAt, tweet.Text)
				}

				if earliestCandidate.ID == 0 || earliestCandidate.ID > tweet.ID {
					if tweet.InReplyToScreenName == candidateTwo {
						earliestCandidate = tweet

						continue
					}

					if tweet.RetweetedStatus != nil && tweet.RetweetedStatus.InReplyToScreenName == candidateTwo {
						earliestCandidate = tweet

						continue
					}

					if strings.Contains(tweet.Text, candidateTwo) {
						earliestCandidate = tweet

						continue
					}
				}
			}

			if newMaxID == maxID {
				break
			}

			maxID = newMaxID
		}

		tweets = append(tweets, connection{
			from:  candidateOne,
			to:    candidateTwo,
			tweet: earliestCandidate,
		})
	}

	for _, tweet := range tweets {
		fmt.Printf(
			"Earliest tweet from %v to %v: ID %v at %v with URL %v and text %v\n",
			tweet.from,
			tweet.to,
			tweet.tweet.ID,
			tweet.tweet.CreatedAt,
			"https://twitter.com/"+path.Join(tweet.from, "status", tweet.tweet.IDStr),
			tweet.tweet.Text,
		)
	}
}
