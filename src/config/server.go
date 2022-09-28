package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/zazhedho/gorental/src/routers"
)

var ServeCmd = &cobra.Command{
	Use:   "server",
	Short: "start api server",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {

		var addrs string = "0.0.0.0:8080"
		if port := os.Getenv("PORT"); port != "" {
			addrs = ":" + port
		}

		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      mainRoute,
		}

		fmt.Println("App running on http://"+addrs, "success")
		srv.ListenAndServe()
		return nil

	} else {
		return err
	}
}
