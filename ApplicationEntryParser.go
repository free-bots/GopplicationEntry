package GoppilcationEntry

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const debugLog = false

// parser for .desktop files
func Parse(path string, removeExecFields bool) *ApplicationEntry {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var applicationEntry *ApplicationEntry
	var actions = make([]*Action, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case IsLineEmpty(line):
			continue
		case IsLineGroup(line):
			applicationEntry, actions = parseGroup(scanner, removeExecFields, line, nil, actions)
			continue
		case IsLineComment(line):
			continue
		default:
			if debugLog {
				fmt.Printf("line not matched %s\n", line)
			}
		}

		if debugLog {
			fmt.Println(SplitLineToKeyValue(scanner.Text()))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if applicationEntry != nil {
		applicationEntry.Actions = actions

		splitFileName := strings.Split(file.Name(), "/")
		length := len(splitFileName)
		applicationEntry.Id = splitFileName[length-1]
	}

	return applicationEntry
}

func isLineDesktopEntry(line string) bool {
	return line == "[Desktop Entry]"
}

func isLineDesktopAction(line string) bool {
	matched, err := regexp.MatchString("^\\[Desktop Action.*]", line)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return matched
}

func parseEntry(scanner *bufio.Scanner, removeExecFields bool) (*ApplicationEntry, *string) {

	entry := new(ApplicationEntry)

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case IsLineEmpty(line):
			continue
		case IsLineComment(line):
			continue
		case IsLineGroup(line):
			return entry, &line
		default:
			OnKey(line, "Version", func(key string, value string) {
				float, err := strconv.ParseFloat(value, 32)
				if err != nil {
					fmt.Println(err)
					return
				}

				entry.Version = float32(float)
			})
			OnKey(line, "Type", func(key string, value string) {
				entry.Type = value
			})
			OnKey(line, "Name", func(key string, value string) {
				entry.Name = value
			})
			OnKey(line, "Comment", func(key string, value string) {
				entry.Comment = value
			})
			OnKey(line, "TryExec", func(key string, value string) {
				entry.TryExec = value
			})
			OnKey(line, "Name", func(key string, value string) {
				entry.Name = value
			})
			OnKey(line, "Exec", func(key string, value string) {
				if removeExecFields {
					entry.Exec = removeExecFieldCodes(value)
				} else {
					entry.Exec = value
				}
			})
			OnKey(line, "Icon", func(key string, value string) {
				entry.Icon = value
			})
			OnKey(line, "MimeType", func(key string, value string) {
				if IsSemicolonList(value) {
					entry.MimeType = GetSemicolonList(value)
				}
			})
			OnKey(line, "Actions", func(key string, value string) {
			})
			OnKey(line, "Categories", func(key string, value string) {
				if IsSemicolonList(value) {
					entry.Categories = GetSemicolonList(value)
				}
			})
		}
	}

	return entry, nil
}

func parseAction(scanner *bufio.Scanner, removeExecFields bool) (*Action, *string) {

	var action = new(Action)

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case IsLineEmpty(line):
			continue
		case IsLineComment(line):
			continue
		case IsLineGroup(line):
			return action, &line
		default:
			OnKey(line, "Name", func(key string, value string) {
				action.Name = value
			})
			OnKey(line, "Exec", func(key string, value string) {
				if removeExecFields {
					action.Exec = removeExecFieldCodes(value)
				} else {
					action.Exec = value
				}
			})
			OnKey(line, "Icon", func(key string, value string) {
				action.Icon = value
			})
		}
	}

	return action, nil
}

func parseGroup(scanner *bufio.Scanner, removeExecFields bool, line string, applicationEntry *ApplicationEntry, actions []*Action) (*ApplicationEntry, []*Action) {
	if isLineDesktopEntry(line) {
		var stopLine *string
		applicationEntry, stopLine = parseEntry(scanner, removeExecFields)
		if stopLine != nil {
			return parseGroup(scanner, removeExecFields, *stopLine, applicationEntry, actions)
		}
	}

	if isLineDesktopAction(line) {
		var newAction, stopLine = parseAction(scanner, removeExecFields)
		if newAction != nil {
			actions = append(actions, newAction)
		}
		if stopLine != nil {
			return parseGroup(scanner, removeExecFields, *stopLine, applicationEntry, actions)
		}
	}

	return applicationEntry, actions
}

func removeExecFieldCodes(value string) string {
	return strings.NewReplacer(
		"%f", "",
		"%F", "",
		"%u", "",
		"%U", "",
		"%d", "",
		"%D", "",
		"%n", "",
		"%N", "",
		"%i", "",
		"%c", "",
		"%k", "",
		"%v", "",
		"%m", "").Replace(value)
}
