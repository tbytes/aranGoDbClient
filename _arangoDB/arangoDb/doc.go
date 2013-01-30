package arangoDb

import (
	"net/url"
)

// If the document-handle points to a non-existing document, then a HTTP 404 is returned and 
//	the body contains an error document

// "If-None-Match" header is given, then it must contain exactly one etag. 
// If the document exists, then a HTTP 200 is returned and the JSON representation of the document is the body
//	The document is returned, if it has a different revision than the given etag. 
//	Otherwise a HTTP 304 is returned.
//	If the document is unchanged (the latest revision), a HTTP 412 is returned.

// If the "If-Match" header is given, then it must contain exactly one etag. 
//	The document is returned, if it has the same revision ad the given etag. 
//	Otherwise a HTTP 412 is returned. 

// As an alternative you can supply the etag in an attribute rev in the URL.

/*
{
  "_id" : "1234567/2345678",
  "_rev" : "3456789",
  "firstName" : "Hugo",
  "lastName" : "Schlonz",
  "address" : {
    "street" : "Strasse 1",
    "city" : "Hier"
  },
  "hobbies" : [
    "swimming",
    "biking",
    "programming"
  ]
}
All documents contain two special fields, the document handle in _id and the etag aka document revision in _rev.

Document Handle: A document handle uniquely identifies a document in the database. It is a string and consists 
	of a collection identifier and a document identifier separated by /.
Document Identifier: A document identifier identifies a document in a given collection. It unique within 
	the collection of the document.
Document Revision: As ArangoDB supports MVCC, documents can exist in more than one revision. 
	The document revision is the MVCC token used to identify a particular revision of a document. 
	It is an integer and unique within the list of document revision for a single document. 
	Earlier revisions of a document have smaller numbers. Document revisions can be used to 
	conditionally update, replace or delete documents in the database. In order to find a particular 
	revision of a document, you need the document handle and the document revision.ArangoDB currently uses 
	64bit unsigned integer values for document revisions. As this datatype is not portable to all client 
	languages, clients should rather use strings to store document revision ids locally.
Document Etag: The document revision enclosed in double quotes. The revision is returned by several 
	HTTP API methods in the Etag HTTP header.

The basic operations (create, read, update, delete) for documents are mapped to the standard 
	HTTP methods (POST, GET, PUT, DELETE). An identifier for the document revision is returned in 
	the "ETag" header field. If you modify a document, you can use the "If-Match" field to detect conflicts. 
	The revision of a document can be checking using the HTTP method HEAD.
*/
