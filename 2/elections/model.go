package elections

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/Pallinder/go-randomdata"
)

// ModelFiles is a type
type ModelFiles struct {
	DirPath             string
	PoliticiansFileName string
	VotesFileNames      []string
}

// Politician is a type
type Politician struct {
	Name  string `json:"name"`
	ID    int    `json:"id,omitempty"`
	Party string `json:"party,omitempty"`
}

// Politicians is a type
type Politicians []Politician

// Vote is a type
type Vote struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	PoliticianID int    `json:"politician_id"`
}

// Votes is a type
type Votes []Vote

func (p Politician) String() string {
	return p.Name + ", of \"" + p.Party + "\""
}

func (v Vote) String() string {
	return v.Name
}

var allPoliticians Politicians
var allVotes Votes

// AllPoliticians does stuff with caching
func (m *ModelFiles) AllPoliticians() (Politicians, error) {
	if allPoliticians == nil {
		file, err := ioutil.ReadFile(path.Join(m.DirPath, m.PoliticiansFileName))
		if err != nil {
			return nil, err
		}
		json.Unmarshal(file, &allPoliticians)
	}
	return allPoliticians, nil
}

// PoliticianFromID does stuff with caching
func (m *ModelFiles) PoliticianFromID(ID int) (Politician, error) {
	ps, err := m.AllPoliticians()
	if err != nil {
		return Politician{}, err
	}
	for _, p := range ps {
		if p.ID == ID {
			return p, nil
		}
	}
	return Politician{}, fmt.Errorf("Politician of ID %d doesn't exist", ID)
}

// AllVotes does stuff with caching
func (m *ModelFiles) AllVotes() (Votes, error) {
	if allVotes == nil {
		allVotes = Votes{}
		for _, fileName := range m.VotesFileNames {
			var allVotesFromThisFile Votes
			file, err := ioutil.ReadFile(path.Join(m.DirPath, fileName))
			if err != nil {
				return nil, err
			}
			json.Unmarshal(file, &allVotesFromThisFile)
			allVotes = append(allVotes, allVotesFromThisFile...)
		}
	}
	return allVotes, nil
}

// NOT FOR THE BOOK

// GenerateAndWriteVotes does stuff
func (m *ModelFiles) GenerateAndWriteVotes(nbFiles int, nbVotesPerFiles int, nbPoliticians int) error {
	for iFile := 1; iFile <= nbFiles; iFile++ {
		fileName := path.Join(m.DirPath, fmt.Sprintf("votes_%d.json", iFile))
		var votes Votes
		for iVote := 0; iVote < nbVotesPerFiles; iVote++ {
			votes = append(votes, Vote{Name: randomdata.FullName(randomdata.RandomGender), ID: fmt.Sprintf("%d_%d", iFile, iVote), PoliticianID: randomdata.Number(3)})
		}
		votesJSON, err := json.Marshal(votes)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(fileName, votesJSON, 0640)
		if err != nil {
			return err
		}
	}
	return nil
}
