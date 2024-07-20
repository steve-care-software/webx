package applications

import (
	"bytes"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

type testExecFunc func()

func TestBuilder_Success(t *testing.T) {
	rootPath := []string{
		"test_files",
	}

	path := []string{
		"dbfile.data",
	}

	defer func() {
		os.RemoveAll(filepath.Join(rootPath...))
	}()

	basePath := append(rootPath, "databases", "my_database")
	commitInnerPath := []string{"commits"}
	chunksInnerPath := []string{"chunks"}
	sizeToChunk := uint(1024)
	splitHashInThisAmount := uint(16)
	application, err := NewBuilder().Create().IsJSON().
		WithBasePath(basePath).
		WithChunksInnerPath(chunksInnerPath).
		WithCommitInnerPath(commitInnerPath).
		WithSizeInBytesToChunk(sizeToChunk).
		WithSplitChunkHashInThisAmountForDirectory(splitHashInThisAmount).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	executions := []testExecFunc{
		// begin invalid database, returns error
		func() {
			_, err := application.Begin(path)
			if err == nil {
				t.Errorf("the error was expected to be valid, nil returned")
				return
			}
		},
		// begin with init existing database, returns error
		func() {
			name := "My Name"
			description := "This is a description"
			pContext, err := application.BeginWithInit(path, name, description)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			expectedBytes := []byte("this is some data")
			err = application.Execute(*pContext, expectedBytes)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Commit(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Push(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			_, err = application.BeginWithInit(path, name, description)
			if err == nil {
				t.Errorf("the error was expected to be valid, nil returned")
				return
			}
		},
		// begin with init, cancel, begin with init, cancel works
		func() {
			name := "My Name"
			description := "This is a description"
			pContext, err := application.BeginWithInit(path, name, description)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			application.Cancel(*pContext)
			pContext, err = application.BeginWithInit(path, name, description)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			application.Cancel(*pContext)
		},
		// init, begin, commit once, push
		func() {
			name := "My Name"
			description := "This is a description"
			pContext, err := application.BeginWithInit(path, name, description)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			expectedBytes := []byte("this is some data")
			err = application.Execute(*pContext, expectedBytes)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Commit(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Push(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retDatabase, err := application.Retrieve(path)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			if retDatabase.MetaData().Name() != name {
				t.Errorf("the metaData name was expected to be '%s', '%s' returned", name, retDatabase.MetaData().Name())
				return
			}

			if retDatabase.MetaData().Description() != description {
				t.Errorf("the metaData description was expected to be '%s', '%s' returned", description, retDatabase.MetaData().Description())
				return
			}
		},
		// begin with init, execute, commit, push, begin, execute twice, commit, push
		func() {
			name := "My Name"
			description := "This is a description"
			pContext, err := application.BeginWithInit(path, name, description)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Execute(*pContext, []byte("this is some data"))
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Commit(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Push(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			pContext, err = application.Begin(path)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			expectedFirstBytes := []byte("this is the first bytes")
			expectedSecondBytes := []byte("this is the second bytes")
			err = application.Execute(*pContext, expectedFirstBytes)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Execute(*pContext, expectedSecondBytes)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Commit(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Push(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retDatabase, err := application.Retrieve(path)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retCommit, err := application.RetrieveCommit(retDatabase.Head().Hash())
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			list := retCommit.Executions().List()
			if len(list) != 2 {
				t.Errorf("the executions was expected to contain %d elements, %d returned", 2, len(list))
				return
			}
		},
		// begin, execute with chunk, commit, push
		func() {
			expectedChunk := []byte{}
			amount := int(sizeToChunk * 2)
			for i := 0; i < amount; i++ {
				expectedChunk = append(expectedChunk, byte(rand.Intn(255)))
			}

			name := "this is a name"
			description := "this is a descriptioon"
			pContext, err := application.BeginWithInit(path, name, description)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Execute(*pContext, expectedChunk)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Commit(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Push(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retDatabase, err := application.Retrieve(path)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retCommit, err := application.RetrieveCommit(retDatabase.Head().Hash())
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			list := retCommit.Executions().List()
			if len(list) != 1 {
				t.Errorf("the executions was expected to contain %d elements, %d returned", 1, len(list))
				return
			}

			if !list[0].IsChunk() {
				t.Errorf("the execution was expected to contain chunk")
				return
			}

			chunkBytes, err := application.RetrieveChunkBytes(list[0].Chunk().FingerPrint())
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			if !bytes.Equal(expectedChunk, chunkBytes) {
				t.Errorf("the returned chunks are invalid")
				return
			}
		},
		// begin, batch, commit, push
		func() {
			expectedChunk := []byte{}
			amount := int(sizeToChunk * 2)
			for i := 0; i < amount; i++ {
				expectedChunk = append(expectedChunk, byte(rand.Intn(255)))
			}

			secondBytes := []byte("this is some bytes")

			name := "this is a name"
			description := "this is a descriptioon"
			pContext, err := application.BeginWithInit(path, name, description)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Batch(*pContext, [][]byte{
				expectedChunk,
				secondBytes,
			})

			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Commit(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			err = application.Push(*pContext)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retDatabase, err := application.Retrieve(path)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retCommit, err := application.RetrieveCommit(retDatabase.Head().Hash())
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			list := retCommit.Executions().List()
			if len(list) != 2 {
				t.Errorf("the executions was expected to contain %d elements, %d returned", 2, len(list))
				return
			}
		},
	}

	for _, oneExecutionFn := range executions {
		oneExecutionFn()
		os.RemoveAll(filepath.Join(rootPath...))
	}

}
