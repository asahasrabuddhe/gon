package createdmg

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCmd(t *testing.T) {
	req := require.New(t)

	cmd, err := Cmd(context.Background())
	defer Close(cmd)

	req.NoError(err)
	req.FileExists(cmd.Path)
	req.FileExists(filepath.Join(cmd.Path, "..", "support", "dmg-license.py"))

	req.NoError(Close(cmd))
	req.NoError(Close(cmd))
}
