package arangoDb

import (
	"encoding/json"
	// "errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "strconv"
	// "strings"
)

//	Collection: A collection consists of documents. It is uniquely identified by 
//	it's collection identifier. It also has a unique name.
type collection struct {
	d               db
	_docCache       []document
	_collectionID   collectionID
	_collectionName string
	_name           string
	_status         string
	_properties     properties
}

type properties struct {
	all []byte
}

//	Collection Identifier: A collection identifier identifies a collection in a 
//	database. It is an integer and is unique within the database.
func (c *collection) ID() string {
	return c._collectionID
}

//	Collection Name: A collection name identifies a collection in a database. 
//	It is an string and is unique within the database. Unlike the collection identifier it is supplied by the creator of the collection. The collection name can consist of letters, digits and the characters _ (underscore) and - (dash). However, the first character must be a letter.
func (c *collection) Name() string {
	return c._name
}

//	The basic operations (create, read, update, delete) for documents are mapped to 
//	the standard HTTP methods (POST, GET, PUT, DELETE).

//	Creating and Deleting Collections

//	POST /_api/collection  
func (c *collection) Create(jsonBody string) error {
	c.d.u.Path = colBaseURL + "/"
	c.d.u.req.Method = "POST"
	err := d.u.get()
	return err
}

//	DELETE /_api/collection/collection-identifier
//	PUT /_api/collection/collection-identifier/truncate

var status = map[string]string{
	"1": "new born collection",
	"2": "unloaded",
	"3": "loaded",
	"4": "in the process of being unloaded",
	"5": "deleted",
}

//	Stats: Getting Information about a Collection
//	GET /_api/collection/collection-identifier
//	id: The identifier of the collection.
//	name: The name of the collection.
//	status: The status of the collection as number.
//	1: new born collection
//	2: unloaded
//	3: loaded
//	4: in the process of being unloaded
//	5: deleted
//	Every other status indicates a corrupted collection.
//	If the collection-identifier is unknown, then a HTTP 404 is returned.
func (c *collection) Stats() (string, error) {
	c.d.u.Path = colBaseURL + "/" + c._collectionID.ID()
	c.d.u.req.Method = "GET"
	var v interface{} // map[string]interface{}

	if x, err := d.u.do(); err != nil {
		y, err := ioutil.ReadAll(x.Body)
		json.Unmarshal(y, &v)
		fmt.Println(y, v)
		return "1", err
	}

	return "2", nil
}

//	GET /_api/collection/collection-identifier/properties
func (c *collection) Properties() (*http.Response, error) {
	c.d.u.Path = colBaseURL + "/" + c._collectionID.ID() + "/properties"
	c.d.u.req.Method = "GET"
	x, err := c.d.u.do()
	return x, err
}

//	GET /_api/collection/collection-identifier/count
func (c *collection) Count() (*http.Response, error) {
	c.d.u.Path = colBaseURL + "/" + c._collectionID.ID() + "/count"
	c.d.u.req.Method = "GET"
	x, err := d.u.do()
	return x, err
}

//	GET /_api/collection/collection-identifier/figures
func (c *collection) Figures() (*http.Response, error) {
	c.d.u.Path = colBaseURL + "/" + c._collectionID.ID() + "/figures"
	c.d.u.get()
	return x, err
}

//	GET /_api/collection/collection-identifier
func (c *collection) Get() (*http.Response, error) {
	c.d.u.Path = colBaseURL + "/" + c._collectionID.ID()
	c.d.u.req.Method = "GET"
	x, err := d.u.do()
	return x, err
}

//	GET /_api/collection/
func (c *collection) LoadCollections() (*http.Response, error) {
	c.d.u.Path = colBaseURL + "/"
	c.d.u.req.Method = "GET"
	x, err := d.u.do()
	return x, err
}

/*
HTTP/1.1 200 OK
content-type: application/json; charset=utf-8

{
  "documents": [
    "/document/6627082/29198126",
    "/document/6627082/29329198",
    "/document/6627082/29263662"
  ]
}
*/

//	Get /_api/document?collection=collection-identifier (Collection)
func (c *collection) getDocs() (*http.Response, error) {
	q := u.Query()
	q.Set("collection", c._collectionID.ID())
	c.d.u.RawQuery = q.Encode()
	c.d.u.Path = docBaseURL
	c.d.u.req.Method = "GET"
	x, err := d.u.do()
	return x, err

}

func (c *collection) GetDocs() {
	r, err := d.getDocs()
	e := json.Unmarshal(r.Body, &d.collections._docCache)
	return err
}

//	Modifying a Collection
//	PUT /_api/collection/collection-identifier/load
func (c *collection) load(jsonBody string) (*http.Response, error) {
	ul := colBaseURL + "/" + c._collectionID + "/load"
	return x, err
}

func (c *collection) Load(jsonBody string) (*http.Response, error) {
	ul := colBaseURL + "/" + c._collectionID + "/load"
	return x, err
}

//	PUT /_api/collection/collection-identifier/unload
func (c *collection) UnLoad(jsonBody string) (*http.Response, error) {
	ul := colBaseURL + "/" + c._collectionID + "/unload"
	return x, err
}

//	PUT /_api/collection/collection-identifier/properties
func (c *collection) UpdateProperties(jsonBody string) (*http.Response, error) {
	ul := colBaseURL + "/" + c._collectionID + "/properties"
	return x, err
}

//	PUT /_api/collection/collection-identifier/rename
func (c *collection) Rename(jsonBody string) (*http.Response, error) {
	ul := colBaseURL + "/" + c._collectionID + "/rename"
	return x, err
}
