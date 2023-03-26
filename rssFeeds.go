//go mod init ***
//go mod tidy
//go install github.com/mmcdole/gofeed@latest

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mmcdole/gofeed"
)

func init_feeds(path string) []string {

	feeds := make([]string, 0)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		feed_string := strings.TrimSpace(scanner.Text())
		if !(strings.HasPrefix(feed_string, "http")) {
			feed_string = "http://" + feed_string
		}
		feeds = append(feeds, feed_string)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initialized the following feeds:")
	for _, item := range feeds {
		fmt.Println("-" + item)
	}

	return feeds

}

// gets data from a feed
func get_feedData(feedURL string) *gofeed.Feed {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	fp.UserAgent = "go - RSSparser v. 0.1"
	feed, err := fp.ParseURLWithContext(feedURL, ctx)
	if err != nil {
		log.Fatal(err)
	}
	return feed
}

func main() {

	configfile := flag.String("file", "", "Path to configurationfile containing RSS-feed URLs")
	history := flag.Int("history", 0, "How far back would you like to get feed-items for? default is 0")
	flag.Parse()

	current_time := time.Now()
	//start off with new from the past * hours
	var hist_len time.Duration = time.Duration(*history)
	lastChecked := time.Now().Add(-time.Hour * hist_len)

	color.Cyan("Welcome to the RSS reader for go")
	color.Yellow("Script started at %s", current_time.Format(time.RFC3339))

	feeds_list := init_feeds(*configfile)

	color.Yellow("Now, lets read some feeds\n\n")

	for {
		for _, rss_feed := range feeds_list {
			//feedname := strings.Split(item, ".")[1]
			feedData := get_feedData(rss_feed)
			parsed_feedtime, _ := time.Parse(time.RFC3339, feedData.Updated)
			timetest := lastChecked.Before((parsed_feedtime))
			if timetest {
				color.Magenta("---" + feedData.Title + "---\n")
				for i := 0; i < 5; i++ {
					parsed_itemtime, _ := time.Parse(time.RFC1123, feedData.Items[i].Published)
					item_timetest := lastChecked.Before((parsed_itemtime))
					if item_timetest {
						color.Blue(feedData.Items[i].Published)
						color.Red(feedData.Items[i].Title)
						fmt.Println("")
						fmt.Println(feedData.Items[i].Description)
						fmt.Println("")
					}
				}

			}
		}
		lastChecked = time.Now()
		time.Sleep(60 * time.Second)
	}
}
