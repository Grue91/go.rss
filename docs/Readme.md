# go.rss #

RSS-feed parser for the command line, written in Go  

![the picture](./img_1.png)


## Usage

Specify which RSS-feeds to read with a configuration file (-file), and optionally how far back you want to load items for with the "-history" option

```go
go run rssFeeds.go -h
  -file string
        Path to configuration file containing RSS-feed URLs
  -history int
        How far back would you like to get feed-items for? default is 0

//Run with
go run rssFeeds.go -file feeds.txt -history 1

//To build executable for your current platform, use:
go build rssFeeds.go 

```

## Dependencies
Uses these great modules:  
https://github.com/fatih/color  
https://github.com/mmcdole/gofeed