package conf

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Port int64
	//Config string
	//LogConfig LoggingConfig
}

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("NETLIFY")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/.example")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	fmt.Printf("port: %d\n", viper.GetInt("port"))

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
