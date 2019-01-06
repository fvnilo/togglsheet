package app

import "github.com/spf13/cobra"

type command func(cmd *cobra.Command, args []string)
