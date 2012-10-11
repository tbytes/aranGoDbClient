package arangoDb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	docBaseURL = "/_api/document"
)

//	Document: Documents in ArangoDB are JSON objects. These objects can be 
//	nested (to any depth) and may contains lists. Each document is unique 
//	identified by its document handle.
type document struct {
	//	All documents contain two special fields, 
	_id             handle //document handle
	_rev            eTag   //document revision enclosed in double quotes.
	_collectionName string
	_doc            docData //data struct
	_docMap         map[string]interface{}
	client          Client
}

type docData struct {
	data interface{}
}

func (d *docData) jsonString() (string, error) {
	x, err := json.Marshal(d.data.([]byte))
	return string(x), err
}

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

func newDocument() (doc *document) {
	return new(document)
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

//	Document Revision: As AvocaodDB supports MVCC, documents can exist in more 
//	than one revision. The document revision is the MVCC token used to 
//	identify a particular revision of a document. It is an integer and unique 
//	within the list of document revision for a single document. Earlier 
//	revision of a document have smaller numbers. In order to find a particular 
//	revision of a document, you need the document handle and the document 
//	revision.
type eTag int

//	HEAD /_api/document/document-handle

//	POST /_api/document?collection=collection-identifier
func (d *document) Create(collectionID string, jsonBody string) (*http.Response, error) {
	x, err := d.client.create(docBaseURL+"?collection=", jsonBody)
	fmt.Println("x: ", x, err)
	return x, err
}

//	GET /_api/document/document-handle
func (d *document) Read(documentHandle string) (*http.Response, error) {
	x, err := d.client.find(docBaseURL + "/" + documentHandle)
	return x, err
}

//	Get /_api/document?collection=collection-identifier (Collection)
func (d *document) GetDocs(collectionID string) (*http.Response, error) {
	url := docBaseURL + "?collection=" + collectionID
	x, err := d.client.find(url)
	return x, err
}

//	PUT /_api/document/document-handle
func (d *document) Update(documentHandle, jsonBody string) (*http.Response, error) {
	x, err := d.client.update(docBaseURL+"/"+documentHandle, jsonBody)
	return x, err
}

//	DELETE /_api/document/document-handle
func (d *document) Delete(documentHandle string) (*http.Response, error) {
	x, err := d.client.delete(docBaseURL + "/" + documentHandle)
	return x, err
}

//	An identifier for the document revision is returned in the "ETag" header field
//	If you modify a document, you can use the "If-Match" field to detect 
//	conflicts. The revision of a document can be checking using the HTTP 
//	method HEAD.
