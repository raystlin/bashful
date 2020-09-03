package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/raystlin/bashful/server"
)

const (
	addrFlagName    = "address"
	addrFlagDefault = ":8083"

	certFlagName = "cert"
	keyFlagName  = "key"
)

var ServerCmd = &cli.Command{
	Name:   "server",
	Usage:  "Starts a bashful server",
	Action: serverAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    addrFlagName,
			Value:   addrFlagDefault,
			Aliases: []string{"a"},
			Usage:   "Address and port to listen",
		},
		&cli.StringFlag{
			Name:  certFlagName,
			Usage: "Server certificate path for https (PEM)",
		},
		&cli.StringFlag{
			Name:  keyFlagName,
			Usage: "Server key path for https (PEM)",
		},
		dbFlag,
	},
}

func serverAction(c *cli.Context) error {
	db, err := NewStore(c)
	if err != nil {
		return err
	}
	defer db.Close()

	certPath := c.String(certFlagName)
	keyPath := c.String(keyFlagName)
	addr := c.String(addrFlagName)

	srv := server.NewBashfulServer(db)

	if certPath == "" || keyPath == "" {
		return srv.ListenAndServe(addr)
	} else {
		return srv.ListenAndServeTLS(addr, certPath, keyPath)
	}
}
