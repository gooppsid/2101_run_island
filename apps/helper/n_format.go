package helper

import "golang.org/x/text/message"

func NumberFormat(n int) string {
	p := message.NewPrinter(message.MatchLanguage("id"))
	return p.Sprintf("%d", n)
}
