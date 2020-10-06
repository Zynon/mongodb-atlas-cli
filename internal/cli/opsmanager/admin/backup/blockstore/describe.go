// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package blockstore

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

var describeTemplate = `ID	URI	SSL	LOAD FACTOR
{{.ID}}	{{.URI}}	{{.SSL}}	{{.LoadFactor}}
`

type DescribeOpts struct {
	cli.OutputOpts
	store        store.BlockstoresDescriber
	blockstoreID string
}

func (opts *DescribeOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.DescribeBlockstore(opts.blockstoreID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli ops-manager admin backup blockstore(s) describe <blockstoreID>
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	opts.Template = describeTemplate
	cmd := &cobra.Command{
		Use:     "describe <blockstoreID>",
		Aliases: []string{"get"},
		Short:   describe,
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.blockstoreID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
