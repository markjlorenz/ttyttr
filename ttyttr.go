package main

import (
  "dapplebeforedawn/ttyttr/api"
  "fmt"
  // "dapplebeforedawn/ttyttr/options"
)

func main() {
  // opts := opts.Options{}
  // opts.Parse()
  // id       := "436310137683197953"
  id       := "434135040256008192"
  endpoint := "https://api.twitter.com/1.1/statuses/show/"
  bearer   := "AAAAAAAAAAAAAAAAAAAAAGrhWAAAAAAAZzMgE%2BKmTuprORfq%2BDkzDe7d4Hc%3Dc0AYgrQHlsoY2Ag22r4c1RtFfCwSS0d9C0WL063T5QSPZGaS1T"

  tweet := api.NewTweet(id, endpoint, bearer)
  data := tweet.Get()
  fmt.Println(data)
}
