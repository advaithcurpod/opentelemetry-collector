// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configretry

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultBackOffSettings(t *testing.T) {
	cfg := NewDefaultBackOffConfig()
	assert.NoError(t, cfg.Validate())
	assert.Equal(t,
		BackOffConfig{
			Enabled:             true,
			InitialInterval:     5 * time.Second,
			RandomizationFactor: 0.5,
			Multiplier:          1.5,
			MaxInterval:         30 * time.Second,
			MaxElapsedTime:      5 * time.Minute,
		}, cfg)
}

func TestInvalidInitialInterval(t *testing.T) {
	cfg := NewDefaultBackOffConfig()
	assert.NoError(t, cfg.Validate())
	cfg.InitialInterval = -1
	assert.Error(t, cfg.Validate())
}

func TestInvalidRandomizationFactor(t *testing.T) {
	cfg := NewDefaultBackOffConfig()
	assert.NoError(t, cfg.Validate())
	cfg.RandomizationFactor = -1
	assert.Error(t, cfg.Validate())
	cfg.RandomizationFactor = 2
	assert.Error(t, cfg.Validate())
}

func TestInvalidMultiplier(t *testing.T) {
	cfg := NewDefaultBackOffConfig()
	assert.NoError(t, cfg.Validate())
	cfg.Multiplier = -1
	assert.Error(t, cfg.Validate())
}

func TestZeroMultiplierIsValid(t *testing.T) {
	cfg := NewDefaultBackOffConfig()
	assert.NoError(t, cfg.Validate())
	cfg.Multiplier = 0
	assert.NoError(t, cfg.Validate())
}

func TestInvalidMaxInterval(t *testing.T) {
	cfg := NewDefaultBackOffConfig()
	assert.NoError(t, cfg.Validate())
	cfg.MaxInterval = -1
	assert.Error(t, cfg.Validate())
}

func TestInvalidMaxElapsedTime(t *testing.T) {
	cfg := NewDefaultBackOffConfig()
	assert.NoError(t, cfg.Validate())
	cfg.MaxElapsedTime = -1
	assert.Error(t, cfg.Validate())
}

func TestDisabledWithInvalidValues(t *testing.T) {
	cfg := BackOffConfig{
		Enabled:             false,
		InitialInterval:     -1,
		RandomizationFactor: -1,
		Multiplier:          0,
		MaxInterval:         -1,
		MaxElapsedTime:      -1,
	}
	assert.NoError(t, cfg.Validate())
}
