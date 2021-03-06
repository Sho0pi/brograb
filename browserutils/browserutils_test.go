package browserutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	profilePath    = "path/to/profile"
	profileDirName = "profile"
)

func TestProfileDir(t *testing.T) {
	assert := assert.New(t)
	profileDir := ProfileDir(profilePath)

	assert.Equal(profileDir.BookmarksDB(), fmt.Sprintf("%s/%s", profilePath, chromiumBookmarksFile))
	assert.Equal(profileDir.CookiesDB(), fmt.Sprintf("%s/%s", profilePath, chromiumCookiesFile))
	assert.Equal(profileDir.CreditCardsDB(), fmt.Sprintf("%s/%s", profilePath, chromiumCreditCardsFile))
	assert.Equal(profileDir.HistoryDB(), fmt.Sprintf("%s/%s", profilePath, chromiumHistoryFile))
	assert.Equal(profileDir.PasswordsDB(), fmt.Sprintf("%s/%s", profilePath, chromiumPasswordsFile))
	assert.Equal(profileDir.Path(), profilePath)
	assert.Equal(profileDir.Name(), profileDirName)
}

func TestFormatChromiumEpoch(t *testing.T) {
	assert := assert.New(t)

	date_added := 13259621094000000
	date := time.Date(2021, 3, 7, 20, 4, 54, 0, time.UTC)

	assert.Equal(FormatChromiumEpoch(int64(date_added)), date)
}
