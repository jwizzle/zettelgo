/*
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jwizzle/zettelgo/zettels"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new zettel.",
	Long: `Create a new zettel, takes a title. The new file is placed inside
Your zettelkasten directory (see config). Its filename is the title prefixed
by a datestring to make it unique. 
It's initially filled with a header template.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		title := strings.ReplaceAll(args[0], " ", "_")
		currentTime := time.Now()
		formattedDate := currentTime.Format("150405_20060102")
		filename := fmt.Sprintf("%s_%s%s", formattedDate, title, zettelCfg.Note_suffix)
		filepath := fmt.Sprintf("%s/%s", zettelCfg.Directory, filename)

		newheader := zettels.Header{
			Title: title,
			Date: currentTime.Format("2006-01-02T15:04"),
			Tags: []string{"#tag"},
			Links: map[string]string{"description": "[[link]]"},
		}
		headerstring, _ := newheader.Display()

		newfile, err := os.Create(filepath)
		if err != nil {
			panic(err)
		}
		defer newfile.Close()
		newfile.WriteString(fmt.Sprintf("---\n%s---", headerstring))

		fmt.Println(filepath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
