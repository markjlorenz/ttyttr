package main

import (
  "dapplebeforedawn/ttyttr/api"
  "dapplebeforedawn/ttyttr/view"
  "dapplebeforedawn/ttyttr/options"
  "fmt"
)

func main() {
  opts := opts.Options{}
  opts.Parse()

  tweet := api.NewTweet(opts.Tweet, opts.Endpoint, opts.Bearer)
  data  := tweet.Get()
  card  := view.NewCard(data, opts.Width)
  fmt.Println(card)
}
