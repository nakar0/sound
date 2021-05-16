package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sound",
		Short: "Play sound effects on the command line",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(successCmd)
	rootCmd.AddCommand(failCmd)
}

const (
	soundsDir = "/Users/shotaro/Documents/Google Drive/Sound"

	successSoundFile = soundsDir + "/success.mp3"
	failureSoundFile = soundsDir + "/failure.mp3"
)

var successCmd = &cobra.Command{
	Use:   "success",
	Short: "Sound Success",
	RunE: func(cmd *cobra.Command, args []string) error {
		return exec.Command("afplay", successSoundFile).Start()
	},
}

var failCmd = &cobra.Command{
	Use:   "failure",
	Short: "Sound Failure",
	RunE: func(cmd *cobra.Command, args []string) error {
		return exec.Command("afplay", failureSoundFile).Start()
	},
}
