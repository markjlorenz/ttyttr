package view

func (c *Card) textLines() (lines []string) {
  var split func(text string) (string)

  split = func(text string) string {
    for i, _ := range text {
      if i == c.width {
        line := split(string(text[i:]))
        lines = append(lines, line)
        return string(text[:i])
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
    if (len(lines[i]) > c.width) {
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
