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
		number      string
		wantNumbers []string
		wantErr     error
	}{
		{
			name:    "GIVEN empty number THEN error",
			number:  "",
			wantErr: spamfighter.ErrEmptyNumber,
		},
		{
			name:    "GIVEN '+' in between of number THEN error",
			number:  "091+123",
			wantErr: spamfighter.ErrInvalidPlusSign,
		},
		{
			name:    "GIVEN more than one '+' in number THEN error",
			number:  "+091+123",
			wantErr: spamfighter.ErrInvalidPlusSign,
		},
		{
			name:    "GIVEN number starts with two '+' THEN error",
			number:  "++091123",
			wantErr: spamfighter.ErrInvalidPlusSign,
		},
		{
			name:    "GIVEN unsupported char in number THEN error",
			number:  "+091123P",
			wantErr: spamfighter.ErrUnsupportedChar,
		},
		{
			name:        "GIVEN a single number THEN contains that single number",
			number:      "+0912345678",
			wantNumbers: []string{"+0912345678"},
		},
		{
			name:        "GIVEN a single number THEN contains that single number",
			number:      "0912345678",
			wantNumbers: []string{"0912345678"},
		},
		{
			name:   "GIVEN a '#' at the ends of number THEN contains 10 numbers",
			number: "091234567#",
			wantNumbers: []string{
				"0912345670", "0912345671", "0912345672", "0912345673", "0912345674",
				"0912345675", "0912345676", "0912345677", "0912345678", "0912345679",
			},
		},
		{
			name:   "GIVEN a '#' in the middle of number THEN contains 10 numbers",
			number: "09123456#0",
			wantNumbers: []string{
				"0912345600", "0912345610", "0912345620", "0912345630", "0912345640",
				"0912345650", "0912345660", "0912345670", "0912345680", "0912345690",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCard, gotErr := spamfighter.CreateCard("Spammer", tt.number)
			if !assert.ErrorIs(t, gotErr, tt.wantErr) {
				return
			}

			gotNumbers := gotCard.Values(vcard.FieldTelephone)
			assert.ElementsMatch(t, tt.wantNumbers, gotNumbers)
		})
	}
}
