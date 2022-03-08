package matcher

import (
	"gopkg.in/telebot.v3"
	"regexp"
)

var (
	yesRegex = regexp.MustCompile(`(?i)(?:^|[[:punct:]]| )да(?:[[:punct:]]| |$)`)
	noRegex  = regexp.MustCompile(`(?i)(?:^|[[:punct:]]| )нет(?:[[:punct:]]| |$)`)
)

var (
	yesSticker = &telebot.Sticker{
		File: telebot.File{
			FileID:   "CAACAgIAAxUAAWImcK6lsP8FhQHuOeYfzg7uPmBgAAL7CwAC-b_oSd0GhUN-DdtPIwQ",
			UniqueID: "AgAD-wsAAvm_6Ek",
		},
		Width:    512,
		Height:   512,
		Animated: true,
	}

	noSticker = &telebot.Sticker{
		File: telebot.File{
			FileID:   "CAACAgIAAxUAAWImcK5qI8cNBdnTWlEuEYu2s0wVAAIKAAOPzYgSnWx4PfXHHo4jBA",
			UniqueID: "AgADCgADj82IEg",
		},
		Width:    512,
		Height:   512,
		Animated: true,
	}
)

func Matches(text string) *telebot.Sticker {
	if yesRegex.MatchString(text) {
		return yesSticker
	} else if noRegex.MatchString(text) {
		return noSticker
	}

	return nil
}
