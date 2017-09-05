package house

const testVersion = 1

var verses = [][]string{
	{"horse and the hound and the horn", ""},
	{"farmer sowing his corn", "belonged to"},
	{"rooster that crowed in the morn", "kept"},
	{"priest all shaven and shorn", "woke"},
	{"man all tattered and torn", "married"},
	{"maiden all forlorn", "kissed"},
	{"cow with the crumpled horn", "milked"},
	{"dog", "tossed"},
	{"cat", "worried"},
	{"rat", "killed"},
	{"malt", "ate"},
	{"house that Jack built.", "lay in"},
}

func Verse(num int) string {
	verse := "This is the " + verses[len(verses)-num][0]
	for i := num - 1; i > 0; i-- {
		verse += "\nthat " + verses[len(verses)-i][1] + " the " + verses[len(verses)-i][0]
	}
	return verse
}

func Song() string {
	var song string
	for i := 1; i <= len(verses); i++ {
		song += Verse(i)
		if i < len(verses) {
			song += "\n\n"
		}
	}
	return song
}
