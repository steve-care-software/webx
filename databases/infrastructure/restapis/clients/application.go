package clients

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/commits"
	"github.com/steve-care-software/webx/databases/domain/commits/histories"
	"github.com/steve-care-software/webx/databases/domain/configs"
	"github.com/steve-care-software/webx/databases/domain/connections"
	"github.com/steve-care-software/webx/databases/domain/contents/references"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type application struct {
	connectionsAdapter connections.Adapter
	baseURL            *url.URL
	client             *resty.Client
}

func createApplication(
	connectionsAdapter connections.Adapter,
	baseURL *url.URL,
	client *resty.Client,
) applications.Application {
	out := application{
		connectionsAdapter: connectionsAdapter,
		baseURL:            baseURL,
		client:             client,
	}

	return &out
}

// Exists returns true if the database exists, false otherwise
func (app *application) Exists(name string) (bool, error) {
	url := fmt.Sprintf(patternURI, app.baseURL.String(), existsURI)
	resp, err := app.client.R().Get(url)
	if err != nil {
		return false, err
	}

	bytes := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return false, errors.New(string(bytes))
	}

	if len(bytes) != 1 {
		str := fmt.Sprintf("the output was expected to contain 1 byte, %d returned", len(bytes))
		return false, errors.New(str)
	}

	// return false if 0, anything else is true:
	return bytes[0] != 0, nil
}

// New creates a new database
func (app *application) New(name string) error {
	url := fmt.Sprintf(patternURI, app.baseURL.String(), rootURI)
	resp, err := app.client.R().Post(url)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		bytes := resp.Body()
		return errors.New(string(bytes))
	}

	return nil
}

// Delete deletes an existing database
func (app *application) Delete(name string) error {
	url := fmt.Sprintf(patternURI, app.baseURL.String(), rootURI)
	resp, err := app.client.R().Delete(url)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		bytes := resp.Body()
		return errors.New(string(bytes))
	}

	return nil
}

// Connections returns the active connections
func (app *application) Connections() (connections.Connections, error) {
	url := fmt.Sprintf(patternURI, app.baseURL.String(), connectionsURI)
	resp, err := app.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	bytes := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(bytes))
	}

	return app.connectionsAdapter.ToConnections(bytes)
}

// Open opens a context at height, height is -1 if the head is requested
func (app *application) Open(name string, height int) (*uint, error) {
	return nil, nil
}

// ContentKeysByKind returns the contentKeys by context and kind
func (app *application) ContentKeysByKind(context uint, kind uint) (references.ContentKeys, error) {
	return nil, nil
}

func (app *application) contentKeys(context uint) (references.ContentKeys, error) {
	return nil, nil
}

// CommitByHash returns the commit by hash
func (app *application) CommitByHash(context uint, hash hash.Hash) (commits.Commit, error) {
	return nil, nil
}

// Histories returns the commits histories on a context
func (app *application) Histories(context uint) (histories.Histories, error) {
	return nil, nil
}

// Read reads a pointer on a context
func (app *application) Read(context uint, pointer references.Pointer) ([]byte, error) {
	return nil, nil
}

// ReadByHash reads content by hash
func (app *application) ReadByHash(context uint, hash hash.Hash) ([]byte, error) {
	return nil, nil
}

// ReadAll read pointers on a context
func (app *application) ReadAll(context uint, pointers []references.Pointer) ([][]byte, error) {
	return nil, nil
}

// ReadAllByHashes reads content by hashes
func (app *application) ReadAllByHashes(context uint, hashes []hash.Hash) ([][]byte, error) {
	return nil, nil
}

// Write writes data to a context
func (app *application) Write(context uint, hash hash.Hash, data []byte, kind uint) error {
	return nil
}

// Cancel cancels a context
func (app *application) Cancel(context uint) error {
	return nil
}

// Commit commits a context
func (app *application) Commit(context uint) error {
	return nil
}

// Share shares the database interactions with a new peer, using the given context
func (app *application) Share(context uint, peer *url.URL) error {
	return nil
}

// Push retrieves the commits from peers, then chain our commits to them using the given configuration
func (app *application) Push(name string, config configs.Config) error {
	return nil
}

// Close closes a context
func (app *application) Close(context uint) error {
	return nil
}
