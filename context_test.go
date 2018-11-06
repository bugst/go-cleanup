//
// Copyright 2018 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package cleanup_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.bug.st/cleanup"
)

func TestInterruptableContext(t *testing.T) {
	ctx, cancel := cleanup.InterruptableContext(context.Background())
	defer cancel()

	err := ctxDelay(ctx)
	require.NoError(t, err)
}

func TestInterruptableContextWithInterruption(t *testing.T) {
	ctx, cancel := cleanup.InterruptableContext(context.Background())
	defer cancel()

	go func() {
		time.Sleep(100 * time.Millisecond)
		// Simulate CTRL+C
		pid := os.Getpid()
		p, err := os.FindProcess(pid)
		require.NoError(t, err)
		err = p.Signal(os.Interrupt)
		require.NoError(t, err)
	}()

	err := ctxDelay(ctx)
	require.Error(t, err)
	require.Equal(t, "interrupted", err.Error())
}

func ctxDelay(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return errors.New("interrupted")
	case <-time.After(time.Second):
		return nil
	}
}
