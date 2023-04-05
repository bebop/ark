package cmd

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:  "ark",
    Short: "ark - empowered by Earth Biosphere to guide your organism engineering",
    Long: `    Ark is a database builder and a utility belt to make your bioengineering efforts easier, enjoyful and fluid.

    With Ark you can go from an organism-specific database to help you on metabolic engineering, to a full search on the large database of sequenced life.

    Choose your adventure, pick your tools and good engineering!`,
    Run: func(cmd *cobra.Command, args []string) {
      /*
       * The list of steps right now:
       * - [ ] Check if a config file is present
       * - [ ] If not found go to interactive config setup
       * - [ ] Run the database setup based on config
       */

       fmt.Println("Hi, I'm ark")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}
