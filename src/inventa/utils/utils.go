package utils

import (
	"log"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Conf holds configuration information
type Conf struct {
	RunTimeMode           string   `yaml:"run_time_mode"`
	LocalJSONFile         string   `yaml:"local_json_file"`
	LocalRouterID         string   `yaml:"local_router_id"`
	LocalASN              int      `yaml:"local_asn"`
	PeerIPv4Address       string   `yaml:"peer_ipv4_address"`
	PeerASN               int      `yaml:"peer_asn"`
	HTTPListenPort        int      `yaml:"http_listen_port"`
	NodeNameStripPatterns []string `yaml:"node_name_strip_patterns"`
	GroupSplitChar        string   `yaml:"group_split_char"`
	GroupSplitIndex       int      `yaml:"group_split_index"`
}

// Log is the logging object
var Log logrus.Logger

// Configs is our shared config object
var Configs Conf

// FindInArray finds an element in an array
func FindInArray(what string, where []string) (idx int, found bool) {
	for i, v := range where {
		if v == what {
			return i, true
		}
	}
	return 0, false
}

// InitConfig initializes the configuration object
func InitConfig() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, &Configs)
	if err != nil {
		log.Fatalln(err)
	}
}

// InitLogger initializes the logger
func InitLogger() {
	Log := logrus.New()
	Log.SetLevel(logrus.InfoLevel)
}

// StripUnwanted removes any substrings from our name string
func StripUnwanted(name string) string {
	for _, pattern := range Configs.NodeNameStripPatterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			log.Fatal(err)
		}
		name = re.ReplaceAllString(name, "")
	}
	return name
}
