# GopplicationEntry
![](./.README/header.png)

<p align="center">XDG Desktop Entry parser written golang</p>

<p align="center">
    <a href="https://github.com/free-bots/GopplicationEnty/releases" alt="GitHub release (latest by date)">
        <img alt="GitHub" src="https://img.shields.io/github/v/release/free-bots/GopplicationEnty?style=flat-square"></a>
    <a href="https://github.com/free-bots/GopplicationEnty/blob/master/LICENSE.md" alt="License">
        <img alt="GitHub" src="https://img.shields.io/github/license/free-bots/GopplicationEnty?style=flat-square"></a>
    <a href="https://github.com/free-bots/GopplicationEnty/graphs/contributors" alt="Contributors">
        <img alt="GitHub" src="https://img.shields.io/github/contributors/free-bots/GopplicationEnty?style=flat-square"></a>
</p>

# Table of Contents

- [GopplicationEntry](#gopplicationentry)
- [Table of Contents](#table-of-contents)
  - [API](#api)
    - [Parse a single file](#parse-a-single-file)
    - [Get all entries of your system](#get-all-entries-of-your-system)
    - [Get all locations for entries](#get-all-locations-for-entries)
  - [Installation](#installation)
  - [Contribution](#contribution)
  - [License](#license)

## API
### Parse a single file
```go
for _, name := range fileNames {
    application := Parse("path to the .desktop file", true)  // set to true if the field codes should be trimmed see: https://specifications.freedesktop.org/desktop-entry-spec/desktop-entry-spec-latest.html#exec-variables  
    if application == nil {
        continue
    }

    applications = append(applications, application)
}
```
### Get all entries of your system
```go
package repositories

import (
	GoppilcationEntry "github.com/free-bots/GopplicationEntry"
)

func FindAll() []*GoppilcationEntry.ApplicationEntry {
	return GoppilcationEntry.FindAllEntries()
}
```
### Get all locations for entries
```go
/* returns an array like 
    [
        "/home/free-bots/.local/share/flatpak/exports/share",
        "/var/lib/flatpak/exports/share",
        "/usr/local/share/",
        "/usr/share/"
    ]
 */
for _, path := range GetApplicationPaths() {

    path = filepath.Join(path, ApplicationDirName)

    fileInfo, err := os.Stat(path)

    if err != nil {
        fmt.Println(err)
        continue
    }

...
```

## Installation
```
go get https://github.com/free-bots/GopplicationEntry
```

## Contribution

If you miss an important feature fell free to contribute or create a feature request.


## License
> MIT License

> Copyright (c) 2021 free-bots

> Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
