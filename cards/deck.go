package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heards", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/* Isso antes do nome da função é um receiver.
Estamos dizendo que qualquer variável do tipo deck
pode utilizar essa função. Como se fose um método.
Em uma analogia com js, é o contexto: o this */
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	/* Aqui criamos uma instância de um Rand
	com um seed baseado no unix time stamp
	para poder ter um valor randômico e diferente
	a cada execução	*/
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	dLenth := len(d)
	for i := range d {
		/* Esse rand.Intn depende de um seed, e por padrão
		o Go utiliza SEMPRE O MESMO SEED. Ou seja, é randômico,
		mas o mesmo resultado a cada execução do programa.
		newPosition := rand.Intn(dLenth - 1) */
		newPosition := r.Intn(dLenth - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
