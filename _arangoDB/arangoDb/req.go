package arangoDb

/*
type collection struct {
	_docCache       []document
	_collectionID   collectionID
	_collectionName string
	_name           string
	_status         string
	_properties     properties
}

POST /_api/collection(creates a collection)
POST /_api/collection
Creates an new collection with a given name. The request must contain an object with the following attributes.
name (The name of the collection.)
waitForSync (optional, default: false): If true then the data is synchronised to disk before returning 
	from a create or update of an document.
journalSize (optional, default is a configuration parameter): The maximal size of a journal or datafile. 
	Note that this also limits the maximal size of a single object. Must be at least 1MB.
isSystem (optional, default is false): If true, create a system collection. In this case collection-name 
	should start with an underscore. End users should normally create non-system collections only. 
	API implementors may be required to create system collections in very special occasions, but normally 
	a regular collection will do.
type (optional, default is 2): the type of the collection to create. The following values for type are valid:
	2: document collection
	3: edges collection


Response:

HTTP/1.1 200 OK
content-type: application/json
location: /_api/collection/179665369
*/

type collectionRR struct {
	Name        string
	Code        int
	WaitForSync bool
	Id          int
	Status      int
	Type        int
	Error       bool
}

func (c *collection) NewCollection() error {
	c.d.u.Path = colBaseURL
	rr := collectionRR{}

	err := d.u.post(rr)
	return err
}

func (c *collectionRR) json() ([]byte, error) {
	b, err := json.Marshal(c)
	return b, err
}
