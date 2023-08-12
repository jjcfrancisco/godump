package main

import (
	"fmt"
	"os"
    "log"
    "bytes"
    "github.com/BurntSushi/toml"
)

type credentials struct {
    env      string `toml:"env"`
	database string `toml:"database"`
    hostname string `toml:"hostname"`
    port     string `toml:"port"`
	user     string `toml:"user"`
	password string `toml:"password"`
}
//type Config struct {
//	Postgres Postgres `toml:"postgres" comment:"Postgres configuration"`
//}

var confPath string = fmt.Sprintf("%s/.config/godump.toml", os.Getenv("HOME"))

func configsExist() error {
    
    _, err := os.Stat(confPath)
    if err != nil {
        return err
    }
    return nil
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

    //fmt.Println(buf.String())

    return fmt.Errorf("")

}
