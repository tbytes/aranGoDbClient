package arangoDb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (d *docData) jsonString() (string, error) {
	x, err := json.Marshal(d.data.([]byte))
	return string(x), err
}

func newDocument() (doc *document) {
	return new(document)
}

// u.Path = "/_api/document"

//	HEAD /_api/document/document-handle

//	POST /_api/document?collection=collection-identifier
func (d *document) create(collectionID string, jsonBody string) (*http.Response, error) {
	d._url.Path = docBaseURL
	v := url.Values{}
	v.Set("collection", collectionID)
	d._url.RawQuery = v.Encode()

	x, err := d.client.create()
	fmt.Println("x: ", x, err)
	return x, err
}

//	GET /_api/document/document-handle
func (d *document) read(documentHandle string) (*http.Response, error) {
	x, err := d.client.find(docBaseURL + "/" + documentHandle)
	return x, err
}

//	PUT /_api/document/document-handle
func (d *document) update(documentHandle, jsonBody string) (*http.Response, error) {
	x, err := d.client.update(docBaseURL+"/"+documentHandle, jsonBody)
	return x, err
}

//	DELETE /_api/document/document-handle
func (d *document) delete(documentHandle string) (*http.Response, error) {
	x, err := d.client.delete(docBaseURL + "/" + documentHandle)
	return x, err
}

//	An identifier for the document revision is returned in the "ETag" header field
//	If you modify a document, you can use the "If-Match" field to detect 
//	conflicts. The revision of a document can be checking using the HTTP 
//	method HEAD.
