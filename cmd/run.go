package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/paulwrubel/stroll/internal"
	"github.com/spf13/cobra"
)

var (
	debug bool
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run FILENAME",
	Short: "run will execute a stroll",
	Long: `Run stroll files (.strl) or input directly.
	
This will run a single stroll file.

If the first argument is '-', it will read from stdin for the stroll text.`,
	Args: cobra.RangeArgs(1, 2),
	Run:  runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
}

func runRun(cmd *cobra.Command, args []string) {
	// read input
	strollBytes, err := readStrollBytes(args[0])
	if err != nil {
		fmt.Printf("error reading stroll bytes: %s\n", err.Error())
		os.Exit(1)
	}

	strollArgs := ""
	if len(args) == 2 {
		strollArgs = args[1]
	}

	// run the program
	stroll, err := internal.NewStroll(string(strollBytes), strollArgs, debug)
	if err != nil {
		fmt.Printf("error planning stroll: %s\n", err.Error())
		os.Exit(1)
	}
	err = stroll.Execute()
	if err != nil {
		fmt.Printf("error strolling: %s.\n", err.Error())
		os.Exit(1)
	}

}

func readStrollBytes(filename string) ([]byte, error) {
	var reader io.Reader
	if filename == "-" {
		reader = os.Stdin
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("error opening stroll file: %s", err.Error())
		}
		reader = file
	}

	strollBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading stroll: %s", err.Error())
	}

	return strollBytes, nil
}
