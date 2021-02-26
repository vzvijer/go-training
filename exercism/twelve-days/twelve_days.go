package twelve

import (
	"fmt"
	"strings"
)

type Gift struct {
	day  string
	name string
}

var gifts = []Gift{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func Verse(n int) string {
	gs := ""
	for i := 0; i < n; i++ {
		if i == 0 {
			gs = fmt.Sprintf("%s", gifts[i].name)
		} else if i == 1 {
			gs = fmt.Sprintf("%s, and %s", gifts[i].name, gs)
		} else {
			gs = fmt.Sprintf("%s, %s", gifts[i].name, gs)
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", gifts[n-1].day, gs)
}

func Song() string {
	song := strings.Builder{}
	for i := 1; i <= len(gifts); i++ {
		song.WriteString(Verse(i))
		if i < len(gifts) {
			song.WriteString("\n")
		}
	}
	return song.String()
}
