package opts

import (
  "flag"
  "fmt"
  "os"
)

type Options struct {
  Endpoint string
  Bearer   string
  Width    int
  Tweet    string
}

func (o *Options) Parse() {
  flag.Usage = usage
  endpoint  := "https://api.twitter.com/1.1/statuses/show/"
  bearer    := "AAAAAAAAAAAAAAAAAAAAAGrhWAAAAAAAZzMgE%2BKmTuprORfq%2BDkzDe7d4Hc%3Dc0AYgrQHlsoY2Ag22r4c1RtFfCwSS0d9C0WL063T5QSPZGaS1T"
  width     := 50

  flag.StringVar(&endpoint, "endpoint", endpoint, "twitter api endpoint")
  flag.StringVar(&bearer, "bearer", bearer, "twitter bearer token")
  flag.IntVar(&width, "width", width, "card width")
  flag.Parse()

  o.Endpoint = endpoint
  o.Bearer   = bearer
  o.Width    = width
  o.Tweet    = flag.Arg(0)
}

func usage() {
  fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
  fmt.Fprintln(os.Stderr, "Takes one position argument, a tweet id")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "Example: ")
  fmt.Fprintln(os.Stderr, "  ttyttr 434135040256008192")
  fmt.Fprintln(os.Stderr, "  requests a tweet from: https://twitter.com/someuser/status/434135040256008192")
}
