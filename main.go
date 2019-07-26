package main

import (
	// stdlib
	"bufio"
	"fmt"
	"os"
	"strconv"
	// universe
	"github.com/universelabs/universe-server/http"
	"github.com/universelabs/universe-server/stormdb"
	"github.com/universelabs/universe-server/internal/config"
)

func main() {
	// init config
	consts, cfgerr := config.NewConstants()
	if cfgerr != nil {
		panic(cfgerr)
	}
	// set up services
	db := stormdb.NewClient()
	db.Open(consts.StormDB.Path)
	server := http.NewServer(consts.Host + ":" + strconv.Itoa(consts.Port), db.Keystore())	
	// launch server
	server.Open()
	// hang for server
    buf := bufio.NewReader(os.Stdin)
    fmt.Printf("Press any key to exit...\n")
    _, err := buf.ReadBytes('\n')
    if err != nil {
        panic(err)
    }
}

