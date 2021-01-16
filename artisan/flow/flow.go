/*
  Onix Config Manager - Artisan
  Copyright (c) 2018-2021 by www.gatblau.org
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/
package flow

import (
	"github.com/gatblau/onix/artisan/core"
	"github.com/gatblau/onix/artisan/data"
)

// a set of authentication credentials for a package registry
type Credential struct {
	User     string
	Password string
	Domain   string
}

type Flow struct {
	Name        string        `yaml:"name"`
	Description string        `yaml:"description"`
	GitURI      string        `yaml:"git_uri,omitempty"`
	AppIcon     string        `yaml:"app_icon,omitempty"`
	Steps       []*Step       `yaml:"steps"`
	Credential  []*Credential `yaml:"credential,omitempty"`
}

// finds if the domain contains credentials
func (f *Flow) HasDomain(domain string) bool {
	for _, credential := range f.Credential {
		if credential.Domain == domain {
			return true
		}
	}
	return false
}

func (f *Flow) StepByFx(fxName string) *Step {
	for _, step := range f.Steps {
		if step.Function == fxName {
			return step
		}
	}
	return nil
}

func (f *Flow) RequiresSource() bool {
	for _, step := range f.Steps {
		if len(step.Package) == 0 && len(step.Function) > 0 {
			return true
		}
	}
	return false
}

func (f *Flow) RequiresKey() bool {
	for _, step := range f.Steps {
		if step.Input != nil && step.Input.Key != nil {
			return true
		}
	}
	return false
}

func (f *Flow) RequiresSecrets() bool {
	for _, step := range f.Steps {
		if step.Input != nil && step.Input.Secret != nil {
			return true
		}
	}
	return false
}

func (f *Flow) GetRegistryUser(packageName string) string {
	name, err := core.ParseName(packageName)
	core.CheckErr(err, "cannot parse package name %s", packageName)
	for _, credential := range f.Credential {
		if credential.Domain == name.Domain {
			return credential.User
		}
	}
	return ""
}

func (f *Flow) GetRegistryPwd(packageName string) string {
	name, err := core.ParseName(packageName)
	core.CheckErr(err, "cannot parse package name %s", packageName)
	for _, credential := range f.Credential {
		if credential.Domain == name.Domain {
			return credential.Password
		}
	}
	return ""
}

type Step struct {
	Name            string      `yaml:"name"`
	Description     string      `yaml:"description,omitempty"`
	Runtime         string      `yaml:"runtime"`
	RuntimeManifest string      `yaml:"runtime_manifest,omitempty"`
	Function        string      `yaml:"function,omitempty"`
	Package         string      `yaml:"package,omitempty"`
	Input           *data.Input `yaml:"input,omitempty"`
}
