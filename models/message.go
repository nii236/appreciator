package models

import "strings"

// Message contains information about the message you wish to send
type Message struct {
	NamesList      []string
	Name           string
	Attended       bool
	Gift           string
	AdditionalNote string
	From           string
}

func (m Message) Save() {
}

func (m *Message) Process() error {
	m.NamesList = titleize(m.NamesList)
	m.Name = toSentence(m.NamesList)
	m.Gift = strings.ToLower(m.Gift)
	return nil
}

func toSentence(namesList []string) string {
	switch len(namesList) {
	case 1:
		return namesList[0]
	case 2:
		return namesList[0] + " and " + namesList[1]
	default:
		temp := strings.Join(namesList[:len(namesList)-1], ", ")
		finalName := namesList[len(namesList)-1]
		return strings.Join([]string{temp, finalName}, " and ")
	}
}

func titleize(namesList []string) []string {

	for i, name := range namesList {
		name = strings.Title(name)
		namesList[i] = name
	}
	return namesList
}
