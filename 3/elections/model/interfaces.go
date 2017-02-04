package elections

// Reader contains all the functions to read the necessary data
type Reader interface {
	AllPoliticians() (Politicians, error)
	PoliticianFromID(ID int) (Politician, error)
	AllVotes() (Votes, error)
}
