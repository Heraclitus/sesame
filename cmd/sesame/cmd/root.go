package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const DefaultAwsRegion = "us-east-1"
const DefaultProfile = "default"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sesame",
	Short: "A brief description of your application",
	Long:  `Sesame makes simple actions simple for AWS simple systems manager.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sesame.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getAwsRegion() string {
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		awsRegion = DefaultAwsRegion
	}
	_, err := fmt.Fprintf(os.Stderr, "AWS region: [%s]\n", awsRegion)
	if err != nil {
		panic(err)
	}
	return awsRegion
}

func getAwsProfile() string {
	awsProfile := os.Getenv("AWS_PROFILE")
	if awsProfile == "" {
		awsProfile = DefaultProfile
	}
	_, err := fmt.Fprintf(os.Stderr, "AWS profile: [%s]\n", awsProfile)
	if err != nil {
		panic(err)
	}
	return awsProfile
}
