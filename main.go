package main

import (
	"os"

	"github.com/mdp/qrterminal"
	"github.com/xlzd/gotp"
)

func main() {
	uri := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO").ProvisioningUri("demoAccountName", "issuerName")
	qrterminal.Generate(uri, qrterminal.L, os.Stdout)
}
