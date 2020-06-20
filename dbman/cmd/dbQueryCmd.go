//   Onix Config Manager - Dbman
//   Copyright (c) 2018-2020 by www.gatblau.org
//   Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
//   Contributors to this project, hereby assign copyright in this code to the project,
//   to be licensed under the same terms as the rest of the code.
package cmd

import (
	"bytes"
	"fmt"
	"github.com/gatblau/onix/dbman/plugins"
	"github.com/gatblau/onix/dbman/util"
	"github.com/spf13/cobra"
)

type DbQueryCmd struct {
	cmd      *cobra.Command
	format   string
	filename string
}

func NewDbQueryCmd() *DbQueryCmd {
	c := &DbQueryCmd{
		cmd: &cobra.Command{
			Use:   "query [name] [args...]",
			Short: "runs a database query",
			Long:  ``,
		},
	}
	c.cmd.Run = c.Run
	c.cmd.Flags().StringVarP(&c.format, "output", "o", "json", "the format of the output - yaml, json, csv")
	c.cmd.Flags().StringVarP(&c.filename, "filename", "f", "", `if a filename is specified, the output will be written to the file. The file name should not include extension.`)

	return c
}

func (c *DbQueryCmd) Run(cmd *cobra.Command, args []string) {
	// check the query name has been passed in
	if len(args) == 0 {
		fmt.Printf("!!! You forgot to tell me the name of the query you want to run\n")
		return
	}
	// get the release manifest for the current application version
	manifest, err := util.DM.GetReleaseInfo(util.DM.Cfg.Get(plugins.AppVersion))
	if err != nil {
		fmt.Printf("!!! I cannot fetch release information: %v\n", err)
		return
	}
	var params []string
	queryName := args[0]
	if len(args) > 1 {
		params = args[1:]
	}
	// find the query definition in the manifest
	query := manifest.GetQuery(queryName)
	if query == nil {
		fmt.Printf("!!! I cannot find query: %v\n", queryName)
		return
	}
	// check the arguments passed in match the query definition
	expectedParams := len(query.Vars)
	providedParams := len(params)
	if expectedParams != providedParams {
		fmt.Printf("!!! The query expected %v parameters but %v were provided\n", varsToString(query.Vars), providedParams)
		return
	}
	result, _, err := util.DM.RunQuery(manifest, query, params)
	if err != nil {
		fmt.Printf("!!! I cannot run query '%s': %s\n", queryName, err)
		return
	}
	util.Print(result, c.format, c.filename)
}

func varsToString(vars []plugins.Var) string {
	buffer := bytes.Buffer{}
	for i, v := range vars {
		buffer.WriteString(v.Name)
		if i < len(vars)-1 {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}
