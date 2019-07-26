package stormdb_test

import (
	// stdlib
	"io/ioutil"
	"os"
	// universe
	"github.com/universelabs/universe-server/stormdb"
)

// test wrapper for stormdb.Client
type Client struct {
	*stormdb.Client
}

// NewClient returns an instance of Client pointing at a temporary file
func NewClient() *Client {
	// generate temporary file
	f, err := ioutil.TempFile("", "universe-stormdb-client-")
	if err != nil {
		panic(err)
	}
	f.Close()

	// create client wrapper
	c := &Client{ Client: stormdb.NewClient(), }
	c.Path = f.Name()

	return c
}

// MustOpenClient returns a new, open instance of Client
func MustOpenClient() *Client {
	c := NewClient()
	if err := c.Open(c.Path); err != nil {
		panic(err)
	}
	return c
}

// Clode closes the client and removes the underlying database
func (c *Client) Close() error {
	defer os.Remove(c.Path)
	return c.Client.Close()
}
