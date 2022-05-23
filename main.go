package main

import (
	"flag"
	"fmt"
)
import "github.com/cosmos/cosmos-sdk/types/bech32"

const oldPrefix = "celes"
const newPrefix = "celestia"

func main() {
	var oldAddr string
	flag.StringVar(&oldAddr, "old-addr", "celes1gmhxchj8ltegsc7uz8gr3jn0j50u7903e57khh", "the old addr as string; e.g. celes1q3v5cugc8cdpud87u4zwy0a74uxkk6u454rcq9")
	flag.Parse()

	hrp, data, err := bech32.DecodeAndConvert(oldAddr)
	if err != nil {
		panic(err)
	}
	if hrp != oldPrefix {
		fmt.Printf("got %v, want %v", hrp, oldPrefix)
		return
	}

	res, err := bech32.ConvertAndEncode(newPrefix, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
