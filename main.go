package main

import (
	"flag"
	hook "github.com/afarid/k8s-webhook-server/pkg"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func main() {
	var certDir string
	var port int
	var debug bool
	flag.StringVar(&certDir, "cert-dir", "./certs", "The directory where the TLS certs are located")
	flag.IntVar(&port, "port", 8443, "The port where the server will listen")
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")
	flag.Parse()

	// Create a new k8s controller manager
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		log.Fatal("unable to start manager", err)
		return
	}
	// Create a new webhook instance
	log.Println("setting up new webhook server")
	hookServer := mgr.GetWebhookServer()
	hookServer.CertDir = certDir
	hookServer.Port = port

	log.Println("registering webhooks to the webhook server")
	hookServer.Register("/validate", &webhook.Admission{Handler: &hook.PodValidator{Client: mgr.GetClient(), Debug: debug}})

	log.Println("starting manager")
	if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Fatal(err, "unable to run manager")
	}

}
