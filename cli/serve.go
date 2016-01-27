package cli

import (
	"github.com/bbuck/dragon-mud/data"
	"github.com/bbuck/dragon-mud/logger"
	"github.com/bbuck/dragon-mud/random"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the DragonMUD server.",
		Long: `Starts the DragonMUD Game server to listen for new player connections.
All lifecycle scripts will be notified during boot and the configuration
information will be processed.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Log().Infof("A %s dragon arrives to serve you today.", getDragonColor())
			logger.Log().WithField("env", viper.GetString("env")).Info("Configuration loaded")

			_, err := data.DefaultFactory.Open()
			if err != nil {
				logger.Log().WithField("error", err.Error()).Fatal("Failed to establish connection with the database.")
			}
			logger.Log().Info("Successfully established connection with database")
		},
	}

	dragonColors = []string{
		"{black+h:white+h}black{reset}",
		"{red+h}red{reset}",
		"{green+h}green{reset}",
		"{yellow+h}gold{reset}",
		"{blue+h}blue{reset}",
		"{white+h}white{reset}",
		"{white+u}silver{reset}",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

func getDragonColor() string {
	index := random.Intn(len(dragonColors))

	return dragonColors[index]
}
