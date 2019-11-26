package config

import (
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Init is a convenience function for populating a struct value with config file, environment variable,
// and/or command-line options. This uses spf13/viper internally to handle this mapping.
// Prefix is the environment variable prefix, config is a value pointer of the config type defined,
// and the optional bind function can be used to define the command-line flags to correspond to
// the config fields. If the "conf" flag is defined, it will be used as the path to a config file.
func Init(prefix string, config interface{}, bind func(*pflag.FlagSet) error) ([]string, error) {
	// Initialize the flags for the config.
	fs := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	// Run bind function from user program to setup flags.
	if bind != nil {
		if err := bind(fs); err != nil {
			return nil, err
		}
	}

	// Parse the flags which default to taking os.Args
	fs.Parse(os.Args[1:])

	// Use viper to merge config sources.
	v := viper.New()

	// Bind flags.
	fs.VisitAll(func(f *pflag.Flag) {
		v.BindPFlag(f.Name, fs.Lookup(f.Name))
	})

	// Use env variables with the provided prefix.
	v.SetEnvPrefix(prefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Read the config file if provided.
	if v.GetString("conf") != "" {
		v.SetConfigFile(v.GetString("conf"))
		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	// Decode resulting options into config.
	if err := v.Unmarshal(config); err != nil {
		return nil, err
	}

	return fs.Args(), nil
}
