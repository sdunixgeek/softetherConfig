package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"softetherConfig/softetherApi"
	"strconv"
)

var (
	secHost = "127.0.0.1"
	secPort = 443
	secPass = ""
	secHub  = "DEFAULT"
)

func main() {
	flag.StringVar(&secHost, "s", LookupEnvOrString("SEC_HOST", secHost), "SoftEther Hostname")
	flag.IntVar(&secPort, "p", LookupEnvOrInt("SEC_PORT", secPort), "SoftEther Port number")
	flag.StringVar(&secPass, "P", LookupEnvOrString("SEC_PASS", secPass), "SoftEther Admin Password")
	flag.StringVar(&secHub, "H", LookupEnvOrString("SEC_HUB", secHub), "SoftEther Hub Name")
	flag.Parse()
	requiredFlags := []string{"P"}
	CheckForRequiredFlags(flag.CommandLine, requiredFlags)
	log.Println("softetherConfig.status=starting")
	// log.Printf("app.config %v\n", getConfig(flag.CommandLine))
	api := softetherApi.NewAPI(secHost, secPort, secPass)
	err := api.HandShake()
	apiTest, err := api.Test()
	check(err, "api.Test")
	log.Println(apiTest)
	api.Disconnect()
}

// LookupEnvOrString Lookup environment variable or set to default
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// LookupEnvOrInt Lookup Environment variable Int type or set default
func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Panicf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}

// CheckForRequiredFlags Verify all required flags are set
func CheckForRequiredFlags(fs *flag.FlagSet, required []string) {
	isFlagSet := make(map[string]bool)
	// Checks all values including Environment variable defaults
	fs.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != "" {
			isFlagSet[f.Name] = true
		}
	})
	for _, req := range required {
		if !isFlagSet[req] {
			fs.Usage()
			log.Fatalf("missing required -%s argument/flag", req)
		}
	}
}

// getConfig grab the configuration from a file
func getConfig(fs *flag.FlagSet) []string {
	cfg := make([]string, 0, 10)
	fs.VisitAll(func(f *flag.Flag) {
		cfg = append(cfg, fmt.Sprintf("%s:%q", f.Name, f.Value.String()))
	})

	return cfg
}

// check for errors
func check(e error, desc string) {
	if e != nil {
		if len(desc) > 0 {
			log.Panicf("%s: %v", desc, e)
		} else {
			log.Panicf("%s: %v", "unknownCall", e)
		}
	}
}
