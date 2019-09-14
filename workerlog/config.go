package workerlog

// rr
// Copyright (C) 2019 Maximilian Pachl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// ---------------------------------------------------------------------------------------
//  imports
// ---------------------------------------------------------------------------------------

import (
	"github.com/spiral/roadrunner/service"
)

// ---------------------------------------------------------------------------------------
//  types
// ---------------------------------------------------------------------------------------

// Config
type Config struct {
	Enable   bool
	Decorate bool
}

// ---------------------------------------------------------------------------------------
//  public members
// ---------------------------------------------------------------------------------------

// Hydrate configuration.
func (c *Config) Hydrate(cfg service.Config) error {
	return cfg.Unmarshal(c)
}

// InitDefaults allows to init blank config with pre-defined set of default values.
func (c *Config) InitDefaults() error {
	c.Enable = true
	return nil
}
