package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Create looks for a config and creates one if it doesnt exist
func Create(path string) {
	if err := viper.SafeWriteConfigAs(path); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfigAs(path)
		}
	}
}

// Load loads configs into the system
func Load() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	viper.AddConfigPath(fmt.Sprintf("%s/%s", home, ".pangolin"))
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
