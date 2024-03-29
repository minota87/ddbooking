package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CommandTestSuite struct {
	suite.Suite
	t *testing.T
}

func (c *CommandTestSuite) SetupSuite() {
	c.t = c.T()
}

func TestCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}

func (c *CommandTestSuite) TestShouldReturnValidCommandForImport() {
	command := createImportCMD()

	assert.NotNil(c.t, command)
	assert.Equal(c.t, command.Name, "import")
	assert.Equal(c.t, command.Usage, "Import Events")
	assert.NotNil(c.t, command.Action)
}

func (c *CommandTestSuite) TestShouldReturnValidCommandForValidate() {
	validate := createValidateCMD()

	assert.NotNil(c.t, validate)
	assert.Equal(c.t, validate.Name, "validate")
	assert.Equal(c.t, validate.Usage, "Validate Events")
	assert.NotNil(c.t, validate.Action)
}

func (c *CommandTestSuite) TestShouldReturnCommandsWhenCreateCommand() {
	commands := createCommands(createImportCMD, createValidateCMD)
	assert.Len(c.t, commands, 2)
}

func (c *CommandTestSuite) TestShouldReturnEmptyWhenCommandsListIsEmpty() {
	commands := createCommands()
	assert.Len(c.t, commands, 0)
}

func (f *FlagTestSuite) TestShouldReturnErrorWhenCommandImportWithNonExistentFlag() {
	c := CreateFakeContextWithFlag("teste", "test")

	err := commandImport(c)

	assert.Error(f.t, err)
}

func (f *FlagTestSuite) TestShouldNotReturnErrorWhenCommandImportWithValidFlag() {
	c := CreateFakeContextWithFlag("file", "../fixtures/events.json")

	err := commandImport(c)

	assert.NoError(f.t, err)
}

func (f *FlagTestSuite) TestShouldReturnErrorWhenCommandValidateWithNonExistentFlag() {
	c := CreateFakeContextWithFlag("teste", "test")

	err := commandValidate(c)

	assert.Error(f.t, err)
}

func (f *FlagTestSuite) TestShouldNotReturnErrorWhenCommandValidateWithValidFlag() {
	c := CreateFakeContextWithFlag("file", "../fixtures/events.json")

	err := commandValidate(c)

	assert.NoError(f.t, err)
}

func (f *FlagTestSuite) TestShouldReturnErrorWhenCommandValidateWithErrorInInputFile() {
	c := CreateFakeContextWithFlag("file", "../fixtures/events_with_error.json")

	err := commandValidate(c)

	assert.Error(f.t, err)
}

func (f *FlagTestSuite) TestShouldReturnErrorWhenCommandImportWithErrorInInputFile() {
	c := CreateFakeContextWithFlag("file", "../fixtures/events_with_error.json")

	err := commandImport(c)

	assert.Error(f.t, err)
}
