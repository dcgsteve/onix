/*
  Onix Config Manager - Artisan
  Copyright (c) 2018-2021 by www.gatblau.org
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/
package flow

import (
	"fmt"
	"github.com/gatblau/onix/artisan/core"
	"github.com/gatblau/onix/artisan/crypto"
	"github.com/gatblau/onix/artisan/data"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// the pipeline generator requires at least the flow definition
// if a build file is passed then step variables can be inferred from it
type Manager struct {
	flow      *Flow
	buildFile *data.BuildFile
	pub       *crypto.PGP
}

func NewFromPath(flowPath, pubKeyPath, buildPath string) (*Manager, error) {
	m := new(Manager)
	flow, err := loadFlow(flowPath)
	if err != nil {
		return nil, fmt.Errorf("cannot load flow definition from %s: %s", flowPath, err)
	}
	m.flow = flow
	// if a build file is defined, then load it
	if len(buildPath) > 0 {
		buildFile, err := data.LoadBuildFile(buildPath)
		if err != nil {
			return nil, fmt.Errorf("cannot load build file from %s: %s", buildPath, err)
		}
		m.buildFile = buildFile
	}
	m.pub, err = crypto.LoadPGP(pubKeyPath)
	core.CheckErr(err, "cannot load public PGP encryption key")
	if m.pub.HasPrivate() {
		return nil, fmt.Errorf("a private PGP key has been provided but a public PGP key is required")
	}
	err = m.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid generator: %s", err)
	}
	return m, nil
}

func NewFromRemote(remotePath string) *Manager {
	return &Manager{}
}

func (m *Manager) FillIn() {
	for _, step := range m.flow.Steps {
		if len(step.Package) > 0 {
			// todo
		} else {
			// if the step has a function
			if len(step.Function) > 0 {
				// add exported inputs to the step
				step.Input = data.ExportInput(step.Function, m.buildFile, m.pub, true)
			} else {
				// do nothing
			}
		}
	}
}

func (m *Manager) YamlString() (string, error) {
	b, err := yaml.Marshal(m.flow)
	if err != nil {
		return "", fmt.Errorf("cannot marshal execution flow: %s", err)
	}
	return string(b), nil
}

func loadFlow(path string) (*Flow, error) {
	var err error
	if len(path) == 0 {
		return nil, fmt.Errorf("flow definition is required")
	}
	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			fmt.Errorf("cannot get absolute path for %s: %s", path, err)
		}
	}
	flowBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read flow definition %s: %s", path, err)
	}
	flow := new(Flow)
	err = yaml.Unmarshal(flowBytes, flow)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal flow definition %s: %s", path, err)
	}
	return flow, nil
}

func (m *Manager) validate() error {
	// check that the steps have the required attributes set
	for _, step := range m.flow.Steps {
		if len(step.Runtime) == 0 {
			return fmt.Errorf("invalid step %s, runtime is missing", step.Name)
		}
	}
	return nil
}