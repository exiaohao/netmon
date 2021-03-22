package main

import (
	"github.com/exiaohao/netmon/pkg/netmon/server"
	"github.com/exiaohao/netmon/pkg/agent"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var (
	listenAddr	string
	agentID		string
	rootCmd = &cobra.Command{
		Use: "notmon",
		Short: "NetMon",
		Long: "NetMon daemon",
		SilenceUsage: true,
	}
	serverCmd = &cobra.Command{
		Use: "server",
		Short: "Run netmon server",
		Long: "Run API server",
		RunE: func(*cobra.Command, []string) error {
			s := new(server.Server)
			s.ListenAddr = listenAddr
			stopCh := setupSignalHandler()
			s.Run(stopCh)
			return nil
		},
	}
	agentCmd = &cobra.Command{
		Use: "agent",
		Short: "Run netmon agent",
		Long: "Run netmon agent",
		RunE: func(*cobra.Command, []string) error {
			a := new(agent.Agent)
			a.AgentID = agentID
			stopCh := setupSignalHandler()
			a.Run(stopCh)
			return nil
		},
	}
)

func init() {
	serverCmd.PersistentFlags().StringVar(&listenAddr, "address", "0.0.0.0:8989", "netmon server API")
	agentCmd.PersistentFlags().StringVar(&agentID, "agentid", "", "netmon agent ID")
}

func main() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(agentCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func setupSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		close(stop)
		<-sigs
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}