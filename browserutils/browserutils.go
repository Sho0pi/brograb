package browserutils

import (
	"path/filepath"
)

const (
	chromiumBookmarksFile   = "Bookmarks"
	chromiumCookiesFile     = "Cookies"
	chromiumHistoryFile     = "History"
	chromiumPasswordsFile   = "Login Data"
	chromiumCreditCardsFile = "Web Data"
)

// ProfileDir represents a profile path of specific profile on chromium based systems.
type ProfileDir string

// Path returns the path of the profile directory.
func (p ProfileDir) Path() string {
	return string(p)
}

// PasswordsDB returns the path to the chromium database containing the login data.
func (p ProfileDir) PasswordsDB() string {
	return filepath.Join(p.Path(), chromiumPasswordsFile)
}

// CreditCardsDB returns the path to the chromium database containing the saved credit cards.
func (p ProfileDir) CreditCardsDB() string {
	return filepath.Join(p.Path(), chromiumCreditCardsFile)
}

// HistoryDB returns the path to the history database.
func (p ProfileDir) HistoryDB() string {
	return filepath.Join(p.Path(), chromiumHistoryFile)
}

// HistoryDB returns the path to the cookies database.
func (p ProfileDir) CookiesDB() string {
	return filepath.Join(p.Path(), chromiumCookiesFile)
}

// HistoryDB returns the path to the bookmarks database.
func (p ProfileDir) BookmarksDB() string {
	return filepath.Join(p.Path(), chromiumBookmarksFile)
}

// GetChromiumProfileDirs returns all the available Chromium profiles directories in the computer.
// With the list of profile directories you will be able to create new Grabbers such as PasswordGrabber etc.
func GetChromiumProfileDirs() ([]ProfileDir, error) {
	return getChromiumProfileDirs(chromiumProfilePath)
}

func getChromiumProfileDirs(profilePattern string) (directories []ProfileDir, err error) {
	histFile := filepath.Join(profilePattern, chromiumHistoryFile)
	histDBFiles, err := filepath.Glob(histFile)
	if err != nil {
		return
	}
	for _, p := range histDBFiles {
		p, _ := filepath.Split(p) // Retrieve the parent directory of the history file -> usually `Default`
		directories = append(directories, ProfileDir(p))
	}
	return
}
