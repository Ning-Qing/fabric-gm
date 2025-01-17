/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package node

import (
	"github.com/VoneChain-CS/fabric-gm/core/ledger/kvledger"
	"github.com/VoneChain-CS/fabric-gm/internal/peer/common"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"path/filepath"
)

func resumeCmd() *cobra.Command {
	resumeChannelCmd.ResetFlags()
	flags := resumeChannelCmd.Flags()
	flags.StringVarP(&channelID, "channelID", "c", common.UndefinedParamValue, "Channel to resume.")
	flags.StringVarP(&rootFSPath, "rootFSPath", "p", common.UndefinedParamValue, "File system path")

	return resumeChannelCmd
}

var resumeChannelCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resumes a channel on the peer.",
	Long:  `Resumes a channel on the peer. When the command is executed, the peer must be offline. When the peer starts after resume, it will receive blocks for the resumed channel.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if channelID == common.UndefinedParamValue {
			return errors.New("Must supply channel ID")
		}
		if rootFSPath == common.UndefinedParamValue {
			return errors.New("Must supply file system path")
		}
		ledgersPath := filepath.Join(rootFSPath, "ledgersData")
		return kvledger.ResumeChannel(ledgersPath, channelID)
	},
}
