package env

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFile(t *testing.T) {
	var err error
	err = os.Setenv("T_KEY_2", "VAL2")
	require.NoError(t, err)
	err = LoadFile(filepath.Join("testdata", "test.env"))
	require.NoError(t, err)
	require.Equal(t, "VAL1", os.Getenv("T_KEY_1"))
	require.Equal(t, "", os.Getenv("T_KEY_2"))
}
