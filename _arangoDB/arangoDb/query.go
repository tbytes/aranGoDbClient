package arangoDb

import (
	"encoding/json"
)

// Query
/*
{
  "hasMore": true,
  "error": false,
  "id": 26011191,
  "result": interface{},
  "code": 201,
  "count": 5
}
*/

/*
This is an introduction to ArangoDB's Http interface for Queries. 
Results of AQL and simple queries are returned as cursors in order to batch the 
communication between server and client. Each call returns a number of documents 
in a batch and an indication, if the current batch has been the final batch. 
Depending on the query, the total number of documents in the result set might or 
might not be known in advance. In order to free server resources the client should 
delete the cursor as soon as it is no longer needed.

HTTP Interface for AQL Query Cursors
Retrieving query results
Single roundtrip
Using a Cursor
Accessing Cursors via HTTP
POST /_api/cursor
POST /_api/query
PUT /_api/cursor/cursor-identifier
DELETE /_api/cursor/cursor-identifier
To run a select query, the query details need to be shipped from the client to the server via 
a HTTP POST request.

Retrieving query results

Single roundtrip
The server will only transfer a certain number of result documents back to the client in one roundtrip. 
This number is controllable by the client by setting the batchSize attribute when issueing the query.

If the complete result can be transferred to the client in one go, the client does not need to 
issue any further request. The client can check whether it has retrieved the complete result set 
by checking the hasMore attribute of the result set. If it is set to false, then the client has 
fetched the complete result set from the server.

Examples
> curl --data @- -X POST --dump - http://localhost:8529/_api/cursor
{ "query" : "FOR u IN users LIMIT 2 RETURN u", "count" : true, "batchSize" : 2 }

HTTP/1.1 201 Created
content-type: application/json

{
  "hasMore": false,
  "error": false,
  "result": [
    {
      "n": 0,
      "_rev": 21030455,
      "_id": "19588663/21030455"
    },
    {
      "n": 1,
      "_rev": 21030455,
      "_id": "19588663/21030455"
    }
  ],
  "code": 201,
  "count": 2
}
*/

/*
Using a Cursor

If the result set contains more documents than should be transferred in a single roundtrip 
(i.e. as set via the batchSize attribute), the server will return the first few documents 
and create a temporary cursor. The cursor identifier will also be returned to the client. 
The server will put the cursor identifier in the id attribute of the response object. 
Furthermore, the hasMore attribute of the response object will be set to true. This is an 
indication for the client that there are additional results to fetch from the server.

Examples
Create and extract first batch:

> curl --data @- -X POST --dump - http://localhost:8529/_api/cursor
{ "query" : "FOR u IN users LIMIT 5 RETURN u", "count" : true, "batchSize" : 2 }

HTTP/1.1 201 Created
content-type: application/json

{
  "hasMore": true,
  "error": false,
  "id": 26011191,
  "result": [
    {
      "n": 0,
      "_rev": 25880119,
      "_id": "23914039/25880119"
    },
    {
      "n": 1,
      "_rev": 25880119,
      "_id": "23914039/25880119"
    }
  ],
  "code": 201,
  "count": 5
}
Extract next batch, still have more:

> curl -X PUT --dump - http://localhost:8529/_api/cursor/26011191

HTTP/1.1 200 OK
content-type: application/json

{
  "hasMore": true,
  "error": false,
  "id": 26011191,
  "result": [
    {
      "n": 2,
      "_rev": 25880119,
      "_id": "23914039/25880119"
    },
    {
      "n": 3,
      "_rev": 25880119,
      "_id": "23914039/25880119"
    }
  ],
  "code": 200,
  "count": 5
}
Extract next batch, done:

> curl -X PUT --dump - http://localhost:8529/_api/cursor/26011191

HTTP/1.1 200 OK
content-type: application/json

{
  "hasMore": false,
  "error": false,
  "result": [
    {
      "n": 4,
      "_rev": 25880119,
      "_id": "23914039/25880119"
    }
  ],
  "code": 200,
  "count": 5
}
Do not do this:

> curl -X PUT --dump - http://localhost:8529/_api/cursor/26011191

HTTP/1.1 400 Bad Request
content-type: application/json

{
  "errorNum": 1600,
  "errorMessage": "cursor not found: disposed or unknown cursor",
  "error": true,
  "code": 400
}
*/

