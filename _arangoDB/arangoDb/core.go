package arangoDb

import (
	// "fmt"
	// "io"
	// "net/http"
	// "encoding/json"
	"log"
)

const (
	bufsize  = 1024
	host     = "localhost"
	port     = "8529"
	is_https = false
	prefix   = "http://"
	sPrefix  = "https://"
)

// Client itself.
type db struct {
	host        string
	port        string
	is_https    bool
	prefix      string
	addr        string
	collections map[string]*collection
	client      *Client
}

func DB() db {
	d := db{}
	return d
}

func (d *db) Default() {
	d.host = host
	d.port = port
	d.prefix = prefix
	d.is_https = is_https
	d.addr = d.prefix + d.host + ":" + d.port
	d.client, _ = NewClient(d.addr)
	return
}
func (d *db) GetCollection(jsonBody string) (col *collection) {
	col = new(collection)
	c := d.client
	col.client = *c
	col._db = d
	d.collections[jsonBody] = col
	return
}

func (d *db) NewCollection(jsonBody string) {
	col, err := addCollection(jsonBody)
	log.Println(err)
	c := d.client
	col.client = *c
	//d.collections = append(d.collections, col)
	d.collections[jsonBody] = col
	return
}
