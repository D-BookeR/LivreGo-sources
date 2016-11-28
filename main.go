package main

import "fmt"

// Politician represents one given politician
type Politician int

const (
	// BlankVote represents a blank vote
	BlankVote Politician = iota
	// VoteChantalRouleau represents a vote for Chantal Rouleau
	VoteChantalRouleau
	// VoteLoraineLagace represents a vote for Loraine Lagace
	VoteLoraineLagace
	// VoteJeanBaptisteEricDorion represents a vote fore JBE Dorion
	VoteJeanBaptisteEricDorion
)

var politicianList = []string{
	"Blank vote",
	"Chantal ROULEAU",
	"Loraine LAGACÉ",
	"Jean-Baptiste-Éric DORION",
}

func (p *Politician) String() string {
	return politicianList[*p]
}

// Elector represents an elector
type Elector struct {
	Name string
}

// Election represents an election
type Election struct {
	Title       string
	Politicians []Politician
	Electors    []Elector
	round       int
}

// NewElection is creating a new election
// with a title, some electors and politicians of cours
// this will returns an Election
func NewElection(title string, el []Elector) *Election {
	return &Election{
		Title:       title,
		Politicians: []Politician{},
		Electors:    el,
		round:       0,
	}
}

// GetResult displays the result of a given election
func (e *Election) GetResult(v []Vote) string {
	//TODO : implement this function
	return ""
}

// Vote represents a vote
type Vote struct {
	IDUser       int
	IDPolitician Politician
}

func (e *Election) String() string {
	if e.round == 0 {
		return fmt.Sprintf("Les deux candidats finalistes sont : %+v", e.Politicians)
	}
	return fmt.Sprintf("Votre président(e) est : %+v", e.Politicians)
}

func main() {

	electorList := []Elector{
		{"Henri LEPIC"},
		{"Mathilde MEILICHZON"},
		{"Mathieu COUPE"},
		{"Ramzi SAYAL"},
		{"Rober MENARD"},
		{"Benjamin DEON"},
		{"Mireille GAILLARD"},
		{"Olivier VASELIN"},
		{"Antoine TARRE"},
	}

	e := NewElection("Présidentielles 2017", electorList)

	round1 := []Vote{
		{0, VoteChantalRouleau},
		{1, VoteJeanBaptisteEricDorion},
		{2, VoteJeanBaptisteEricDorion},
		{3, VoteLoraineLagace},
		{4, BlankVote},
		{5, VoteJeanBaptisteEricDorion},
		{6, VoteLoraineLagace},
		{7, BlankVote},
		{8, VoteLoraineLagace},
		{9, VoteJeanBaptisteEricDorion},
	}

	fmt.Println(e.GetResult(round1))

	round2 := []Vote{
		{0, VoteJeanBaptisteEricDorion},
		{1, VoteJeanBaptisteEricDorion},
		{2, VoteJeanBaptisteEricDorion},
		{3, VoteLoraineLagace},
		{4, VoteLoraineLagace},
		{5, VoteJeanBaptisteEricDorion},
		{6, VoteLoraineLagace},
		{7, VoteLoraineLagace},
		{8, VoteLoraineLagace},
		{9, VoteJeanBaptisteEricDorion},
	}

	fmt.Println(e.GetResult(round2))
}
