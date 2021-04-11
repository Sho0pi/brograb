package brograb

import (
	"errors"
	"github.com/sho0pi/brograb/browserutils"
	"path/filepath"
)

const (
	chromiumBookmarksFile   = "Bookmarks"
	chromiumCookiesFile     = "Cookies"
	chromiumHistoryFile     = "History"
	chromiumPasswordsFile   = "Login Data"
	chromiumCreditCardsFile = "Web Data"
)

var (
	BadDestFile = errors.New("bad destination type")
)

// Grabber is a basic data grabbing interface. Use Next and Scan to fetch the next grabbing data.
type Grabber interface {

	// Next prepares the next fetch data for reading with the Scan method. It
	// returns true on success, or false if there is no next result row or an error
	// happened while preparing it. Err should be consulted to distinguish between
	// the two cases.
	//
	// Every call to Scan, even the first one, must be preceded by a call to Next.
	Next() bool

	// Scan copies the current fetched data into the interface pointed at
	// by dest.
	Scan(dest interface{}) error

	// Err returns the error, if any, that was encountered during iteration.
	Err() error

	// Close stops the iteration, and closes all open databases, connections,
	// and files used to fetch the browser data.
	Close() error
}

// ChromeBetaProfileDirs returns all the available Chrome Beta profiles directories in the computer.
// With the list of profile directories you will be able to create new Grabbers such as PasswordGrabber etc.
func ChromeBetaProfileDirs() ([]browserutils.ProfileDir, error) {
	return getChromiumBasedProfileDirs(browserutils.ChromeBetaProfilePath)
}

// ChromeProfileDirs returns all the available Chrome profiles directories in the computer.
// With the list of profile directories you will be able to create new Grabbers such as PasswordGrabber etc.
func ChromeProfileDirs() ([]browserutils.ProfileDir, error) {
	return getChromiumBasedProfileDirs(browserutils.ChromeProfilePath)
}

// ChromeProfileDirs returns all the available Chromium profiles directories in the computer.
// With the list of profile directories you will be able to create new Grabbers such as PasswordGrabber etc.
func ChromiumProfileDirs() ([]browserutils.ProfileDir, error) {
	return getChromiumBasedProfileDirs(browserutils.ChromiumProfilePath)
}

func getChromiumBasedProfileDirs(profilePattern string) (directories []browserutils.ProfileDir, err error) {
	histDBPattern := filepath.Join(profilePattern, chromiumHistoryFile)
	// Uses the history database file to get all the profile directories containing it.
	histDBFiles, err := filepath.Glob(histDBPattern)
	if err != nil {
		return
	}
	for _, p := range histDBFiles {
		p, _ := filepath.Split(p) // Retrieve the parent directory - the profile directory.
		directories = append(directories, browserutils.ProfileDir(p))
	}
	return
}
