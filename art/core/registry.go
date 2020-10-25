/*
  Onix Config Manager - Art
  Copyright (c) 2018-2020 by www.gatblau.org
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/
package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// a repository in the registry
type repository struct {
	// the readable name of the artefact
	Name string `json:"name"`
	// the reference name of the artefact corresponding to different builds
	Artefacts []*artefact `json:"artefacts"`
}

// get the artefact by name or nil if not found in the registry
func (repo *repository) artefact(name string) (*artefact, bool) {
	for _, artefact := range repo.Artefacts {
		if artefact.Name == name {
			return artefact, true
		}
	}
	return nil, false
}

type artefact struct {
	// the internal reference name of the artefact
	Name string `json:"name"`
	// the list of Tags associated with the artefact
	Tags []string `json:"tags"`
}

// the local registry containing the repositories
type registry struct {
	Repositories []*repository `json:"repositories"`
}

// get the repository by name or nil if not found in the registry
func (r *registry) repo(name string) (*repository, bool) {
	for _, repo := range r.Repositories {
		if repo.Name == name {
			return repo, true
		}
	}
	return nil, false
}

// create a registry management structure
func NewRegistry() *registry {
	r := &registry{
		Repositories: []*repository{},
	}
	// load registry
	r.load()
	return r
}

// load the content of the registry
func (r *registry) load() {
	// check if registry file exist
	_, err := os.Stat(r.file())
	if err != nil {
		// then assume registry.json is not there: try and create it
		r.save()
	}
}

// the local path to the registry
func (r *registry) path() string {
	return fmt.Sprintf("%s/.%s", homeDir(), cliName)
}

// return the registry full file name
func (r *registry) file() string {
	return fmt.Sprintf("%s/registry.json", r.path())
}

// save the state of the registry
func (r *registry) save() {
	regBytes, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(r.file(), regBytes, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// add the file and seal to the registry
func (r *registry) add(filename string, name string) {
	// gets the full base name (with extension)
	basename := filepath.Base(filename)
	// gets the basename directory only
	basenameDir := filepath.Dir(filename)
	// gets the base name extension
	basenameExt := path.Ext(basename)
	// gets the base name without extension
	basenameNoExt := strings.TrimSuffix(basename, path.Ext(basename))
	// if the file to add is not a zip file
	if basenameExt != ".zip" {
		log.Fatal(errors.New(fmt.Sprintf("the registry can only accept zip files, the extension provided was %s", basenameExt)))
	}
	// move the zip file to the registry folder
	err := os.Rename(filename, fmt.Sprintf("%s/%s", r.path(), basename))
	if err != nil {
		log.Fatal(err)
	}
	// now move the seal file to the registry folder
	err = os.Rename(fmt.Sprintf("%s/%s.json", basenameDir, basenameNoExt), fmt.Sprintf("%s/%s.json", r.path(), basenameNoExt))
	if err != nil {
		log.Fatal(err)
	}
	// does the artefact exist?
	if repo, exists := r.repo(name); exists {
		// does the reference exists in the artefact?
		if _, artefactExists := repo.artefact(basenameNoExt); artefactExists {
			log.Fatal(fmt.Sprintf("cannot add duplicate artefact %s to registry", basenameNoExt))
		} else {
			// creates the reference
			art := &artefact{
				Name: basenameNoExt,
				Tags: nil,
			}
			// adds the reference to the artefact
			arts := append(repo.Artefacts, art)
			repo.Artefacts = arts
		}
	} else {
		// creates a new artefact
		repos := append(r.Repositories, &repository{
			Name: name,
			Artefacts: []*artefact{
				{
					Name: basenameNoExt,
					Tags: nil,
				},
			},
		})
		r.Repositories = repos
	}
	// persist the changes
	r.save()
}
