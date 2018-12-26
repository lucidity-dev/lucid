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
package builder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/lucidity-dev/lucid/parser"
)

func BuildProject(project parser.Project) {
	fmt.Println("Starting Services: ")
	for _, service := range project.Services {
		startService(service.Src, service.Run)
	}
}

func startService(dir string, run string) {
	run_cmd := strings.Fields(run)
	run = run_cmd[0]
	cmd := exec.Command(run, run_cmd[1:]...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	//err := cmd.Run()
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
}
