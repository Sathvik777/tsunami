/*
Copyright © 2019 sathvik.katam

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "sample-http",
  Short: "A test server to load test",
  Long: `A test server to load test and push prometheus metrics`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yaml)")
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".sample-http" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".sample-http")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

