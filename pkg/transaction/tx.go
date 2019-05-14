package transaction

// Tx holds the main Transaction struct
// that operations are performed on.
type Tx struct {
	Errors []error
	// TODO: Hook this up to queries
	// Query  []query.Query
}

// NewTransaction creates a new Tx and returns it. This also creates a
// snapshot of the collection at creation time and acquires a lock on the Collection.
func NewTransaction() (*Tx, error) {
	panic("not implemented")
}

// Commit enters a change into the database permanently
func (t *Tx) Commit() error {
	panic("not implemented")
}

// Rollback will return the database state to it's earlier state before work was done.
func (t *Tx) Rollback() error {
	panic("not implemented")
}

// Add adds a change to the transaction and returns an error if it was unsuccessful.
func (t *Tx) Add(coll string, update interface{}) error {
	panic("not implemented")
}
