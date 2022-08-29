/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

// svcPortsCmd represents the svcPorts command
var svcPortsCmd = &cobra.Command{
	Use:   "svcPorts",
	Short: "show clearly service ports in tabular mode",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ns, err := cmd.Flags().GetString("namespace")
		if err != nil {
			panic(err.Error())
		}
		var data = [][]string{}
		var header = []string{}
		clientset := ClientSet(genericclioptions.NewConfigFlags(true))
		svc, err := clientset.CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, arg := range args {
			for _, v1 := range svc.Items {
				for _, v2 := range v1.Spec.Ports {
					appProtocol := ""
					if v2.AppProtocol != nil {
						appProtocol = *v2.AppProtocol
					}
					data = append(data,
						[]string{
							"",
							v2.Name,
							string(v2.Protocol),
							strconv.Itoa(int(v2.Port)),
							v2.TargetPort.StrVal,
							strconv.Itoa(int(v2.NodePort)),
							appProtocol,
						},
					)
				}
			}
			header = []string{
				"Service Name",
				"Name",
				"Protocol",
				"Port",
				"TargetPort",
				"NodePort",
				"AppProtocol",
			}
			renderTable(header, data, arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(svcPortsCmd)
	svcPortsCmd.PersistentFlags().String("namespace", "", "service's namespace")
}

// ClientSet k8s clientset
func ClientSet(configFlags *genericclioptions.ConfigFlags) *kubernetes.Clientset {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		panic("kube config load error")
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {

		panic("gen kube config error")
	}
	return clientSet
}

func renderTable(header []string, data [][]string, arg string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Rich([]string{arg}, []tablewriter.Colors{tablewriter.Color(21)})
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.NumLines()
	table.SetHeader(header)
	table.Render() // Send output
}
