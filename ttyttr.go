package main

import (
  "dapplebeforedawn/ttyttr/api"
  "dapplebeforedawn/ttyttr/options"
  "fmt"
)

func main() {
  opts := opts.Options{}
  opts.Parse()

  // id       := "436310137683197953"
  // id       := "434135040256008192"

  tweet := api.NewTweet(opts.Tweet, opts.Endpoint, opts.Bearer)
  data  := tweet.Get()
  fmt.Println(data)
}
