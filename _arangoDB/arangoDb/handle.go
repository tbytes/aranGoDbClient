package arangoDb

//	Document Handle: A document handle uniquely identifies a document in the 
//	database. It is a string and consists of a collection identifier and a 
//	document identifier separated by /.
type handle struct {
	_name         string
	_collectionID string
	//	Document Identifier: A document identifier identifies a document in a 
	//	given collection. It is an integer and is unique within the collection of 
	//	the document.
	_docID string
}

func (h *handle) CollectionIdInt() (int64, error) {
	x, err := strconv.ParseInt(h._collectionID, 10, 64)
	return x, err
}

func (h *handle) DocIdInt() (int64, error) {
	x, err := strconv.ParseInt(h._docID, 10, 64)
	return x, err
}

func (h *handle) parseHandle(m map[string]interface{}) (err error) {

	if x := m["_id"]; x != nil {
		h._name = x.(string)
		s := strings.Split(h._name, "/")
		h._collectionID = s[0]
		h._docID = s[1]
		return err
	}
	return errors.New("No handle data")
}
