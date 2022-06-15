package gouuid

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNotFreeze(t *testing.T) {
	id1 := Gen()
	id2 := Gen()
	fmt.Println(id1, id2)
	require.NotEqual(t, "test", id1)
	require.NotEqual(t, "test", id2)
	require.Len(t, id1, 36)
	require.Len(t, id2, 36)
	require.NotEqual(t, id1, id2)
}

func TestFreeze(t *testing.T) {
	Freeze("test")
	require.Equal(t, "test", Gen())
}
