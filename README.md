# Spam Fighter

## Problem

Nowaday, spammers use bots to make spam calls from a range of thousands
of numbers. However, our phones, iPhones included, do NOT usually support
blocking a range of numbers, but every single number only. There are some
3rd applications to fill the gap. Though, allowing applications to read callers
is really a concern.

## Solution

Therefore, this simple tool is created. Instead of asking users to install any
tools on their mobiles, it merely creates contacts of the spammers with every
possible numbers in widely accepted format of *.vcf. Users can import the vcf
files into their Contacts app to block the spammers.

## Prerequites

- Golang 1.19+

## Usage

```
$ go run cmd/main.go -help
  -name string
        The contact name for the spammer (default "Spammer")
  -number string
        Phone number(s) of the spammer. The number can start with '+'.
        '#' can also be used to substitute any digits form '0' to '9',
        e.g. +84598382### matches all the numbers from +84598382000 to +84598382999.

$ go run cmd/main.go -name=Spammer -number=+84598382### 
DONE! Import Spammer.vcf file to your Contacts app to fight against spammers!
```