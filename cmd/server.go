package cmd

import (
	"github.com/spf13/cobra"
	"grpc-sai-test/server"

	log "grpc-sai-test/store/log"
	_ "grpc-sai-test/store/mysql"
	_ "grpc-sai-test/store/redis"
)


var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Logger.Println("Recover error : %v", err)
			}
		}()

		server.Run()
	},
}

func init() {


	//生成证书时的hostname需要满足hostname的格式，改为grpc.sai

	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50052", "server port")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./conf/certs/server.pem", "cert pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./conf/certs/server.key", "cert key path")
	serverCmd.Flags().StringVarP(&server.CertServerName, "cert-name", "", "grpc.sai", "server's hostname")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "proto", "path to the directory which contains swagger definitions")


	rootCmd.AddCommand(serverCmd)
}
