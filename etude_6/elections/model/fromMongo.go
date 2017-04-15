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

// Winner return the winner directly, by performing a MongoDB aggregation
func (m *FromMongo) Winner() (Politician, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return Politician{}, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.VotesCollection)
	resp := bson.M{}

	// db.votes.aggregate([{$group:{_id: "$politician_id", count: {$sum: 1}}},{$sort: {count: -1}},{$limit:1}])
	pipe := c.Pipe([]bson.M{{"$group": bson.M{"_id": "$politician_id", "count": bson.M{"$sum": 1}}}, bson.M{"$sort": bson.M{"count": -1}}, bson.M{"$limit": 1}})
	err = pipe.One(&resp)
	if err != nil {
		return Politician{}, err
	}

	// Now that we have the ID, we need to perform another query to find the actual politician
	p, err := m.PoliticianFromID(resp["_id"].(int))
	if err != nil {
		return Politician{}, err
	}

	return p, nil
}

// SaveVote takes a vote and saves it in the database
func (m *FromMongo) SaveVote(v Vote) error {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.VotesCollection)
	err = c.Insert(v)
	if err != nil {
		return err
	}

	return nil
}
