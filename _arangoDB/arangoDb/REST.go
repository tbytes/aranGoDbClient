package arangoDb

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type client struct {
	conn     *httputil.ClientConn
	resource *url.URL
}

// Creates a client for the specified resource.
//
// The resource is the url to the base of the resource (i.e.,
// http://127.0.0.1:3000/snips/)
func newClient(resource string) (*Client, error) {
	var client = new(client)
	var err error

	// setup host
	if client.resource, err = url.Parse(resource); err != nil {
		return nil, err
	}

	// Setup conn
	var tcpConn net.Conn
	if tcpConn, err = net.Dial("tcp", client.resource.Host); err != nil {
		return nil, err
	}
	client.conn = httputil.NewClientConn(tcpConn, nil)

	return client, nil
}

// Closes the clients connection
func (client *client) Close() {
	client.conn.Close()
}

// General Request method used by the specialized request methods to create a request
func (c *client) newRequest(method string, id string) (*http.Request, error) {
	request := new(http.Request)
	var err error
	fmt.Println(c)
	request.ProtoMajor = 1
	request.ProtoMinor = 1
	request.TransferEncoding = []string{"chunked"}
	request.Method = method

	// Generate Resource-URI and parse it
	uri := c.resource.String() + id
	fmt.Println(uri)
	if request.URL, err = url.Parse(uri); err != nil {
		return nil, err
	}

	return request, nil
}

// Send a request
func (c *client) request(request *http.Request) (*http.Response, error) {
	var err error
	var response *http.Response

	// Send the request
	if err = c.conn.Write(request); err != nil {
		return nil, err
	}

	// Read the response
	if response, err = c.conn.Read(request); err != nil {
		return nil, err
	}

	return response, nil
}

// GET /resource/
func (c *client) index() (*http.Response, error) {
	var request *http.Request
	var err error

	if request, err = c.newRequest("GET", ""); err != nil {
		return nil, err
	}
	return c.request(request)
}

// GET /resource/id
func (c *client) find(id string) (*http.Response, error) {
	var request *http.Request
	var err error

	if request, err = c.newRequest("GET", id); err != nil {
		return nil, err
	}

	return c.request(request)
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

// POST /resource
func (c *client) create(urlId, body string) (*http.Response, error) {
	var request *http.Request
	var err error

	if request, err = c.newRequest("POST", urlId); err != nil {
		return nil, err
	}
	request.Body = nopCloser{bytes.NewBufferString(body)}
	return c.request(request)
}

// PUT /resource/id
func (c *client) update(id string, body string) (*http.Response, error) {
	var request *http.Request
	var err error
	if request, err = c.newRequest("PUT", id); err != nil {
		return nil, err
	}

	request.Body = nopCloser{bytes.NewBufferString(body)}

	return client.request(request)
}

// Parse a response-Location-URI to get the ID of the worked-on snip
func (c *client) IdFromURL(urlString string) (string, error) {
	var uri *url.URL
	var err error
	if uri, err = url.Parse(urlString); err != nil {
		return "", err
	}

	return string(uri.Path[len(client.resource.Path):]), nil
}

// DELETE /resource/id
func (c *client) delete(id string) (*http.Response, error) {
	var request *http.Request
	var err error
	if request, err = c.newRequest("DELETE", id); err != nil {
		return nil, err
	}

	return c.request(request)
}
