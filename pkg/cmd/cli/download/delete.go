package download

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vmware-tanzu/velero/pkg/client"
	"github.com/vmware-tanzu/velero/pkg/cmd/util/output"
	"github.com/vmware-tanzu/velero/pkg/util/logging"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewDeleteCommand(f client.Factory, use string) *cobra.Command {
	var listOptions metav1.ListOptions
	logLevelFlag := logging.LogLevelFlag(logrus.InfoLevel)
	formatFlag := logging.NewFormatFlag()

	c := &cobra.Command{
		Use:  fmt.Sprintf("%s [NAMES]", use),
		Short: "Delete downloads",
		Run: func(c *cobra.Command, args []string) {
			logLevel := logLevelFlag.Parse()
			logger := logging.DefaultLogger(logLevel, formatFlag.Parse())
			logger.Infof("The command, datamgr download delete, is called")
		},
	}

	c.Flags().StringVarP(&listOptions.LabelSelector, "selector", "l", listOptions.LabelSelector, "only show items matching this label selector")

	output.BindFlags(c.Flags())

	return c
}