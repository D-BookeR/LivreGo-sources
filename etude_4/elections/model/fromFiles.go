package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"sync"
)

// FromFiles holds the information and instruments the liaison of the model with flat JSON files
type FromFiles struct {
	DirPath             string   // le répertoire dans lequel tous le fichiers sont
	PoliticiansFileName string   // le seul fichier dans lequel on trouve les politiciens
	VotesFileNames      []string // tous les fichiers contenant des votes
}

var allPoliticians Politicians

// AllPoliticians fetches all politicians from JSON file if cache is empty, returns the cache otherwise
func (m *FromFiles) AllPoliticians() (Politicians, error) {
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
	return Politician{}, fmt.Errorf("Politician of ID %d doesn't exist", ID)
}

var allVotes Votes

func oneVoteFile(fileName string, m *FromFiles, wg *sync.WaitGroup, mutex *sync.Mutex) {
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
func (m *FromFiles) AllVotes() (Votes, error) {
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
