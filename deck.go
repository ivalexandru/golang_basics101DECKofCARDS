package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//extends the behaviour of []string
type deck []string

func newDeck() deck {
	cards := deck{} // empty deck

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//functie ce mere aplicata pe acest custom type
// d deck e recieverul
// d = the actual copy of the deck we're working with
// deck = every variable of type deck can call this func on itself
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// (dec, deck) means we wanna return 2 separate values, both of type deck
// bagi cele 2 returnuri in variabile cand invoci functia,
// ex: hand, remainingCards := deal(cards, 5)
func deal(d deck, handSize int) (deck, deck) {
	//intai returnezi de la inceput pana la nr 'mainii' (not inclusive)
	// apoi returnezi restul pachetului de carti (de la handSize pana la sf  )
	return d[:handSize], d[handSize:]
}

//import 'strings' package
// func Join(elems []string, sep string) string
// Join concatenates the elements of its first argument to create a single string.
// The separator string sep is placed between elements in the resulting string.

// deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// func WriteFile(filename string, data []byte, perm os.FileMode) error
//aka if an error occurs, it will return it
// error is an actual type
// perm (issions) = daca nu exista fisierul, vor fi acordate permisiunile..
func (d deck) saveToFile(fileName string) error {
	//folosesc functia din package respectiv, apoi ca al doilea arg
	// convertesc  in []byte ce rezulta din functia CUSTOM toString()
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// func ReadFile(filename string) ([]byte, error)
// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF
// from Read as an error to be reported.
//daca nu sunt erori, err == nil

// https://golang.org/pkg/os/#Exit

// func Exit(code int)
// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.

// For portability, the status code should be in the range [0, 125].

//we'll convert the []byte into a string then that into a []string then to a deck

// func Split(s, sep string) []string
// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.

//reads from hdd
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1) // 0 e succes, oriceALtceva err
	}
	// string(bs) //ceva,altceva,
	s := strings.Split(string(bs), ",")
	return deck(s) // convert that []string to a deck
}

//shuffle cards - randomize stuff
// https://golang.org/pkg/math/rand/#Intn
// func Intn(n int) int
// Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
// Seeding with the same value results in the same random sequence each run.
// For different numbers, seed with a different value, such as
// time.Now().UnixNano(), which yields a constantly-changing number.

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d { // i = index
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] //swap them
	}
}
