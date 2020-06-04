package env

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestStringVar(t *testing.T) {
	var val string
	var err error
	err = os.Setenv("T_KEY_1", "")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_2", "\n\t")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_3", "   Hello ")
	require.NoError(t, err)

	err = StringVar(&val, "T_KEY_0", "world")
	require.NoError(t, err)
	require.Equal(t, "world", val)
	err = StringVar(&val, "T_KEY_1", "world1")
	require.NoError(t, err)
	require.Equal(t, "world1", val)
	err = StringVar(&val, "T_KEY_2", "world2")
	require.NoError(t, err)
	require.Equal(t, "world2", val)
	err = StringVar(&val, "T_KEY_3", "world2")
	require.NoError(t, err)
	require.Equal(t, "Hello", val)
}

func TestInt64Var(t *testing.T) {
	var val int64
	var err error

	err = os.Setenv("T_KEY_1", "")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_2", "\n\t")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_3", "   11 ")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_4", "   ss")
	require.NoError(t, err)

	err = Int64Var(&val, "T_KEY_0", 11)
	require.NoError(t, err)
	require.Equal(t, int64(11), val)
	err = Int64Var(&val, "T_KEY_1", 11)
	require.NoError(t, err)
	require.Equal(t, int64(11), val)
	err = Int64Var(&val, "T_KEY_2", 11)
	require.NoError(t, err)
	require.Equal(t, int64(11), val)
	err = Int64Var(&val, "T_KEY_3", 12)
	require.NoError(t, err)
	require.Equal(t, int64(11), val)
	err = Int64Var(&val, "T_KEY_4", 12)
	require.Error(t, err)
}

func TestUint64Var(t *testing.T) {
	var val uint64
	var err error

	err = os.Setenv("T_KEY_1", "")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_2", "\n\t")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_3", "   11 ")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_4", "   ss")
	require.NoError(t, err)

	err = Uint64Var(&val, "T_KEY_0", 14)
	require.NoError(t, err)
	require.Equal(t, uint64(14), val)
	err = Uint64Var(&val, "T_KEY_1", 12)
	require.NoError(t, err)
	require.Equal(t, uint64(12), val)
	err = Uint64Var(&val, "T_KEY_2", 18)
	require.NoError(t, err)
	require.Equal(t, uint64(18), val)
	err = Uint64Var(&val, "T_KEY_3", 12)
	require.NoError(t, err)
	require.Equal(t, uint64(11), val)
	err = Uint64Var(&val, "T_KEY_4", 12)
	require.Error(t, err)
}

func TestBoolVar(t *testing.T) {
	var val bool
	var err error

	err = os.Setenv("T_KEY_1", "")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_2", "   t")
	require.NoError(t, err)
	err = os.Setenv("T_KEY_3", "   x")
	require.NoError(t, err)

	err = BoolVar(&val, "T_KEY_0", true)
	require.NoError(t, err)
	require.True(t, val)

	val = true
	err = BoolVar(&val, "T_KEY_1", false)
	require.NoError(t, err)
	require.False(t, val)

	val = false
	err = BoolVar(&val, "T_KEY_2", false)
	require.NoError(t, err)
	require.True(t, val)

	err = BoolVar(&val, "T_KEY_3", false)
	require.Error(t, err)
}

func TestFloat64Var(t *testing.T) {
	var val float64
	var err error

	err = os.Setenv("T_KEY_1", " 2.8")
	require.NoError(t, err)

	err = Float64Var(&val, "T_KEY_1", 1.8)
	require.NoError(t, err)
	require.Equal(t, float64(2.8), val)
}
