package main


import "flag"

var output string


func init() {
    flag.StringVar(&output, "o", "all", "display name (output)")
}

func main() {
	flag.Parse()

    Run()
}
