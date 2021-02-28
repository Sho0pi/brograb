# BroGrab

brograb is an open-source Go package that help your go modules fetch and grab browser data
(passwords, bookmarks, cookies, history etc...) from all the common browsers. Our goal is to create cross platform API
to fetch and decrypt the browser data. 

This project is heavily inspire by [HackBrowserData](https://github.com/moonD4rk/HackBrowserData)

> Statement: This package is limited to security research only, and the user assumes all legal and related
> responsibilities arising from its use! The author assumes no legal responsibility!


## Installing

To start using BroGrab, run `go get`:

```shell
$ go get github.com/sho0pi/brograb
```

## Example

To iterate on Chrome passwords is simple as:

```go
package main

import "github.com/sho0pi/brograb/passwords"

func main() {
	grabber, err := passwords.NewChromeGrabber()

	// Using sql like API
	if grabber.Next() {
		var password PasswordEntry
		err := grabber.Scan(&password)
		println(password)
	}

	// Using concurrency and go channels
	for password := grabber.Iterate() {
		println(password.URL, password.Username, password.Password)
	}
}

```

## TODO

- [] Option to grab passwords.
- [] Option to grab bookmarks.
- [] Option to grab history.
- [] Option to grab cookies.
- [] Option to grab credit cards.
- [] Support for Firefox.
- [] Create CLI??
- [] Support for IE?
- [] Support for Safari?
