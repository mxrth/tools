package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"math"
	"os"
)

//go:generate go run wordlist2go.go eff_large_wordlist
//go:generate go run wordlist2go.go eff_short_wordlist_1
//go:generate go run wordlist2go.go diceware_german

var lst = flag.String("list", "eff", "select wordlist to use. Choose from 'eff', 'eff_short' and 'de'")
var num = flag.Int("n", 5, "number of words to use")
var ent = flag.Bool("entropy", false, "print estimated entropy")

func main() {
	flag.Parse()
	var list []string
	switch *lst {
	case "eff":
		list = eff_large_wordlist
	case "eff_short":
		list = eff_short_wordlist_1
	case "de":
		list = diceware_german
	}
	for i := 1; i <= *num; i++ {
		fmt.Printf("%s", randWord(list))
		if i != *num {
			fmt.Print("-")
		}
	}
	fmt.Println()

	entropy := math.Log2(float64(len(list))) * float64(*num)
	if entropy < 32 {
		fmt.Fprintf(os.Stderr, "WARNING: low passphrase entropy of %.2f bits\n", entropy)
	}
	if *ent {
		fmt.Printf("Estimated entropy: %.2f bits\n", entropy)
	}
}

func randWord(list []string) string {
	bs := make([]byte, 4)
	n, err := rand.Read(bs)
	if err != nil || n != 4 {
		panic(err)
	}
	i := binary.BigEndian.Uint32(bs)
	return list[int(i)%len(list)]
}
