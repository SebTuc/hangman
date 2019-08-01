package hangman

import "strings"

type Game struct {
	State        string   //Game state
	Letters      []string // Letters in the word to finf
	FoundLetters []string //Good guesses
	UsedLetters  []string //Used letters
	TurnsLeft    int      //Remaining attempts
}

func New(turn int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turn,
	}

	return g
}

func (g *Game) MakeAGuess(guess string) {

	guess = strings.ToUpper(guess)

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}

	} else {
		g.State = "badGuess"
		g.DecrementeTurn(guess)
		if hasLost(g.TurnsLeft) {
			g.State = "lost"
		}
	}

}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess

		}
	}
}

func (g *Game) DecrementeTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)

}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}

	return false
}

func hasWon(letters []string, FoundLetters []string) bool {
	for i := range letters {
		if letters[i] != FoundLetters[i] {
			return false
		}
	}

	return true
}

func hasLost(TurnsLeft int) bool {
	if TurnsLeft <= 0 {
		return true
	}

	return false
}
