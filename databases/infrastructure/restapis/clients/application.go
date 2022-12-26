package clients

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/configs"
	"github.com/steve-care-software/webx/databases/domain/connections"
	"github.com/steve-care-software/webx/databases/domain/contents/references"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type application struct {
	commitsAdapter              references.CommitsAdapter
	commitAdapter               references.CommitAdapter
	connectionsAdapter          connections.Adapter
	referenceContentKeysAdapter references.ContentKeysAdapter
	baseURL                     *url.URL
	client                      *resty.Client
}

func createApplication(
	commitsAdapter references.CommitsAdapter,
	commitAdapter references.CommitAdapter,
	connectionsAdapter connections.Adapter,
	referenceContentKeysAdapter references.ContentKeysAdapter,
	baseURL *url.URL,
	client *resty.Client,
) applications.Application {
	out := application{
		commitsAdapter:              commitsAdapter,
		commitAdapter:               commitAdapter,
		connectionsAdapter:          connectionsAdapter,
		referenceContentKeysAdapter: referenceContentKeysAdapter,
		baseURL:                     baseURL,
		client:                      client,
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

// Open opens a context on a given database
func (app *application) Open(name string) (*uint, error) {
	url := fmt.Sprintf(patternURI, app.baseURL.String(), contextURI)
	resp, err := app.client.R().Post(url)
	if err != nil {
		return nil, err
	}

	bytes := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(bytes))
	}

	if len(bytes) != 8 {
		str := fmt.Sprintf("the output was expected to contain 8 bytes, %d returned", len(bytes))
		return nil, errors.New(str)
	}

	context := uint(binary.LittleEndian.Uint64(bytes))
	return &context, nil
}

// ContentKeysByKind returns the contentKeys by context and kind
func (app *application) ContentKeysByKind(context uint, kind uint) (references.ContentKeys, error) {
	contentKeysURI := fmt.Sprintf(contentKeysByKindURI, context, kind)
	url := fmt.Sprintf(patternURI, app.baseURL.String(), contentKeysURI)
	resp, err := app.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	bytes := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(bytes))
	}

	return app.referenceContentKeysAdapter.ToContentKeys(bytes)
}

// CommitByHash returns the commit by hash
func (app *application) CommitByHash(context uint, hash hash.Hash) (references.Commit, error) {
	commitURI := fmt.Sprintf(commitByHashURI, context, hash.String())
	url := fmt.Sprintf(patternURI, app.baseURL.String(), commitURI)
	resp, err := app.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	bytes := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(bytes))
	}

	return app.commitAdapter.ToCommit(bytes)
}

// Commits returns the commits on a context
func (app *application) Commits(context uint) (references.Commits, error) {
	commitURI := fmt.Sprintf(commitsURI, context)
	url := fmt.Sprintf(patternURI, app.baseURL.String(), commitURI)
	resp, err := app.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	bytes := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(bytes))
	}

	return app.commitsAdapter.ToCommits(bytes)
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
	url := fmt.Sprintf(patternURI, app.baseURL.String(), contextURI)
	resp, err := app.client.R().Delete(url)
	if err != nil {
		return err
	}

	bytes := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return errors.New(string(bytes))
	}

	return nil
}
