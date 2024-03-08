package shell

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommand(t *testing.T) {
	err := Run("ls / | grep usr", WithLogFunc(log.Printf))
	require.NoError(t, err)
}
