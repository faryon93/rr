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
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/spiral/roadrunner"
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"
	rrhttp "github.com/spiral/roadrunner/service/http"
)

// ---------------------------------------------------------------------------------------
//  constants
// ---------------------------------------------------------------------------------------

// ID contains default service name.
const ID = "workerlog"

// ---------------------------------------------------------------------------------------
//  types
// ---------------------------------------------------------------------------------------

// Service provides ability to forward stderr of the workers to stdout of roadrunner.
type Service struct {
	cfg    *Config
	logger *logrus.Logger
}

// ---------------------------------------------------------------------------------------
//  public members
// ---------------------------------------------------------------------------------------

// Init initializes the service.
func (s *Service) Init(cfg *Config, http *rrhttp.Service) (ok bool, err error) {
	if !cfg.Enable {
		return false, nil
	}

	s.logger = rr.Logger
	s.cfg = cfg

	http.AddListener(s.onEvent)

	return true, nil
}

// ---------------------------------------------------------------------------------------
//  private members
// ---------------------------------------------------------------------------------------

// onEvent handles events received by the service.
// Only StderrOutput events are processed by this service.
func (s *Service) onEvent(event int, ctx interface{}) {
	// we are only intrested in the stderr events
	if event != roadrunner.EventStderrOutput {
		return
	}

	// each line in the stderr buffer should be outputted individually
	for _, line := range strings.Split(string(ctx.([]byte)), "\n") {
		if line == "" {
			continue
		}

		if s.cfg.Decorate && s.logger != nil {
			s.logger.Warnln(line)
		} else {
			fmt.Println(line)
		}
	}
}
