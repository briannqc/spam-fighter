package spamfighter_test

import (
	"testing"

	"github.com/briannqc/spamfighter"
	"github.com/emersion/go-vcard"
	"github.com/stretchr/testify/assert"
)

func TestCreateCard(t *testing.T) {
	tests := []struct {
		name        string
		numbers     []string
		wantNumbers []string
		wantErr     error
	}{
		{
			name:    "GIVEN empty number THEN error",
			numbers: []string{""},
			wantErr: spamfighter.ErrEmptyNumber,
		},
		{
			name:    "GIVEN '+' in between of number THEN error",
			numbers: []string{"091+123"},
			wantErr: spamfighter.ErrInvalidPlusSign,
		},
		{
			name:    "GIVEN more than one '+' in number THEN error",
			numbers: []string{"+091+123"},
			wantErr: spamfighter.ErrInvalidPlusSign,
		},
		{
			name:    "GIVEN number starts with two '+' THEN error",
			numbers: []string{"++091123"},
			wantErr: spamfighter.ErrInvalidPlusSign,
		},
		{
			name:    "GIVEN unsupported char in number THEN error",
			numbers: []string{"+091123P"},
			wantErr: spamfighter.ErrUnsupportedChar,
		},
		{
			name:        "GIVEN an abs number THEN contains that single number",
			numbers:     []string{"+0912345678"},
			wantNumbers: []string{"+0912345678"},
		},
		{
			name:        "GIVEN an ans number THEN contains that single number",
			numbers:     []string{"0912345678"},
			wantNumbers: []string{"0912345678"},
		},
		{
			name:    "GIVEN a '#' at the ends of number THEN contains 10 numbers",
			numbers: []string{"091234567#"},
			wantNumbers: []string{
				"0912345670", "0912345671", "0912345672", "0912345673", "0912345674",
				"0912345675", "0912345676", "0912345677", "0912345678", "0912345679",
			},
		},
		{
			name:    "GIVEN a '#' in the middle of number THEN contains 10 numbers",
			numbers: []string{"09123456#0"},
			wantNumbers: []string{
				"0912345600", "0912345610", "0912345620", "0912345630", "0912345640",
				"0912345650", "0912345660", "0912345670", "0912345680", "0912345690",
			},
		},
		{
			name:        "GIVEN two abs numbers THEN contains these abs numbers",
			numbers:     []string{"0912345678", "0912345679"},
			wantNumbers: []string{"0912345678", "0912345679"},
		},
		{
			name:    "GIVEN two number patterns THEN contains all numbers of these patterns",
			numbers: []string{"091234567#", "09123456#0"},
			wantNumbers: []string{
				"0912345670", "0912345671", "0912345672", "0912345673", "0912345674",
				"0912345675", "0912345676", "0912345677", "0912345678", "0912345679",
				"0912345600", "0912345610", "0912345620", "0912345630", "0912345640",
				"0912345650", "0912345660", "0912345670", "0912345680", "0912345690",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCard, gotErr := spamfighter.CreateCard("Spammer", tt.numbers)
			if !assert.ErrorIs(t, gotErr, tt.wantErr) {
				return
			}

			gotNumbers := gotCard.Values(vcard.FieldTelephone)
			assert.ElementsMatch(t, tt.wantNumbers, gotNumbers)
		})
	}
}
