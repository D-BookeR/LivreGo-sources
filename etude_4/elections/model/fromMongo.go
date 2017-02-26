package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// FromMongo holds the information and instruments the liaison of the model with MongoDB
type FromMongo struct {
	Server                string
	DbName                string
	PoliticiansCollection string
	VotesCollection       string
}

// AllPoliticians fetches all politicians from MongoDB
func (m *FromMongo) AllPoliticians() (Politicians, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.PoliticiansCollection)
	var result []Politician
	err = c.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// PoliticianFromID finds the Politician value when given its ID
func (m *FromMongo) PoliticianFromID(ID int) (Politician, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return Politician{}, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.PoliticiansCollection)
	var result Politician
	err = c.Find(bson.M{"id": ID}).One(&result)
	if err != nil {
		return Politician{}, err
	}

	return result, nil
}

// AllVotes fetches all votes from MongoDB
func (m *FromMongo) AllVotes() (Votes, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.VotesCollection)
	var result []Vote
	err = c.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
