package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/liqotech/liqo/pkg/liqoctl/factory"
	"github.com/liqotech/liqo/pkg/liqoctl/network"
	"github.com/liqotech/liqo/pkg/liqoctl/output"
)

func main() {
	// Percorsi dei kubeconfig (assumendo che i file siano nella stessa cartella del main.go)
	localKubeconfig := "./rome-remote.yaml"
	remoteKubeconfig := "./milan-remote.yaml"

	// Imposta KUBECONFIG per la factory locale
	os.Setenv("KUBECONFIG", localKubeconfig)
	localFactory := factory.NewForLocal()
	if err := localFactory.Initialize(); err != nil {
		log.Fatalf("Errore nell'inizializzazione della localFactory: %v", err)
	}

	// Imposta KUBECONFIG per la factory remota
	os.Setenv("KUBECONFIG", remoteKubeconfig)
	remoteFactory := factory.NewForRemote()
	if err := remoteFactory.Initialize(); err != nil {
		log.Fatalf("Errore nell'inizializzazione della remoteFactory: %v", err)
	}

	// Non impostiamo esplicitamente il namespace: verrà preso dal kubeconfig.
	// Creiamo le opzioni per il comando network.
	opts := network.NewOptions(localFactory)
	opts.RemoteFactory = remoteFactory

	// Imposta il tipo di Gateway Server a NodePort
	opts.ServerGatewayType = "NodePort"

	// Abilita log dettagliati tramite i printer
	localFactory.Printer = output.NewLocalPrinter(true, true)
	remoteFactory.Printer = output.NewRemotePrinter(true, true)

	// Stampa alcune informazioni per verificare i kubeconfig e il namespace caricato
	fmt.Println("Informazioni localFactory:")
	fmt.Printf("Namespace: %s\n", localFactory.Namespace)
	fmt.Printf("RESTConfig: %+v\n\n", localFactory.RESTConfig)

	fmt.Println("Informazioni remoteFactory:")
	fmt.Printf("Namespace: %s\n", remoteFactory.Namespace)
	fmt.Printf("RESTConfig: %+v\n\n", remoteFactory.RESTConfig)

	// Crea un contesto con timeout (ad esempio, 10 minuti)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Creazione dei cluster
	fmt.Println("Creazione del cluster locale (cluster1)...")
	cluster1, err := network.NewCluster(ctx, localFactory, remoteFactory, true)
	if err != nil {
		log.Fatalf("Errore nella creazione di cluster1: %v", err)
	}
	fmt.Printf("Cluster1 creato: %+v\n\n", cluster1)

	fmt.Println("Creazione del cluster remoto (cluster2)...")
	cluster2, err := network.NewCluster(ctx, remoteFactory, localFactory, true)
	if err != nil {
		log.Fatalf("Errore nella creazione di cluster2: %v", err)
	}
	fmt.Printf("Cluster2 creato: %+v\n\n", cluster2)

	// In questo esempio, non chiamiamo esplicitamente initNetworkConfigs
	// perché questa operazione viene gestita internamente in RunConnect.

	// Esecuzione del comando "network connect"
	fmt.Println("Esecuzione del comando 'network connect'...")
	if err := opts.RunConnect(ctx); err != nil {
		log.Fatalf("Errore durante l'esecuzione di 'network connect': %v", err)
	}
	fmt.Println("Operazione 'network connect' completata con successo.")
}

