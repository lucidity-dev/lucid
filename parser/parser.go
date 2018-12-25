// Copyright © 2018 Rishi Desai
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package parser

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Project struct {
	Services []Service `yaml:"services"`
}

type Service struct {
	Name        string   `yaml:"name"`
	Src         string   `yaml:"src"`
	Publishers  []string `yaml:"publishers"`
	Subscribers []string `yaml:"subscribers"`
	Build       string   `yaml:"build"`
	Run         string   `yaml:"run"`
}

func ParseConfig(data []byte, file_type string) {
	if file_type == "yaml" {
		parseYaml(data)
	}
}

func parseYaml(data []byte) {
	fmt.Println(string(data))
	p := Project{}
	err := yaml.UnmarshalStrict(data, &p)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
		os.Exit(-1)
	}

	// TODO: Return project
	for _, s := range p.Services {
		fmt.Printf("service: \n %v \n", s)
	}
}
