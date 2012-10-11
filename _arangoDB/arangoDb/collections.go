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

const (
	colBaseURL = "/_api/collection"
)

//	Collection: A collection consists of documents. It is uniquely identified by 
//	it's collection identifier. It also has a unique name.
type collection struct {
	_docCache     []document
	_collectionID string
	_name         string
	_status       string
	client        Client
	_properties   properties
	_db           *db
}

func addCollection(jsonBody string) (col *collection, err error) {
	col = new(collection)
	htp, err := col.Create(jsonBody)
	log.Println(htp)
	return
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
func (c *collection) Create(jsonBody string) (*http.Response, error) {
	x, err := c.client.create(colBaseURL, jsonBody)
	fmt.Println("x: ", x, err)
	return x, err
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
	url := colBaseURL + "/" + c._collectionID
	var v map[string]interface{}
	if x, err := c.client.find(url); err != nil {
		log.Println(err)
		y, err := ioutil.ReadAll(x.Body)
		json.Unmarshal(y, &v)
		fmt.Println(v)
		return "1", err
	}
	return "2", nil
}

//	GET /_api/collection/collection-identifier/properties
func (c *collection) Properties() (*http.Response, error) {
	x, err := c.client.find(colBaseURL + "/" + c._collectionID + "/properties")
	return x, err
}

//	GET /_api/collection/collection-identifier/count
func (c *collection) Count() (*http.Response, error) {
	x, err := c.client.find(colBaseURL + "/" + c._collectionID + "/count")
	return x, err
}

//	GET /_api/collection/collection-identifier/figures
func (c *collection) Figures() (*http.Response, error) {
	x, err := c.client.find(colBaseURL + "/" + c._collectionID + "/figures")
	return x, err
}

//	GET /_api/collection/collection-identifier
func (c *collection) Get() (*http.Response, error) {
	x, err := c.client.find(colBaseURL + "/" + c._collectionID)
	return x, err
}

//	GET /_api/collection/
func (d *db) LoadCollections() (*http.Response, error) {
	x, err := c.client.find(colBaseURL + "/")
	return x, err
}

//	Modifying a Collection
//	PUT /_api/collection/collection-identifier/load
func (c *collection) Load(jsonBody string) (*http.Response, error) {
	x, err := c.client.update(colBaseURL+"/"+c._collectionID+"/load", jsonBody)
	return x, err
}

//	PUT /_api/collection/collection-identifier/unload
func (c *collection) UnLoad(jsonBody string) (*http.Response, error) {
	x, err := c.client.update(colBaseURL+"/"+c._collectionID+"/unload", jsonBody)
	return x, err
}

//	PUT /_api/collection/collection-identifier/properties
func (c *collection) UpdateProperties(jsonBody string) (*http.Response, error) {
	x, err := c.client.update(colBaseURL+"/"+c._collectionID+"/properties", jsonBody)
	return x, err
}

//	PUT /_api/collection/collection-identifier/rename
func (c *collection) Rename(jsonBody string) (*http.Response, error) {
	x, err := c.client.update(colBaseURL+"/"+c._collectionID+"/rename", jsonBody)
	return x, err
}
