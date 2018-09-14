package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/donbattery/kkhcli/utils"
)

// aaddCmd represents the aadd command
var (
	aaddCmd = &cobra.Command{
		Use:     "add",
		Short:   "Adds new avatar(s) by path to the KKHC database",
		Aliases: []string{"a"},
		Long: `kkhcli avatar add sends new avatar(s) by path to the KKHC database

Example:
	kkhcli avatar add . (sends *.png from pwd to the kkhc server)
	kkhcli avatar add ~/Documents/Pictures/Avatars ~/Desktop/Avatars (sends *.png from your Documents/Pictures/Avatars and Desktop/Avatars to the kkhc server)`,
		Run: func(cmd *cobra.Command, args []string) {
			var msg utils.Response
			for _, arg := range args {
				getPNGs(arg)
			}
			if len(avatarList) == 0 {
				log.Fatal("Not a single PNG was found")
			}
			for _, avatar := range avatarList {
				resp, err := utils.SendAvatar(avatar)
				if err != nil {
					fmt.Printf("Error sending avatar picture %s\n%s\n", avatar, err)
				} else {
					if err := json.Unmarshal(resp, &msg); err == nil {
						fmt.Printf("%v\n", msg.Msg)
					}
					uploaded++
				}
			}
			fmt.Printf("%v avatars has been uploaded to the KKHC Server\n", uploaded)
		},
	}
	avatarList []string
	uploaded   int
)

func init() {
	avatarCmd.AddCommand(aaddCmd)
}

func visit(path string, f os.FileInfo, err error) error {
	if pathType, _ := checkPath(path); pathType == "png" {
		avatarList = append(avatarList, path)
	}
	return nil
}

func getPNGs(arg string) {

	pathType, _ := checkPath(arg)

	switch pathType {
	case "dir":
		filepath.Walk(arg, visit)
	case "png":
		avatarList = append(avatarList, arg)
	}

}

func isPNG(path string) bool {
	var size int64
	// Get File
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	// Get proper sized header
	fi, err := file.Stat()
	if err != nil {
		return false
	}
	if fi.Size() < 512 {
		size = fi.Size()
	} else {
		size = 512
	}
	header := make([]byte, size)

	if _, err := io.ReadFull(file, header[:]); err != nil {
		return false
	}
	fmt.Printf("\nHEADER\n%v\n", header)
	return http.DetectContentType(header) == "image/png"
}

func checkPath(path string) (string, error) {
	file, err := os.Stat(path)
	if err != nil {
		return "", errors.New("Cannot open path")
	}
	switch mode := file.Mode(); {
	case mode.IsDir():
		return "dir", nil
	case mode.IsRegular():
		if !isPNG(path) {
			return "", errors.New("Not a PNG file")
		}
		return "png", nil
	}
	return "", errors.New("Unknown type")
}