/*
Accessing Cursors via HTTP

POST /_api/cursor(creates a cursor)
POST /_api/cursor
The query details include the query string plus optional query options and bind parameters. These values 
need to be passed in a JSON representation in the body of the POST request.The following attributes 
can be used inside the JSON object:
	query: contains the query string to be executed (mandatory)
	count: boolean flag that indicates whether the number of documents found should be returned as "count" attribute in the result set (optional). Calculating the "count" attribute might have a performance penalty for some queries so this option is turned off by default.
	batchSize: maximum number of result documents to be transferred from the server to the client in one roundtrip (optional). If this attribute is not set, a server-controlled default value will be used.
	bindVars: key/value list of bind parameters (optional).

If the result set can be created by the server, the server will respond with HTTP 201. The body of the response 
will contain a JSON object with the result set.The JSON object has the following properties:
	error: boolean flag to indicate that an error occurred (false in this case)
	code: the HTTP status code
	result: an array of result documents (might be empty if query has no results)
	hasMore: a boolean indicator whether there are more results available on the server
	count: the total number of result documents available (only available if the query was executed with the count attribute set.
	id: id of temporary cursor created on the server (optional, see above)

If the JSON representation is malformed or the query specification is missing from the request, the server 
will respond with HTTP 400.The body of the response will contain a JSON object with additional error details. 
The object has the following attributes:
	error: boolean flag to indicate that an error occurred (true in this case)
	code: the HTTP status code
	errorNum: the server error number
	errorMessage: a descriptive error message

If the query specification is complete, the server will process the query. If an error occurs during query 
processing, the server will respond with HTTP 400. Again, the body of the response will contain details 
about the error.\

Examples
Executes a query and extract the result in a single go:
> curl --data @- -X POST --dump - http://localhost:8529/_api/cursor
{ "query" : "FOR u IN users LIMIT 2 RETURN u", "count" : true, "batchSize" : 2 }

HTTP/1.1 201 Created
content-type: application/json

{
  "hasMore": false,
  "error": false,
  "result": [
    {
      "n": 0,
      "_rev": 21030455,
      "_id": "19588663/21030455"
    },
    {
      "n": 1,
      "_rev": 21030455,
      "_id": "19588663/21030455"
    }
  ],
  "code": 201,
  "count": 2
}
Bad queries:
> curl -X POST --dump - http://localhost:8529/_api/cursor

HTTP/1.1 400 Bad Request
content-type: application/json

{
  "errorNum": 1503,
  "code": 400,
  "error": true,
  "errorMessage": "query specification invalid"
}
> curl --data @- -X POST --dump - http://localhost:8529/_api/cursor
{ "query" : "FOR u IN unknowncollection LIMIT 2 RETURN u.n", "count" : true, "bindVars" : {}, "batchSize" : 2 }

HTTP/1.1 400 Bad Request
content-type: application/json

{
  "code": 400,
  "error": true,
  "errorMessage": "unable to open collection '%s': unable to open collection 'unknowncollection'",
  "errorNum": 1510
}
*/

/*
POST /_api/query(parses a query)
POST /_api/query
To validate a query string without executing it, the query string can be passed to the server via an 
HTTP POST request.These query string needs to be passed in the attribute query of a JSON object as the 
body of the POST request.If the query is valid, the server will respond with HTTP 200 and return the 
names of the bind parameters it found in the query (if any) in the "bindVars" attribute of the response.
The server will respond with HTTP 400 in case of a malformed request, or if the query contains a parse error. 
The body of the response will contain the error details embedded in a JSON object.
Examples
Valid query:
> curl --data @- -X POST --dump - http://localhost:8529/_api/query
{ "query" : "FOR u IN users FILTER u.name == @name LIMIT 2 RETURN u.n" }

HTTP/1.1 200 OK
content-type: application/json

{
  "error": false,
  "bindVars": [
    "name"
  ],
  "code": 200
}
Invalid query:
> curl --data @- -X POST --dump - http://localhost:8529/_api/query
{ "query" : "FOR u IN users FILTER u.name = @name LIMIT 2 RETURN u.n" }

HTTP/1.1 400 Bad Request
content-type: application/json

{
  "errorNum": 1501,
  "errorMessage": "parse error: %s: parse error: 1:29 syntax error, unexpected assignment near ' = @name LIMIT 2 RETURN u.n'",
  "error": true,
  "code": 400
}
*/

/*
PUT /_api/cursor(reads next batch from a cursor)
PUT /_api/cursor/cursor-identifier
If the cursor is still alive, returns an object with the following attributes.
id: the cursor-identifier
result: a list of documents for the current batch
hasMore: false if this was the last batch
count: if present the total number of elements
Note that even if hasMore returns true, the next call might still return no documents. If, however, hasMore is false, then the cursor is exhausted. Once the hasMore attribute has a value of false, the client can stop.The server will respond with HTTP 200 in case of success. If the cursor identifier is ommitted or somehow invalid, the server will respond with HTTP 404.
Examples
Valid request for next batch:
> curl -X PUT --dump - http://localhost:8529/_api/cursor/26011191

HTTP/1.1 200 OK
content-type: application/json

{
  "hasMore": true,
  "error": false,
  "id": 26011191,
  "result": [
    {
      "n": 2,
      "_rev": 25880119,
      "_id": "23914039/25880119"
    },
    {
      "n": 3,
      "_rev": 25880119,
      "_id": "23914039/25880119"
    }
  ],
  "code": 200,
  "count": 5
}
Missing identifier
> curl -X PUT --dump - http://localhost:8529/_api/cursor

HTTP/1.1 400 Bad Request
content-type: application/json

{
  "code": 400,
  "errorMessage": "bad parameter",
  "errorNum": 400,
  "error": true
}

Unknown identifier
> curl -X PUT --dump - http://localhost:8529/_api/cursor/123456

HTTP/1.1 400 Bad Request
content-type: application/json

{
  "code": 400,
  "errorNum": 1600,
  "error": true,
  "errorMessage": "cursor not found: disposed or unknown cursor"
}
*/

/*
DELETE /_api/cursor(deletes a cursor)
DELETE /_api/cursor/cursor-identifier
Deletes the cursor and frees the resources associated with it.The cursor will automatically be destroyed on the server when the client has retrieved all documents from it. The client can also explicitly destroy the cursor at any earlier time using an HTTP DELETE request. The cursor id must be included as part of the URL.In case the server is aware of the cursor, it will respond with HTTP 202. Otherwise, it will respond with 404.Cursors that have been explicitly destroyed must not be used afterwards. If a cursor is used after it has been destroyed, the server will respond with HTTP 404 as well.Note: the server will also destroy abandoned cursors automatically after a certain server-controlled timeout to avoid resource leakage.
Examples
> curl -X DELETE --dump - http://localhost:8529/_api/cursor/8679702

HTTP/1.1 202 Accepted
content-type: application/json

{
  "code": 202,
  "id": "8679702",
  "error": false
}
*/

type query struct {
	hasMore bool
	err     bool
	id
}
