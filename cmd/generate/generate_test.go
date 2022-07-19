package generate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestCommand(t *testing.T) {
	app := cli.NewApp()
	app.Commands = []*cli.Command{Command()}

	testCases := []struct {
		name      string
		commands  []string
		setup     func()
		assertion func(t *testing.T, err error)
		tearDown  func()
	}{
		{
			name:     "OK",
			commands: []string{"", "generate", "moduleName"},
			setup:    func() {},
			assertion: func(t *testing.T, err error) {
				require.NoError(t, err)
				assertModuleDir(t, "moduleName")
			},
			tearDown: func() {
				os.RemoveAll("moduleName")
			},
		},
		{
			name:     "No module neme input",
			commands: []string{"", "generate"},
			setup:    func() {},
			assertion: func(t *testing.T, err error) {
				require.Error(t, err)
			},
			tearDown: func() {},
		},
		{
			name:     "Module directory already exists",
			commands: []string{"", "generate", "existDir"},
			setup: func() {
				os.Mkdir("existDir", 0755)
			},
			assertion: func(t *testing.T, err error) {
				require.Error(t, err)
			},
			tearDown: func() {
				os.RemoveAll("existDir")
			},
		},
		{
			name:     "Module directory already exists with force flag",
			commands: []string{"", "generate", "-f", "existDirForce"},
			setup: func() {
				os.Mkdir("existDirForce", 0755)
			},
			assertion: func(t *testing.T, err error) {
				require.NoError(t, err)
				// MAYBE: check whether template files are updated.
			},
			tearDown: func() {
				os.RemoveAll("existDirForce")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			tc.setup()
			defer tc.tearDown()

			err := app.Run(tc.commands)
			tc.assertion(t, err)
		})
	}
}

func assertModuleDir(t *testing.T, moduleName string) {
	stat, err := os.Stat(moduleName)
	require.NoError(t, err)
	require.True(t, stat.IsDir())

	stat, err = os.Stat(filepath.Join(moduleName, "outputs.tf"))
	require.NoError(t, err)
	require.Zero(t, stat.Size())
	stat, err = os.Stat(filepath.Join(moduleName, "main.tf"))
	require.NoError(t, err)
	require.Zero(t, stat.Size())
	stat, err = os.Stat(filepath.Join(moduleName, "variables.tf"))
	require.NoError(t, err)
	require.Zero(t, stat.Size())

	stat, err = os.Stat(filepath.Join(moduleName, "README.md"))
	require.NoError(t, err)
	require.NotZero(t, stat.Size())
}
