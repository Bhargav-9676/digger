package digger

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUp() (string, func()) {
	tempDir := createTempDir()
	return tempDir, func() {
		deleteTempDir(tempDir)
	}
}

func TestDiggerConfigFileDoesNotExist(t *testing.T) {
	dg, err := NewDiggerConfig("")
	assert.NoError(t, err, "expected error to be not nil")
	assert.Equal(t, dg.Projects[0].Name, "default", "expected default project to have name 'default'")
	assert.Equal(t, dg.Projects[0].Dir, ".", "expected default project dir to be '.'")
}

func TestDiggerConfigWhenMultipleConfigExist(t *testing.T) {
	tempDir, teardown := setUp()
	defer teardown()

	_, err := os.Create(path.Join(tempDir, "digger.yaml"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Create(path.Join(tempDir, "digger.yml"))
	if err != nil {
		t.Fatal(err)
	}

	dg, err := NewDiggerConfig(tempDir)
	assert.Error(t, err, "expected error to be returned")
	assert.ErrorContains(t, err, ErrDiggerConfigConflict.Error(), "expected error to match target error")
	assert.Nil(t, dg, "expected diggerConfig to be nil")
}

func TestDiggerConfigWhenOnlyYamlExists(t *testing.T) {
	tempDir, teardown := setUp()
	defer teardown()

	_, err := os.Create(path.Join(tempDir, "digger.yaml"))
	if err != nil {
		t.Fatal(err)
	}

	dg, err := NewDiggerConfig(tempDir)
	assert.NoError(t, err, "expected error to be nil")
	assert.NotNil(t, dg, "expected digger config to be not nil")
}

func TestDiggerConfigWhenOnlyYmlExists(t *testing.T) {
	tempDir, teardown := setUp()
	defer teardown()

	_, err := os.Create(path.Join(tempDir, "digger.yml"))
	if err != nil {
		t.Fatal(err)
	}

	dg, err := NewDiggerConfig(tempDir)
	assert.NoError(t, err, "expected error to be nil")
	assert.NotNil(t, dg, "expected digger config to be not nil")
}

func TestDefaultValuesForWorkflowConfiguration(t *testing.T) {
	tempDir, teardown := setUp()
	defer teardown()

	f, err := os.Create(path.Join(tempDir, "digger.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	digger_yml := `
projects:
- name: dev
  branch: /main/
  dir: .
  workspace: default
`
	_, err2 := f.WriteString(digger_yml)
	if err2 != nil {
		log.Fatal(err2)
	}

	dg, err := NewDiggerConfig(tempDir)
	assert.NoError(t, err, "expected error to be nil")
	assert.Equal(t, dg.Projects[0].WorkflowConfiguration.OnPullRequestPushed[0], "digger plan")
	assert.Equal(t, dg.Projects[0].WorkflowConfiguration.OnPullRequestClosed[0], "digger unlock")
	assert.Equal(t, dg.Projects[0].WorkflowConfiguration.OnCommitToDefault[0], "digger apply")
}

func createTempDir() string {
	dir, err := os.MkdirTemp("", "tmp")
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func deleteTempDir(name string) {
	err := os.RemoveAll(name)
	if err != nil {
		log.Fatal(err)
	}
}