package env

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestCandidateFiles(t *testing.T) {
	err := os.Setenv("PROFILE", "")
	require.NoError(t, err)
	c := CandidateFiles()
	require.Equal(t, []string{".env", filepath.Join("config", ".env"), filepath.Join("conf", ".env")}, c)
	err = os.Setenv("PROFILE", "GoTest  ")
	require.NoError(t, err)
	c = CandidateFiles()
	require.Equal(t, []string{".gotest.env", filepath.Join("config", ".gotest.env"), filepath.Join("conf", ".gotest.env"), ".env", filepath.Join("config", ".env"), filepath.Join("conf", ".env")}, c)
}
