package main

import (
    "slices"
    "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)
// FILE PICKER -> https://github.com/charmbracelet/bubbletea/blob/master/examples/file-picker/main.go

type model struct {
    err        error
	feedback   string
	cursor     int
	choice     string
    menu       []string
    title      string
    current    string
    previous   string
    goback     bool
    textInputs []textinput.Model
    creds      credentials
}

// Main menu components
var mmItems = []string{"Dump", "Restore", "Configs"}
var mmTitle = " MAIN menu "
var mmCurrent = "main"

// Configs menu components
var cmItems = []string{"Add db", "Edit db", "Remove db"}
var cmTitle = " CONFIGS menu "
var cmCurrent = "configs"
var cmPrevious = "Main"

// Add db menu components
var adbItems = []string{}
var adbTitle = " ADD DB menu "
var adbCurrent = "add-db"
var adbPrevious = "Configs"
var adbFocused = 0

// Check options menu components
var coItems = []string{"Save", "Cancel", "Test connection"}
var coTitle = " CHECK OPTIONS menu "
var coCurrent = "check-options"
var coPrevious = "Add db"

// Check save menu components
var csSuccessFb = "Credentials succesfully saved."
var csEnvExistsFb = "Environment already exists in configs."

type (
	errMsg error
)

func exceptions(m model) bool {

	if m.current == "add-db" {
		return true
	} else {
		return false
	}

}

func newModel(m model) model {

    var choice string

    if (m.goback) {
        choice = m.previous 
    } else if (m.current == "add-db") {
        choice = "Check Options"
    } else {
        choice = m.menu[m.cursor]
    }

    switch choice {
    case "Main":
        m = model{cursor: 0, menu: mmItems, title: mmTitle, current: mmCurrent}
    
    case "Dump":
    case "Restore":
    case "Configs":
        m = model{cursor: 0, menu: cmItems, title: cmTitle, current: cmCurrent, previous: cmPrevious}
    
    case "Add db":

        inputs := adbTextInputs()
        
        m = model{cursor: 0, menu: adbItems, title: adbTitle, current: adbCurrent,
        previous: adbPrevious, textInputs: inputs}

    case "Check Options":

        userCreds := credentials{env: m.textInputs[env].Value(),
                                 database: m.textInputs[database].Value(),
                                 hostname: m.textInputs[hostname].Value(),
                                 port: m.textInputs[port].Value(),
                                 user: m.textInputs[username].Value(),
                                 password: m.textInputs[password].Value()}

        m = model{cursor: 0, menu: coItems, title: coTitle, current: coCurrent,
        previous: coPrevious, creds: userCreds}

    case "Save":

		// Check if environment already exists, else, create new environment.
		envExists := checkEnvExists(m.creds.env)
		if envExists {
			m.feedback = csEnvExistsFb
		}

        // Check godump.toml exists. If does not exist, it creates one. If
        // similar credentials exists, it feeds an error to m.err.
        err := configsExist()
        if err != nil {
            createEmtpyConfigs()
        }
        err = saveCredentials(m)
        if err != nil {
            m.err = err 
        }

        m = model{cursor: 0, menu: cmItems, title: cmTitle, current: cmCurrent,
		previous: cmPrevious, feedback: csSuccessFb}

    }

    return m

}

func goBack(m model) model {
    if m.current == "main" {
        return m
	} else {
        m.goback = true
        m = newModel(m)
        return m
    }
}

func nextTextarea(m model, msg tea.Msg) model {

    textinputMenus := []string{"add-db"}

    if slices.Contains(textinputMenus, m.current) {
        for i := range m.textInputs {
	    	m.textInputs[i].Blur()
	    }
	    m.textInputs[m.cursor].Focus()
    }

    return m
}
