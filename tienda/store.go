package tienda

type Store struct {
	Name        string
	Description string
	Contact     string
	Score       int
}

func (n *Store) SetName(Name string) {
	n.Name = Name
}

func (n *Store) SetDescription(Description string) {
	n.Description = Description
}

func (n *Store) SetContact(Contact string) {
	n.Contact = Contact
}

func (n *Store) SetScore(Score int) {
	n.Score = Score
}

func (n Store) GetName() string {
	return n.Name
}

func (n Store) GetDescription() string {
	return n.Description
}

func (n Store) GetContact() string {
	return n.Contact
}

func (n Store) GetScore() int {
	return n.Score
}
