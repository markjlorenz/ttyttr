package view

func (c *Card) textLines() (lines []string) {
  var split func(text string) (string)

  split = func(text string) string {
    white := ' '
    lastWhite := 0

    var line string
    var splitAt int

    for i, char := range text {
      if char == white {
        lastWhite = i
      }

      if i == c.width {
        // if there was no white space, then we can't do anything
        if lastWhite == 0 {
          splitAt = i
        } else {
          splitAt = lastWhite
        }
        line = split(string(text[splitAt:]))
        lines = append(lines, line)
        return string(text[:splitAt])
      }
    }
    return(text) // short line
  }

  line := split(c.tweet.Text)
  lines = append(lines, line)
  lines = reverse(lines)
  return
}

func reverse(lines []string) (reversed []string) {
  for i, _ := range lines {
    line    := lines[len(lines)-1-i]
    reversed = append(reversed, line)
  }

  return
}

func (c *Card) padLines(lines []string) []string  {
  for i, line := range lines {
    fillLength := (c.width - len(line))/2
    lines[i] = repeat(fillLength) + line + repeat(fillLength)
    if (len(lines[i]) < c.width) {
      lines[i] = lines[i] + FillChar
    }
  }

  return lines
}

func fenceLines(lines []string) []string  {
  for i, line := range lines {
    lines[i] = fence(line)
  }

  return lines
}
