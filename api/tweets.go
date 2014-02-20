package api

import (
  "dapplebeforedawn/ttyttr/serialize"
  "net/http"
  "bytes"
)

type Tweet struct {
  client   *http.Client
  Request  *http.Request
  Response *http.Response
}

func NewTweet(tweetId string, endpoint string, bearer string) *Tweet {
  contentType := ".json"
  client := &http.Client{}
  req, err := http.NewRequest("GET", endpoint+tweetId+contentType, nil)
  if err != nil { panic(err) }

  req.Header.Add("Authorization", "Bearer "+bearer)
  return &Tweet{
    client:  client,
    Request: req,
  }
}

func (t *Tweet) Get() *serialize.Tweet {
  resp, err := t.client.Do(t.Request)
  if err != nil { panic(err) }

  data   := new(bytes.Buffer)
  _, err  = data.ReadFrom(resp.Body)
  if err != nil { panic(err) }

  return serialize.NewTweet(data.Bytes())
}


