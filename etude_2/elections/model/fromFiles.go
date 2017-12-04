package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

// FromFiles holds the information and instruments the liaison of the model with flat JSON files
type FromFiles struct {
	DirPath             string   // le répertoire dans lequel tous le fichiers sont
	PoliticiansFileName string   // le seul fichier dans lequel on trouve les politiciens
	VotesFileNames      []string // tous les fichiers contenant des votes
}

// Politician contains all data about one given politician
type Politician struct {
	Name  string `json:"name"`
	ID    int    `json:"id,omitempty"`
	Party string `json:"party,omitempty"`
}

// Politicians is a set of politicians
type Politicians []Politician

// Vote is the information registered when a voter votes
type Vote struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	PoliticianID int    `json:"politician_id"`
}

// Votes is a set of votes
type Votes []Vote

func (p *Politician) String() string {
	return fmt.Sprintf("%s, de \"%s\"", (*p).Name, (*p).Party)
}

func (v *Vote) String() string {
	return v.Name
}

var allPoliticians Politicians

// AllPoliticians fetches all politicians from JSON file if cache is empty, returns the cache otherwise
func (m *FromFiles) AllPoliticians() (Politicians, error) {
	if allPoliticians == nil {
		file, err := ioutil.ReadFile(path.Join(m.DirPath, m.PoliticiansFileName))
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(file, &allPoliticians); err != nil {
			return nil, err
		}
	}
	return allPoliticians, nil
}

// PoliticianFromID finds the Politician value when given its ID
func (m *FromFiles) PoliticianFromID(ID int) (Politician, error) {
	ps, err := m.AllPoliticians()
	if err != nil {
		return Politician{}, err
	}
	for _, p := range ps {
		if p.ID == ID {
			return p, nil
		}
	}
	return Politician{}, fmt.Errorf("politician of ID %d doesn't exist", ID)
}

var allVotes Votes

// AllVotes fetches all votes from JSON file if cache is empty, returns the cache otherwise
func (m *FromFiles) AllVotes() (Votes, error) {
	if allVotes == nil {
		allVotes = Votes{}
		for _, fileName := range m.VotesFileNames {
			var votesInFile Votes
			file, err := ioutil.ReadFile(path.Join(m.DirPath, fileName))
			if err != nil {
				return nil, err
			}
			if err = json.Unmarshal(file, &votesInFile); err != nil {
				return nil, err
			}
			allVotes = append(allVotes, votesInFile...)
		}
	}
	return allVotes, nil
}
