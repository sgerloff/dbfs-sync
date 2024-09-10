package main

import (
	"context"
	"fmt"
	"os"

	"github.com/databricks/cli/cmd/root"
	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()
	cmd := &cobra.Command{
		Use:           "dbfs-sync [flags] SRC DST",
		Short:         "Databricks CLI extension for syncing files to DBFS",
		Version:       "0.0.0",
		SilenceUsage:  false,
		SilenceErrors: false,
	}
	cmd.SetContext(ctx)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println(args)
		return nil
	}

	err := root.Execute(ctx, cmd)
	if err != nil {
		os.Exit(1)
	}
}
