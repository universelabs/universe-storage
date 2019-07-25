package stormdb

import (
	// universe
	"github.com/universelabs/universe-server/universe"
	// deps
	"github.com/asdine/storm"
)

// Represents a client to the underlying stormDB instance.
type Client struct {
	db *storm.DB

	// Services
	keystore Keystore

	// Filename of the stormDB database
	Path string
}

func NewClient() *Client {
	c := &Client{}
	c.keystore.client = c
	return c
}

// Open and initialize stormDB
func (c *Client) Open(path string) error {
	var err error
	c.db, err = storm.Open(path)
	if err != nil {
		return err
	}
	c.Path = path
	err = c.db.Init(&universe.Wallet{})
	return err
}

func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

func (c *Client) Keystore() universe.Keystore { return &c.Keystore }
