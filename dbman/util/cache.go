//   Onix Config Db - Dbman
//   Copyright (c) 2018-2020 by www.gatblau.org
//   Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
//   Contributors to this project, hereby assign copyright in this code to the project,
//   to be licensed under the same terms as the rest of the code.
package util

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"strings"
)

// dbman root configuration management struct
type Cache struct {
	cfg *viper.Viper
}

// create a new Cache instance
func NewCache() *Cache {
	rootCfg := &Cache{
		cfg: viper.New(),
	}
	rootCfg.load()
	return rootCfg
}

// load the Cache
func (c *Cache) load() {
	home := homeDir()
	c.cfg.AddConfigPath(home)
	c.cfg.SetConfigName(".dbman")
	c.cfg.SetConfigType("toml")
	err := c.cfg.ReadInConfig()
	if err != nil {
		c.setPath(home)
		c.setName("default")
		c.save()
		err := c.cfg.ReadInConfig()
		if err != nil {
			fmt.Printf("oops! I cannot save root configuration: %v", err)
		}
	}
}

// save the config path & name to file
func (c *Cache) save() {
	err := c.cfg.WriteConfig()
	if err != nil {
		// the file does not exist then try create it
		err := c.cfg.SafeWriteConfig()
		if err != nil {
			fmt.Printf("oops! I cannot save cache: %v", err)
		}
	}
}

// return the config file path to use
func (c *Cache) path() string {
	return c.cfg.GetString("path")
}

// return the config name to use
func (c *Cache) name() string {
	return c.cfg.GetString("name")
}

// return the config file name to use without extension
func (c *Cache) filename() string {
	return fmt.Sprintf(".dbman_%v", c.name())
}

func (c *Cache) setName(name string) {
	// check file name should not contain a file extension
	if strings.Contains(name, ".toml") ||
		strings.Contains(name, ".json") ||
		strings.Contains(name, ".yaml") ||
		strings.Contains(name, ".yml") ||
		strings.Contains(name, ".yaml") ||
		strings.Contains(name, ".txt") {
		fmt.Printf("oops! I found a file extension in the configuration filename '%v': it should not contain any extension", name)
		return
	}

	// invalid if the file name is prepended by '.'
	if strings.Index(name, ".") == 1 {
		fmt.Printf("oops! I found an invalid name '%v': it should not start with '.'", name)
		return
	}
	c.cfg.Set("name", name)
}

func (c *Cache) setPath(value string) {
	c.cfg.Set("path", value)
}

// get the home directory
func homeDir() string {
	// find home directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("oops! I cannot find the home directory: %v\n", err)
		return ""
	}
	return home
}
