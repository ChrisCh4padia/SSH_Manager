package main

import (
	"bytes"
	_ "embed"
	"io/ioutil"
	"net/url"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	//"fyne.io/fyne/layout"
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
)

//go:embed hello.txt
var content string
var stringers string
var result1 string
var Dropdownwords string
var counter int
var wording []string
var s string

func dropdown() {
	// ################################## Profile Selection Dropdown ##################################
	counter := 0
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {

		value := scanner.Text()

		// Test our first words method.
		result1 = firstWords(value, 1)
		stringers = stringers + "." + result1

		if value != "" {
			counter = counter + 1

		}

	}

	words := SplitAny(stringers, ".")

	switch {
	case counter == 0:
		wording = []string{}
	case counter == 1:
		wording = []string{words[0]}
	case counter == 2:
		wording = []string{words[0], words[1]}
	case counter == 3:
		wording = []string{words[0], words[1], words[2]}
	case counter == 4:
		wording = []string{words[0], words[1], words[2], words[3]}
	case counter == 5:
		wording = []string{words[0], words[1], words[2], words[3], words[4]}
	case counter == 6:
		wording = []string{words[0], words[1], words[2], words[3], words[4], words[5]}
	case counter == 7:
		wording = []string{words[0], words[1], words[2], words[3], words[4], words[5], words[6]}
	case counter == 8:
		wording = []string{words[0], words[1], words[2], words[3], words[4], words[5], words[6], words[7]}
	case counter == 9:
		wording = []string{words[0], words[1], words[2], words[3], words[4], words[5], words[6], words[7], words[8]}
	case counter == 10:
		wording = []string{words[0], words[1], words[2], words[3], words[4], words[5], words[6], words[7], words[8], words[9]}
	case counter == 11:
		wording = []string{words[0], words[1], words[2], words[3], words[4], words[5], words[6], words[7], words[8], words[9], words[10]}
	}
	fmt.Println(words)
}

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func firstWords(value string, count int) string {
	// Loop over all indexes in the string.
	for i := range value {
		// If we encounter a space, reduce the count.
		if value[i] == ' ' {
			count -= 1
			// When no more words required, return a substring.
			if count == 0 {
				return value[0:i]
			}
		}
	}
	// Return the entire string.
	return value
}

