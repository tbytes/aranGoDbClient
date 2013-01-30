package arangoDb

import (
	// "fmt"
	// "io"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

const (
	bufsize  = 1024
	host     = "localhost"
	port     = "8529"
	is_https = false
	prefix   = "http://"
	sPrefix  = "https://"

	colBaseURL    = "/_api/collection"
	docBaseURL    = "/_api/document"
	edgeBaseURL   = "/_api/edge"
	simpleBaseURL = "/_api/simple"
	sdminBaseURL  = "/admin"
	indexBaseURL  = "/_api/index"
)

type request http.Request

func (u *uri) do() error {
	u.req.URL = u.URL
	u.req.Header = u.Header
	resp, err := u.client.Do(u.req)
	return err
}

func (u *uri) post() error {
	u.req.Method = "POST"
	u.req.URL = u.URL
	u.req.Header = u.Header
	u.req.Body = u.Body
	resp, err := u.client.Do(u.req)
	return err
}

func (u *uri) get() error {
	u.req.Method = "GET"
	u.req.URL = u.URL
	u.req.Header = u.Header
	resp, err := u.client.Do(u.req)
	return err
}

func (u *uri) put() error {
	u.req.Method = "PUT"
	u.req.URL = u.URL
	u.req.Header = u.Header
	resp, err := u.client.Do(u.req)
	return err
}

func (u *uri) delete() error {
	u.req.Method = "DELETE"
	u.req.URL = u.URL
	u.req.Header = u.Header
	resp, err := u.client.Do(u.req)
	return err
}

type uri struct {
	url.URL
	is_https bool
	req      request
	Header   http.Header
	Body     io.Closer
	client   http.Client
}

type db struct {
	collections     map[string]collectionID
	prevCollection  collection
	c               collection // currentCollection
	nextCollection  collection
	collectionCache []collection
	u               uri
}

//	Document: Documents in ArangoDB are JSON objects. These objects can be 
//	nested (to any depth) and may contains lists. Each document is unique 
//	identified by its document handle.
type document struct {
	//	All documents contain two special fields, 
	_id     handle  //document handle
	_rev    eTag    //document revision enclosed in double quotes.
	_doc    docData //data struct
	_docMap map[string]interface{}
	_url    url.URL
}

type docData struct {
	data []byte
}

func (d *docData) marshal(v interface{}) error {
	d.data, err = json.Marshal(v)
	return err
}

func (d *docData) Get(v *interface{}) error {
	err := d.unmarshal(v)
	return err
}

func (d *docData) Set(v interface{}) error {
	err := d.marshal(v)
	return err
}

func (d *docData) unmarshal(v *interface{}) error {
	err := json.Unmarshal(d.data, v)
	return err
}

//	Document Revision: As AvocaodDB supports MVCC, documents can exist in more 
//	than one revision. The document revision is the MVCC token used to 
//	identify a particular revision of a document. It is an integer and unique 
//	within the list of document revision for a single document. Earlier 
//	revision of a document have smaller numbers. In order to find a particular 
//	revision of a document, you need the document handle and the document 
//	revision.
type eTag int

type collectionID struct {
	string
}

func DB() db {
	d := db{}
	d.init()
	return d
}

func (d *db) init() {
	d.u.is_https = is_https
	d.u.Scheme = prefix
	d.u.Host = d.host + ":" + d.port
	d.u.client = &http.Client{}
	return
}
