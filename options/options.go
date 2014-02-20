// package opts
//
// import (
//   "flag"
//   "strings"
//   "fmt"
//   "os"
// )
//
// type Options struct {
//   Blueprint []string
// }
//
// func (o *Options) Parse() {
//   flag.Usage = usage
//   default_blueprint := "PRP VBD NN IN NNS DT NN VBG PRP$ NN"
//   var blueprint string
//
//   flag.StringVar(&blueprint, "blueprint", default_blueprint, "the sentence structure to be output")
//   flag.Parse()
//   o.Blueprint = strings.Fields(blueprint)
// }
//
// func usage() {
//   fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
//   fmt.Fprintln(os.Stderr, "Feed the corpus into stdin.")
//   flag.PrintDefaults()
//   fmt.Fprintln(os.Stderr, "Example: ")
//   fmt.Fprintln(os.Stderr, "  talky < corpus.txt")
// }
