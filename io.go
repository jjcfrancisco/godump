package main

import (
	"fmt"
	"os"
    "log"
    "bytes"
    "github.com/BurntSushi/toml"
)

type inputConf struct {
    env      string `toml:"env"`
	database string `toml:"database"`
    hostname string `toml:"hostname"`
    port     string `toml:"port"`
	user     string `toml:"user"`
	password string `toml:"password"`
}

type tomlConf struct {
	Database string `toml:"database"`
	Hostname string `toml:"hostname"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
}

var confPath string = fmt.Sprintf("%s/.config/godump.toml", os.Getenv("HOME"))

func configsExist() error {
    
    _, err := os.Stat(confPath)
    if err != nil {
        return err
    }
    return nil
}

func checkEnvExists(e string) bool {

	var data map[string]tomlConf

	_, err := toml.DecodeFile(confPath, &data)
	if err != nil {
		log.Fatal(err)
	}

	for env := range data{
		if env == e {
			return true
		}
	}

	return false

}

func loadConfigs() []*inputConf {

	var data map[string]tomlConf
	var inputConfs []*inputConf


	_, err := toml.DecodeFile(confPath, &data)
	if err != nil {
		log.Fatal(err)
	}

	for env, creds := range data{
		conf := new(inputConf)
		conf.env = env
		conf.database = creds.Database
		conf.hostname = creds.Hostname
		conf.port = creds.Port
		conf.user = creds.User
		conf.password = creds.Password

		inputConfs = append(inputConfs, conf)
	}

	return inputConfs

}

func editEnv(ic *inputConf) {

	confs := loadConfigs()

	for _, conf := range confs {
		if ic.env == conf.env {
			conf.env = ic.env
			conf.database = ic.database
			conf.hostname = ic.hostname
			conf.port = ic.port
			conf.user = ic.user
			conf.password = ic.password
		}
	}

	var buf = new(bytes.Buffer)

	for _, conf := range confs {
		err := toml.NewEncoder(buf).Encode(map[string]interface{}{
			conf.env: map[string]string{
				"database": conf.database,
				"hostname": conf.hostname,
    	        "port": conf.port,
    	        "user": conf.user,
    	        "password": conf.password,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}

    file, err := os.OpenFile(confPath, os.O_WRONLY, 0644)
    if err != nil {
       panic(err)
    }
    defer file.Close()

    // Write the TOML data to the file.
    _, err = file.Write(buf.Bytes())
    if err != nil {
        panic(err)
    }

}

func removeEnv(env string) {

	confs := loadConfigs()
	var new_confs []*inputConf

	for _, conf := range confs {
		if env == conf.env {
			continue
		} else {
			new_confs = append(new_confs, conf)
		}
	}

	var buf = new(bytes.Buffer)

	for _, new_conf := range new_confs {
		err := toml.NewEncoder(buf).Encode(map[string]interface{}{
			new_conf.env: map[string]string{
			"database": new_conf.database,
			"hostname": new_conf.hostname,
    	    "port": new_conf.port,
    	    "user": new_conf.user,
    	    "password": new_conf.password,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}

    file, err := os.OpenFile(confPath, os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
       panic(err)
    }
    defer file.Close()

    // Write the TOML data to the file.
    _, err = file.Write(buf.Bytes())
    if err != nil {
        panic(err)
    }

}

func createEmtpyConfigs() {
    myfile, err := os.Create(confPath)
    if err != nil {
        log.Fatal("Error when creating godump.toml")
    }
    myfile.Close()
}

func saveCredentials(m model) error {

	var buf = new(bytes.Buffer)

	err := toml.NewEncoder(buf).Encode(map[string]interface{}{
		m.creds.env: map[string]string{
			"database": m.creds.database,
			"hostname": m.creds.hostname,
            "port": m.creds.port,
            "user": m.creds.user,
            "password": m.creds.password,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// If file exist A
	// else B
	
    //file, err := os.OpenFile(confPath, os.O_WRONLY, 0644)
	file, err := os.OpenFile(confPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
       panic(err)
    }
    defer file.Close()

    // Write the TOML data to the file.
    _, err = file.Write(buf.Bytes())
    if err != nil {
        panic(err)
    }

	// SORT OUT ERROR
    return fmt.Errorf("")

}

func createDir(d string) error {

	_, err := os.Stat(d)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(d, 0755)
			if err != nil {
				return err
			}
		}
    } 

    return nil
}
