package arangoDb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
This is an introduction to ArangoDB's REST interface for edges.

AvocacadoDB offers also some graph functionality. A graph consists of nodes, edges and properties. 
ArangoDB stores the information how the nodes relate to each other aside from the properties.

So a graph data model always consists of two collections: the relations between the nodes in the graphs 
are stored in an "edges collection", the nodes in the graph are stored in documents in regular collections.

Example:

the edge collection stores the information that a company's reception is sub-unit to the services unit and 
the services unit is sub-unit to the CEO. You would express this relationship with the _from and _to property
the "normal" collection stores all the properties about the reception, 
	e.g. that 20 people are working there 
	and the room number etc.
_from is the document handle of the linked vertex (incoming relation), 
_to is the document handle of the 
linked vertex (outgoing relation).

Documents, Identifiers, Handles
Edge: Edges in ArangoDB are special documents. In addition to the internal attributes _id and _rev, they have two 
attributes _form and _to, which contain document handles namely the start-point and the end-point of the edge.

Address and ETag of an Edge
All documents in ArangoDB have a document handle. This handle uniquely defines a document and is managed by 
ArangoDB. All documents are found under the URI
http://server:port/_api/document/document-handle

For edges you can use the special address
http://server:port/_api/edge/document-handle

For example: Assume that the document handle, which is stored in the _id field of the edge, is 7254820/362549736, 
then the URL of that edge is:
http://localhost:8529/_api/edge/7254820/362549736
*/
type edge struct {
	//	All documents contain four special fields, 
	_id   handle // document handle
	_rev  eTag   // document revision enclosed in double quotes.
	_form handle // the start-point of the edge.
	_to   handle // the end-point of the edge.

	_nodes
	_edges
	_properties properties
	_db         *db
}

// POST /_api/edge?collection=collection-identifier&from=from-handle&to=to-handle
//	Creates a new edge in the collection identified by the collection-identifier. 
//	A JSON representation of the document must be passed as the body of the POST request. 
//	The object handle of the start point must be passed in from-handle. The object handle 
//	of the end point must be passed in to-handle.
//	Example:	> curl --data @- -X POST --dump - http://localhost:8529/_api/edge?collection=7848004&from=7848004/9289796&to=7848004/9355332
//				{ "e" : 1 }
func (c *collection) Create(jsonBody string) (*http.Response, error) {
	x, err := c.client.create(colBaseURL, jsonBody)
	fmt.Println("x: ", x, err)
	return x, err
}

// GET /_api/edge (reads an edge)
// GET /_api/edge/document-handle
//	Example:	> curl -X GET --dump - http://localhost:8529/_api/edge/7848004/9683012

// POST /_api/edge?collection=collection-identifier&from=from-handle&to=to-handle
// PUT /_api/edge/document-handle
// DELETE /_api/edge/document-handle
// HEAD /_api/edge/document-handle

// GET /_api/edges/collection-identifier?vertex=vertex-handle&directory=direction

// GET /_api/edges/collection-identifier?vertex=vertex-handle&direction=any
//	Returns the list of edges starting or ending in the vertex identified by vertex-handle.
//	Example:	http://localhost:8529/_api/edges/17501660?vertex=17501660/18419164

// GET /_api/edges/collection-identifier?vertex=vertex-handle&direction=in
//	Returns the list of edges ending in the vertex identified by vertex-handle.
//	Example:	http://localhost:8529/_api/edges/17501660?vertex=17501660/18419164&direction=in

// GET /_api/edges/collection-identifier?vertex=vertex-handle&direction=out
//	Returns the list of edges starting in the vertex identified by vertex-handle.
//	Example:	http://localhost:8529/_api/edges/17501660?vertex=17501660/18419164&direction=out
