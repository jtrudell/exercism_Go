package twelve

const testVersion = 1

type christmasDay struct {
	day  string
	gift string
}

var verses = []christmasDay{
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

func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}

func Verse(n int) string {
	verse := verses[n-1]
	day := verse.day
	gifts := ""

	if n == 1 {
		gifts = " " + verse.gift
	} else {
		for i := n; i > 0; i-- {
			gifts += appendGift(i)
		}
	}

	return "On the " + day + " day of Christmas my true love gave to me," + gifts + "."
}

func appendGift(n int) string {
	verse := verses[n-1]
	gift := verse.gift
	phrase := " " + gift + ","

	if n == 1 {
		phrase = " and " + gift
	}

	return phrase
}
