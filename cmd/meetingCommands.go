// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/jack-cheng/CLI-agenda/errors"
	"github.com/spf13/cobra"
)

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "Create a new meeting",
	Long: `To create a new meetng, please enter the unique title of your meeting,
				 which cannot be the same as any existed meeting, participators in the meeting,
				 start time and end time of your meeting with the format YYYY-MM-DD.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the arguments
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			errors.ErrorMsg("title of your meeting is required!")
		}

		participatorArg, _ := cmd.Flags().GetString("participators")
		if participatorArg == "" {
			errors.ErrorMsg("at least one participators is required!")
		}
		participators := strings.Split(participatorArg, "+")

		startTime, _ := cmd.Flags().GetString("starttime")
		if startTime == "" {
			errors.ErrorMsg("start time of your meeting is required!")
		}

		endTime, _ := cmd.Flags().GetString("endtime")
		if endTime == "" {
			errors.ErrorMsg("end time of your meeting is required!")
		}

		// TODO: check if arguments are valid and create a new meeting in the memory
		fmt.Println("Meeting " + title + " created.")
		fmt.Println("participators: " + participators[0])
		fmt.Println("start time: " + startTime)
		fmt.Println("end time: " + endTime)
	},
}

func init() {
	// add createMeeting command
	createMeetingCmd.Flags().StringP("title", "t", "", "enter the title of your meeting.")
	createMeetingCmd.Flags().StringP("participators", "p", "", "enter the participators of your meeting.")
	createMeetingCmd.Flags().StringP("starttime", "s", "", "enter the start time (YYYY-MM-DD) of your meeting.")
	createMeetingCmd.Flags().StringP("endtime", "e", "", "enter the end time (YYYY-MM-DD) of your meeting.")

	RootCmd.AddCommand(createMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
