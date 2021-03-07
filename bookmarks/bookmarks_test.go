package bookmarks

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
	"time"
)

var (
	emptyResult = gjson.Parse(`{}`)
)

func TestBookmark_setDateAdded(t *testing.T) {
	assert := assert.New(t)
	var bookmark Bookmark

	date_added := 13259621094000000
	date := time.Date(2021, 3, 7, 20, 4, 54, 0, time.UTC)

	validJson := fmt.Sprintf(`{"date_added": %d}`, date_added)
	entry := gjson.Parse(validJson)

	assert.NoError(bookmark.setDateAdded(entry))
	assert.Equal(bookmark.DateAdded, date)

	assert.Error(bookmark.setDateAdded(emptyResult))
}

func TestBookmark_setName(t *testing.T) {
	assert := assert.New(t)
	var bookmark Bookmark
	name := "Hello World!"

	validJson := fmt.Sprintf(`{"name": "%s"}`, name)
	entry := gjson.Parse(validJson)

	assert.NoError(bookmark.setName(entry))
	assert.Equal(bookmark.Name, name)

	assert.Error(bookmark.setName(emptyResult))
}

func TestBookmark_setChildren(t *testing.T) {
	assert := assert.New(t)
	var bookmark Bookmark

	children := []interface{}{"hello", "world"}

	validJson := fmt.Sprintf(`{"children": ["%s", "%s"]}`, children...)
	entry := gjson.Parse(validJson)

	assert.NoError(bookmark.setChildren(entry))
	assert.Equal(bookmark.index, -1)
	assert.Equal(len(bookmark.bookmarks), 2)
	assert.Equal(bookmark.bookmarks[0].String(), children[0])

	assert.Error(bookmark.setChildren(emptyResult))

}

func TestBookmark_setURL(t *testing.T) {
	assert := assert.New(t)
	var bookmark Bookmark
	url := "http://github.com/sho0pi/brograb"

	validJson := fmt.Sprintf(`{"url":"%s"}`, url)
	entry := gjson.Parse(validJson)

	assert.NoError(bookmark.setURL(entry))
	assert.Equal(bookmark.URL, url)

	assert.Error(bookmark.setURL(emptyResult))
}

func TestBookmark_setType(t *testing.T) {
	assert := assert.New(t)
	var bookmark Bookmark

	folderType := `{"type": "folder"}`
	entry := gjson.Parse(folderType)

	assert.NoError(bookmark.setType(entry))
	assert.Equal(bookmark.Type, FOLDER)

	urlType := `{"type": "url"}`
	entry = gjson.Parse(urlType)

	assert.NoError(bookmark.setType(entry))
	assert.Equal(bookmark.Type, URL)

	assert.Error(bookmark.setType(emptyResult))
}
