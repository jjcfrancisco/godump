package main

import (
    "log"
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
	trigger    bool
	env		   string
    textInputs []textinput.Model
    search     textinput.Model
    creds      inputConf
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

// Dump menu components
var dmItems = []string{"All database", "Single schema"}
var dmTitle = " DUMP menu "
var dmCurrent = "dump"
var dmPrevious = "Main"

// Dump all menu components
var daItems = []string{}
var daTitle = " TYPE FILENAME menu "
var daCurrent = "type-filename"
var daPrevious = "All database"
var daFocused = 0

// List dbs menu components
var ldbsItems = []string{}
var ldbsTitle = " LIST DBS menu "
var ldbsCurrent = "list-dbs"
var ldbsPrevious = "Dump"
var ldbsFocused = 0

// Add db menu components
var adbItems = []string{}
var adbTitle = " ADD DB menu "
var adbCurrent = "add-db"
var adbPrevious = "Configs"
var adbFocused = 0

// db to edit menu components
var db2eTitle = " DB TO EDIT menu "
var db2eCurrent = "db-to-edit"
var db2ePrevious = "Configs"
var db2eFocused = 0

// db to remove menu components
var db2rTitle = " DB TO REMOVE menu "
var db2rCurrent = "db-to-remove"
var db2rPrevious = "Configs"
var db2rFocused = 0

// Edit single db menu components
var esdbItems = []string{}
var esdbTitle = " EDIT db menu "
var esdbCurrent = "edit-single-db"
var esdbPrevious = "Edit db"
var esdbFocused = 0

// Check options menu components
var coItems = []string{"Save", "Cancel", "Test connection"}
var coTitle = " CHECK OPTIONS menu "
var coCurrent = "check-options"

// Check save menu components
var savedFb = "Credentials succesfully saved."
var existsFb = "Environment already exists in configs."
var deletedFb = "Credentials successfully deleted."
var pingSuccessFb = "Connection established ✓"
var pingFailFb = "Connection failed ✗"

type (
	errMsg error
)

func exceptions(m model) bool {

	if m.current == "add-db" || m.current == "edit-single-db" || m.current == "type-filename" {
		return true
	} else {
		return false
	}

}

func newModel(m model) model {

    var choice string
	var previous string

    if m.goback {
        choice = m.previous 
    } else if m.current == "add-db" {
        choice = "Check Options"
		previous = "Add db"
    } else if m.current == "edit-single-db" {
		choice = "Check Options"
		previous = "Edit single db"
	} else if m.current == "db-to-edit" { 
		choice = "Edit single db"
	} else if m.current == "db-to-remove" {
		choice = "Remove single db"
	} else if m.current == "list-dbs" {
        choice = "Type filename"
    } else if m.current == "type-filename" {
        choice = "Dump all"
    } else {
        choice = m.menu[m.cursor]
    }

    switch choice {
    case "Main":
        m = model{cursor: 0, menu: mmItems, title: mmTitle, current: mmCurrent}
    
    case "Dump":
        m = model{cursor: 0, menu: dmItems, title: dmTitle, current: dmCurrent, previous: dmPrevious}

    case "All database":

        var dbs []string
        confs := loadConfigs()

        for _, db := range confs {
            dbs = append(dbs, db.env)
        }

        m = model{cursor: 0, menu: dbs, title: ldbsTitle, current: ldbsCurrent,
        previous: ldbsPrevious, env: dbs[m.cursor]}


    case "Type filename":

        inputs := inputSearch()

        m = model{cursor: 0, menu: daItems, title: daTitle, current: daCurrent,
        previous: daPrevious, search: inputs, env: m.env}

    case "Dump all":

        var creds inputConf
        confs := loadConfigs()

        for _, conf := range confs {
            if m.env == conf.env {
                creds = *conf
            }
        }
        
        m.creds = creds
        err := dump("all", m)
        if err != nil {
            log.Fatal(err)
        }

    case "Restore":

    case "Configs":
        m = model{cursor: 0, menu: cmItems, title: cmTitle, current: cmCurrent, previous: cmPrevious}
    
    case "Add db":

		var inputs []textinput.Model

		if !m.trigger {
			inputs = adbTextInputs()
		} else {
			inputs = m.textInputs
		}

        
        m = model{cursor: 0, menu: adbItems, title: adbTitle, current: adbCurrent,
        previous: adbPrevious, textInputs: inputs}

	case "Edit db":

		var db2eItems []string
		confs := loadConfigs()

		for _, conf := range confs {
			db2eItems = append(db2eItems, conf.env)	
		}

		m = model{cursor: 0, menu: db2eItems, title: db2eTitle, current: db2eCurrent,
		previous: db2ePrevious}

	case "Edit single db":

		var inputs []textinput.Model

		ic := &inputConf{}
		confs := loadConfigs()
		for _, conf := range confs {
			env := m.menu[m.cursor]
			if env == conf.env {
				ic.env = conf.env
				ic.database = conf.database
				ic.hostname = conf.hostname
				ic.port = conf.port
				ic.user = conf.user
				ic.password = conf.password
			}
		}

		if !m.trigger {
			inputs = edbTextInputs(ic)
		} else {
			inputs = m.textInputs
		}

		m = model{cursor: 0, menu: esdbItems, title: esdbTitle, current: esdbCurrent,
		previous: esdbPrevious, textInputs: inputs}

	case "Remove db":

		var db2rItems []string
		confs := loadConfigs()

		for _, conf := range confs {
			db2rItems = append(db2rItems, conf.env)	
		}

		m = model{cursor: 0, menu: db2rItems, title: db2rTitle, current: db2rCurrent,
		previous: db2rPrevious}

	case "Remove single db":
		env := m.menu[m.cursor]
		removeEnv(env)

        m = model{cursor: 0, menu: cmItems, title: cmTitle, current: cmCurrent,
		previous: cmPrevious, feedback: deletedFb}

    case "Check Options":

        userCreds := inputConf{env: m.textInputs[env].Value(),
                               database: m.textInputs[database].Value(),
                               hostname: m.textInputs[hostname].Value(),
                               port: m.textInputs[port].Value(),
                               user: m.textInputs[username].Value(),
                               password: m.textInputs[password].Value()}

        m = model{cursor: 0, menu: coItems, title: coTitle, current: coCurrent,
		previous: previous, creds: userCreds, textInputs: m.textInputs, trigger: true}

    case "Save":

		if m.previous == "Add db" {
			// Check if environment already exists, else, create new environment.
			envExists := checkEnvExists(m.creds.env)
			if envExists {
				m.feedback = existsFb
			} else {
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
				previous: cmPrevious, feedback: savedFb}
			}
		} else if m.previous == "Edit single db" {
			editEnv(&m.creds)
			m.feedback = savedFb 
		}
	case "Cancel":

		m = model{cursor: 0, menu: cmItems, title: cmTitle, current: cmCurrent,
		previous: cmPrevious}

	case "Test connection":
        _, err := NewConn(&m.creds) 
        if err != nil {
            m.feedback = pingFailFb
        } else {
            m.feedback = pingSuccessFb
        }

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

    textinputMenus := []string{"add-db", "edit-single-db"}

    if slices.Contains(textinputMenus, m.current) {
        for i := range m.textInputs {
	    	m.textInputs[i].Blur()
	    }
	    m.textInputs[m.cursor].Focus()
    }

    return m
}
