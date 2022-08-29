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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

// svcPortsCmd represents the svcPorts command
var svcPortsCmd = &cobra.Command{
	Use:   "svcPorts",
	Short: "Show service ports in tabular mode",
	Long: `Show service ports in tabular mode.

	Command example:

	kubectl svcPorts --namespace ingress-controller haproxy-one-kubernetes-ingress haproxy-two-kubernetes-ingress`,
	Run: func(cmd *cobra.Command, args []string) {
		ns, err := cmd.Flags().GetString("namespace")
		if err != nil {
			panic(err.Error())
		}
		var data = [][]string{}
		var header = []string{
			"Service Name",
			"Name",
			"Protocol",
			"Port",
			"TargetPort",
			"NodePort",
			"AppProtocol",
		}
		clientset := ClientSet(genericclioptions.NewConfigFlags(true))

		for _, arg := range args {
			svc, err := clientset.CoreV1().Services(ns).Get(context.TODO(), arg, metav1.GetOptions{})
			if err != nil {
				panic(err.Error())
			}
			DrawSvcTable(clientset, svc.Spec.Ports, ns, header, data, arg)
		}
		svc, err := clientset.CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, svc := range svc.Items {
			DrawSvcTable(clientset, svc.Spec.Ports, ns, header, data, svc.Name)
		}
	},
}

func DrawSvcTable(clientset *kubernetes.Clientset, ports []v1.ServicePort, ns string, header []string, data [][]string, arg string) {
	for _, v := range ports {
		appProtocol := ""
		if v.AppProtocol != nil {
			appProtocol = *v.AppProtocol
		}
		data = append(data,
			[]string{
				"",
				v.Name,
				string(v.Protocol),
				strconv.Itoa(int(v.Port)),
				v.TargetPort.StrVal,
				strconv.Itoa(int(v.NodePort)),
				appProtocol,
			},
		)

	}
	renderTable(header, data, arg)
}

func init() {
	rootCmd.AddCommand(svcPortsCmd)
	svcPortsCmd.PersistentFlags().String("namespace", "default", "service's namespace")
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
