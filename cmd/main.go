package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/briannqc/spamfighter"
	"github.com/emersion/go-vcard"
)

func main() {
	name := flag.String("name", "Spammer", "The contact name for the spammer")
	numberPatterns := flag.String(
		"numbers",
		"",
		`Comma separated phone numbers of the spammer. Each number can start with '+'.
'#' can also be used to substitute any digits form '0' to '9',
e.g. +84598382### matches all the numbers from +84598382000 to +84598382999;
+845983824##,+845983826## matches all the numbers of +84598382400 ~ +84598382499 and +84598382600 ~ +84598382699`)

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Create a vcf files of the spammers to fight against them!")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr,
			`Nowaday, spammers use bots to make spam calls from a range of thousands of numbers. However, our phones, iPhones
included, do NOT usually support blocking a range of numbers, but every single number only. There are some 3rd
applications to fill the gap. Though, allowing applications to read callers is really a concern. Therefore,
this simple tool is created. Instead of asking users to install any tools on their mobiles, it merely creates
contacts of the spammers with every possible numbers in widely accepted format of *.vcf. Users can import the
vcf files into their Contacts app to block the spammers.`)

		flag.PrintDefaults()
	}
	flag.Parse()

	card, err := spamfighter.CreateCard(*name, strings.Split(*numberPatterns, ","))
	if err != nil {
		log.Fatal("Creating a vCard failed", err)
	}

	filename := fmt.Sprintf("%s.vcf", *name)
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("Creating contact file failed", err)
	}

	if err := vcard.NewEncoder(f).Encode(card); err != nil {
		log.Fatal("Writing contact file failed", err)
	}
	_ = f.Close()
	log.Printf("DONE! Import %v file to your Contacts app to fight against spammers!", filename)
}
