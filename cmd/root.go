package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	RootCMD = &cobra.Command{
		Use: "avvkctl",
		Short: "Use it",
		Run: nil,
	}
	NameCMD = &cobra.Command{
		Use: "name",
		Short: "print name",
		Run: printName,
	}
)

func printName(cmd *cobra.Command, args []string)  {
	fmt.Println("vipin")
}

func init()  {
	RootCMD.AddCommand(NameCMD)
}