func main() {

	spacer := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	label1 := widget.NewLabel("Please enter Serveraddress or Hostname:")
	Serveraddress := widget.NewEntry()
	label2 := widget.NewLabel("Please enter your Username, if neccesary:")
	Username := widget.NewEntry()
	label3 := widget.NewLabel("If required, please enter the Path to your SSH Key File:")
	Password := widget.NewEntry()
	labelProfileinfo := widget.NewLabel("Profiles are used to connect to network Devices quicker.")
	labelRemoveProfile := widget.NewLabel("Choose Profile to Remove:")

	btn03 := widget.NewButtonWithIcon("Browse SSH Key File", theme.FolderIcon(), func() {
		sshfile, err := dialog.File().Filter("All Files", "*").Load()
		fmt.Println(sshfile)
		fmt.Println("Error:", err)
		Password.SetText(sshfile)
	})
	btn01 := widget.NewButtonWithIcon("Connect", theme.ConfirmIcon(), func() {
		if Serveraddress.Text == "" {
			spacer.SetText("widget.NewEntry 'ServerAddress' can not be Empty!")
		} else {
			spacer.SetText("")
			if Username.Text == "" {
				os := runtime.GOOS
				switch {
				case os == "linux":
					cmd := exec.Command("konsole", "-e", "ssh", Serveraddress.Text)
					if err := cmd.Start(); err != nil {
						log.Fatal(err)
					}
				case os == "windows":
					cmd := exec.Command("cmd", "/C", "start", "ssh", Serveraddress.Text)
					if err := cmd.Start(); err != nil {
						log.Fatal(err)
					}
				default:
					fmt.Println("Your Platform is not supported, please either use Linux or Windows")
				}
			}
		}
		if Username.Text != "" && Password.Text == "" {
			os := runtime.GOOS
			switch {
			case os == "linux":
				linuxcombstring := Username.Text + "@" + Serveraddress.Text
				cmd := exec.Command("konsole", "-e", "ssh", linuxcombstring)
				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
			case os == "windows":
				wincombstring := Username.Text + "@" + Serveraddress.Text
				cmd := exec.Command("cmd", "/C", "start", "ssh", wincombstring)
				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
			default:
				fmt.Println("Your Platform is not supported, please either use Linux or Windows")
			}
		}
		if Username.Text != "" && Password.Text != "" {
			fmt.Println("Testerei")
			os := runtime.GOOS
			switch {
			case os == "linux":
				linuxcombstring := Username.Text + "@" + Serveraddress.Text
				cmd := exec.Command("konsole", "-e", "ssh", "-i", Password.Text, linuxcombstring)
				fmt.Println("konsole", "-e", "ssh", "-i", Password.Text, linuxcombstring)
				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
			case os == "windows":
				wincombstring := Username.Text + "@" + Serveraddress.Text
				cmd := exec.Command("cmd", "/C", "start", "ssh", "-i", Password.Text, wincombstring)
				fmt.Println("cmd", "/C", "start", "ssh", "-i", Password.Text, wincombstring)
				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
			default:
				fmt.Println("Your Platform is not supported, please either use Linux or Windows")
			}
		}

	})
	btn02 := widget.NewButtonWithIcon("Cancel", theme.ContentClearIcon(), func() {
		Serveraddress.SetText("")
		Password.SetText("")
		Username.SetText("")
	})
	Profileselection := widget.NewSelect(
		wording,
		func(s string) {

			fmt.Printf(s)
		})
	ProfAdderror := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	LabelProfName := widget.NewLabel("Please Enter ProfileName:")
	profName := widget.NewEntry()
	profServeraddress := widget.NewEntry()
	profUsername := widget.NewEntry()
	profPassword := widget.NewEntry()
	AddProfileFile := widget.NewButtonWithIcon("Browse SSH Key File", theme.FolderIcon(), func() {
		Proffile, err := dialog.File().Filter("All Files", "*").Load()
		fmt.Println(Proffile)
		fmt.Println("Error:", err)
		profPassword.SetText(Proffile)
		dropdown()
		Profileselection.Refresh()

	})
	dropdown()

	AddProfButton := widget.NewButtonWithIcon("Add Profile", theme.ContentAddIcon(), func() {

		if profName.Text == "" {
			ProfAdderror.SetText("widget.NewEntry 'NewProfileName' can not be Empty!")
		} else {

			f, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}

			defer f.Close()
			NewProfinput := profName.Text + " " + profServeraddress.Text + " " + profUsername.Text + " " + profPassword.Text + "\n"
			if _, err = f.WriteString(NewProfinput); err != nil {
				panic(err)
				dropdown()
				Profileselection.Refresh()
			}
			dropdown()
			ProfAdderror.SetText("Profile" + " " + profName.Text + " has been created!")
			profName.SetText("")
			profServeraddress.SetText("")
			profUsername.SetText("")
			profPassword.SetText("")
			Profileselection.Refresh()
		}
		Profileselection.Refresh()
	})
	spacerRem := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	LabelProfRemName := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	LabelProfRemName1 := widget.NewLabel("Selected Profile to delete:")

	ProfileRemselection := widget.NewSelect(
		wording,
		func(s string) {
			LabelProfRemName.Text = s
			LabelProfRemName.Refresh()
		})

	btnRemoveProfile := widget.NewButtonWithIcon("Remove Profile", theme.ContentRemoveIcon(), func() {
		if LabelProfRemName.Text == "" {
			LabelProfRemName.SetText("Please select a valid Profile!")
		}
		if LabelProfRemName.Text == "Please select a valid Profile!" {
			LabelProfRemName.SetText("Please select a valid Profile!")
		} else {
			name := LabelProfRemName.Text
			spacerRem.SetText("Profile " + name + " deleted!")
			input, err := ioutil.ReadFile("hello.txt")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			output := bytes.Replace(input, []byte(name), []byte(""), -1)
			if err = ioutil.WriteFile("hello.txt", output, 0666); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			LabelProfRemName.Text = ""
		}
	})
	IconWidget := widget.NewLabelWithStyle("Icon from:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	iconurl, _ := url.Parse("https://vectorified.com/download-image#ssh-icon-7.png")
	iconlink := widget.NewHyperlink("https://vectorified.com/", iconurl)

	url, _ := url.Parse("https://github.com/ChrisCh4padia/")
	hyperlink := widget.NewHyperlink("github.com/ChrisCh4padia/", url)
	Helplabel := widget.NewLabelWithStyle("Documentation can be found under:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	contenttab1 := widget.NewVBox(spacer, label1, Serveraddress, label2, Username, label3, Password, btn03, btn01, btn02)
	contenttab3 := widget.NewVBox(labelProfileinfo, LabelProfName, profName, label1, profServeraddress, label2, profUsername, label3, profPassword, AddProfileFile, AddProfButton, ProfAdderror)
	contenttab4 := widget.NewVBox(labelRemoveProfile, ProfileRemselection, LabelProfRemName1, LabelProfRemName, spacerRem, btnRemoveProfile)
	contenttab5 := widget.NewVBox(IconWidget, iconlink)
	contenttab6 := widget.NewVBox(Helplabel, hyperlink)
	app := app.New()
	w := app.NewWindow("SSH-Manager")
	icon, _ := fyne.LoadResourceFromPath("C:/Users/crenner/Downloads/ssh-icon-7.png")
	w.SetIcon(icon)
	var tabs *widget.TabContainer
	var tab1, tab3, tab4, tab2, tab5 *widget.TabItem
	tab1 = widget.NewTabItemWithIcon("Home", theme.HomeIcon(), contenttab1)
	tab3 = widget.NewTabItemWithIcon("Add Profiles", theme.ContentAddIcon(), contenttab3)
	tab4 = widget.NewTabItemWithIcon("Delete Profiles", theme.ContentRemoveIcon(), contenttab4)
	tab5 = widget.NewTabItemWithIcon("Info", theme.InfoIcon(), contenttab5)
	tab2 = widget.NewTabItemWithIcon("Help", theme.HelpIcon(), contenttab6)
	tabs = widget.NewTabContainer(tab1, tab3, tab4, tab2, tab5)
	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)
	w.Resize(fyne.Size{600, 600})
	w.ShowAndRun()

}
