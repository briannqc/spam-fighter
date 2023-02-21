package spamfighter

import (
	"errors"

	"github.com/emersion/go-vcard"
)

var (
	ErrEmptyNumber     = errors.New("number is empty")
	ErrInvalidPlusSign = errors.New("'+' can only be the first char")
	ErrUnsupportedChar = errors.New("only '+', digit and '#' char are supported")
)

func CreateCard(name string, numbers []string) (vcard.Card, error) {
	if len(numbers) == 0 {
		return nil, ErrEmptyNumber
	}
	card := vcard.Card{}
	card.AddValue(vcard.FieldVersion, "3.0")

	card.AddName(&vcard.Name{GivenName: name})

	for _, number := range numbers {
		if err := addNumbersOfPatternToCard(card, number); err != nil {
			return nil, err
		}
	}
	return card, nil
}

func addNumbersOfPatternToCard(card vcard.Card, number string) error {
	if len(number) == 0 {
		return ErrEmptyNumber
	}

	for pos, ch := range number {
		if ch == '+' {
			if pos == 0 {
				continue
			}
			return ErrInvalidPlusSign
		}
		if ch != '#' && (ch < '0' || '9' < ch) {
			return ErrUnsupportedChar
		}
	}
	backtrack(card, []byte(number), 0)
	return nil
}

func backtrack(card vcard.Card, pattern []byte, idx int) {
	if idx >= len(pattern) {
		addPhonenumberToCard(card, string(pattern))
		return
	}
	ch := pattern[idx]
	if ch == '+' || ('0' <= ch && ch <= '9') {
		backtrack(card, pattern, idx+1)
		return
	}
	if ch == '#' {
		for d := byte('0'); d <= '9'; d++ {
			pattern[idx] = d
			backtrack(card, pattern, idx+1)
			pattern[idx] = '#'
		}
		return
	}
}

func addPhonenumberToCard(card vcard.Card, number string) {
	cellphone := vcard.Params{}
	cellphone.Add(vcard.ParamType, vcard.TypeCell)
	cellphone.Add(vcard.ParamType, vcard.TypeVoice)

	card.Add(vcard.FieldTelephone, &vcard.Field{Params: cellphone, Value: number})
}
