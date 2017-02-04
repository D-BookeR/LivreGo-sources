package elections

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"sync"

	"github.com/Pallinder/go-randomdata"
)

// ModelFiles holds the information and instruments the liaison of the model with flat JSON files
type ModelFiles struct {
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

func (p Politician) String() string {
	return p.Name + ", of \"" + p.Party + "\""
}

func (v Vote) String() string {
	return v.Name
}

var allPoliticians Politicians

// AllPoliticians fetches all politicians from JSON file if cache is empty, returns the cache otherwise
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

// PoliticianFromID finds the Politician value when given its ID
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

var allVotes Votes

func oneVoteFile(fileName string, m *ModelFiles, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	var allVotesFromThisFile Votes
	file, err := ioutil.ReadFile(path.Join(m.DirPath, fileName))
	if err != nil {
		fmt.Printf("File %s had an issue: ", fileName)
		fmt.Println(err)
		return
	}
	json.Unmarshal(file, &allVotesFromThisFile)
	mutex.Lock()
	allVotes = append(allVotes, allVotesFromThisFile...)
	mutex.Unlock()
}

// AllVotes fetches all votes from JSON file if cache is empty, returns the cache otherwise
func (m *ModelFiles) AllVotes() (Votes, error) {
	if allVotes == nil {
		var wg sync.WaitGroup
		var mutex = &sync.Mutex{}
		allVotes = Votes{}
		wg.Add(len(m.VotesFileNames))
		for _, fileName := range m.VotesFileNames {
			go oneVoteFile(fileName, m, &wg, mutex)
		}
		wg.Wait()
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
