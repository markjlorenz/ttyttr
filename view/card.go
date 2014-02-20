package view

import (
  "dapplebeforedawn/ttyttr/serialize"
  "strings"
)

const TL       string = "⌈"
const TR       string = "⌉"
const BL       string = "⌊"
const BR       string = "⌋"
const HR       string = "-"
const VB       string = "|"
const FillChar string = " "

type Card struct {
  width  int
  tweet  *serialize.Tweet
}

func NewCard(tweet *serialize.Tweet, width int) *Card {
  return &Card{
    width: width,
    tweet: tweet,
  }
}

func (c *Card) String() string {
  lines := []string{
    c.topBorder(),
    fence(c.topLine()),
    fence(c.secondLine()),
    fence(c.blankLine()),
  }

  textLines := fenceLines(c.padLines(c.textLines()))
  lines = append(lines, textLines...)

  lines = append(
    lines,
    fence(c.blankLine()),
    fence(c.bottomLine()),
    c.bottomBorder(),
  )

  return strings.Join(lines, "\n")
}

func (c *Card) topBorder() string {
  return TL + strings.Repeat("-", c.width) + TR
}

func (c *Card) bottomBorder() string {
  return BL + strings.Repeat("-", c.width) + BR
}

func (c *Card) topLine() string {
  username     := c.tweet.User.Name
  createdAt    := c.tweet.CreatedAt
  fillerLength := c.width - len(username) - len(createdAt)

  return username + repeat(fillerLength) + string(createdAt)
}

func (c *Card) secondLine() string {
  screenName     := c.tweet.User.ScreenName
  fillerLength := c.width - len(screenName) - 1

  return "@" + screenName + repeat(fillerLength)
}

func (c *Card) bottomLine() string {
  screenName   := c.tweet.User.ScreenName
  id           := c.tweet.Id
  detail       := "/" + screenName + "/status/" + id
  fillerLength := c.width - len(detail)

  return detail + repeat(fillerLength)
}

func (c *Card)blankLine() string {
  return repeat(c.width)
}

func repeat(length int) string {
  return strings.Repeat(FillChar, length)
}

func fence(line string) string {
  return VB + line + VB
}
