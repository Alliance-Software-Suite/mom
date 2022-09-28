package main

import flag "github.com/spf13/pflag"

func main() {
	flag.StringP("resource", "r", "MONEY", "Resource to send")
	flag.Float64P("amount", "a", 0, "Amount to send")

	flag.IntP("nation", "n", 0, "Nation to send to (cannot be used with alliance)")
	flag.IntP("alliance", "l", 0, "Alliance to send to (cannot be used with nation)")

	flag.Parse()
}
