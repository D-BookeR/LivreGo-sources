package model

import "fmt"

// Politician contains all data about one given politician
type Politician struct {
	Name  string `json:"name" bson:"name"`
	ID    int    `json:"id,omitempty" bson:"id,omitempty"`
	Party string `json:"party,omitempty" bson:"party,omitempty"`
}

// Politicians is a set of politicians
type Politicians []Politician

// Vote is the information registered when a voter votes
type Vote struct {
	Name         string `json:"name" bson:"name"`
	ID           string `json:"id" bson:"id"`
	PoliticianID int    `json:"politician_id" bson:"politician_id"`
}

// Votes is a set of votes
type Votes []Vote

func (p *Politician) String() string {
	return fmt.Sprintf("%s, de \"%s\"", (*p).Name, (*p).Party)
}

func (v *Vote) String() string {
	return (*v).Name
}
