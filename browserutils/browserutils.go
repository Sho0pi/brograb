package browserutils

import (
	"github.com/sho0pi/brograb/bookmarks"
	"github.com/sho0pi/brograb/history"
	"path/filepath"
	"time"
)

const (
	chromiumBookmarksFile   = "Bookmarks"
	chromiumCookiesFile     = "Cookies"
	chromiumHistoryFile     = "History"
	chromiumPasswordsFile   = "Login Data"
	chromiumCreditCardsFile = "Web Data"
)

// ProfileDir represents a profile directory of specific profile on chromium based systems.
type ProfileDir string

func (p ProfileDir) BookmarksGrabber(area bookmarks.BookmarkArea) (*bookmarks.Bookmark, error) {
	return bookmarks.NewChromiumGrabber(p, area)
}

func (p ProfileDir) HistoryGrabber(date time.Time) (*history.ChromiumGrabber, error) {
	return history.NewChromiumGrabber(p, date)
}

// Path returns the path of the profile directory.
func (p ProfileDir) Path() string {
	return string(p)
}

// Name returns the name of the profile directory.
//
// Note: In case of multiply users in the system, the Name still could be the same. (Usually: Default)
func (p ProfileDir) Name() string {
	return filepath.Base(p.Path())
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

// FormatChromiumEpoch format the epoch from the chromium db, to a human readable format.
func FormatChromiumEpoch(epoch int64) time.Time {
	t := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
	d := time.Duration(epoch)
	for i := 0; i < 1000; i++ {
		t = t.Add(d)
	}
	return t
}
