package serialize

import (
  "encoding/json"
  "strings"
)

type Tweet struct {
  CreatedAt twitterTime `json:"created_at"`
  Id        string      `json:"id_str"`
  Text      string
  User      twitterUser
}

type twitterTime string
func (tt *twitterTime) UnmarshalJSON(data []byte) error {
  fields := strings.Fields(string(data))
  date   := fields[2] +"-"+ fields[1]
  *tt     = twitterTime(date)
  return nil
}

type twitterUser struct {
  Name       string
  ScreenName string  `json:"screen_name"`
}

func NewTweet(data []byte) (tweet *Tweet) {
  err := json.Unmarshal(data, &tweet)
  if err != nil { panic(err) }
  return
}
