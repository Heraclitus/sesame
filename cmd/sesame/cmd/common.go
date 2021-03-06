package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/cobra"
	"os"
	"runtime/debug"
)

func ValidateArgsFunc() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("unexpected argument [%s]", args[0])
		}
		return nil
	}
}

func exitOnError(err error) {
	if err != nil {
		_, err := fmt.Fprintln(os.Stdout, err)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

type SesameError struct {
	msg         string
	reportStack bool
}

func (m *SesameError) Error() string {
	if m.reportStack {
		debug.PrintStack()
	}
	return m.msg
}

type SSMCommand struct {
	svc *ssm.SSM
}

func (ssmCommand *SSMCommand) conf() {
	profile := getAwsProfile()
	config := &aws.Config{
		Region:                        aws.String(getAwsRegion()),
		CredentialsChainVerboseErrors: aws.Bool(true),
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            *config,
		Profile:           profile,
	})
	exitOnError(err)
	ssmCommand.svc = ssm.New(sess)
}

func (ssmCommand *SSMCommand) thingDo() {
	os.Exit(200)
}
