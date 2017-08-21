package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/lePereT/gowebapp/daemon"
)

var assetsPath string

func processFlags() *daemon.Config {
	cfg := &daemon.Config{} // cfg is a pointer to a daemon.Config struct as defined in the other package


	/* these first two statements take the command line arguments and 
	associate them to pre-defined variables in daemon.Config. The third 
	associates a flag with the assetsPath declared earlier. In all cases   
	need to pass pointers to the flag declaration function (mutates) */
	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&cfg.Db.ConnectString, "db-connect", "host=/var/run/postgresql dbname=gowebapp sslmode=disable", "DB Connect String")
	flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")

	flag.Parse()
	return cfg
}

func setupHttpAssets(cfg *daemon.Config) {
	log.Printf("Assets served from %q.", assetsPath)
	cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
	cfg := processFlags()

	setupHttpAssets(cfg)

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}
