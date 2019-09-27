/*
Copyright © 2019 Matias Charriere <m.charriere@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"

	"github.com/mcharriere/mongocheck/pkg/config"
	"github.com/mcharriere/mongocheck/pkg/conn"
)

var Uri string
var Wait bool
var Tries int
var Interval int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mongocheck",
	Short: "Check if MongoDB server is UP",
	Long: `Check if MongoDB server is accepting connections.
Use --host to set the target.
Use --wait to wait for the connection to be establised.
`,
	Run: func(cmd *cobra.Command, args []string) {

		cfg := Config.New()
		cfg.SetUri(Uri)

		cnx := Conn.New(cfg)

		if Wait {
			for i := 0; i <= Tries; i++ {
				err := cnx.Check()
				if err == nil {
					break
				}
				fmt.Printf("%d # Connetion not establised.\n", i)
				time.Sleep(time.Duration(Interval) * time.Second)
			}
		} else {
			err := cnx.Check()
			if err != nil {
				fmt.Println("Connetion not establised")
				os.Exit(1)
			}
		}
		fmt.Println("Connetion establised.")
		os.Exit(0)
	},
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

	rootCmd.PersistentFlags().StringVarP(&Uri, "uri", "u", "mongodb://localhost:27017", "url to connect to")
	rootCmd.PersistentFlags().BoolVarP(&Wait, "wait", "w", false, "wait to establish connection")
	rootCmd.PersistentFlags().IntVarP(&Interval, "interval", "i", 10, "interval between checks")
	rootCmd.PersistentFlags().IntVarP(&Tries, "tries", "n", 10, "max tries")
}
