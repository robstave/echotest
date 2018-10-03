package cmd

import (
	"echotest/conf"
	gw "echotest/gateway"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "example",
		Run: run,
	}

	rootCmd.Flags().IntP("port", "p", 0, "Port to do things on")
	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}
	fmt.Println("ppppppppp")

	fmt.Println(config.Port)
	gw.RunEcho(config.Port)
}
