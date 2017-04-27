/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file.  As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 *                                                         Guillaume de Sagazan
 * ----------------------------------------------------------------------------
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var start = time.Now()
var last = time.Now()

func getTimeStr(relative, delta bool, format string) string {
	if relative {
		return time.Since(start).String()
	}
	if delta {
		s := time.Since(last).String()
		last = time.Now()

		return s
	}
	return time.Now().Format(format)
}

func main() {
	relative := flag.Bool("relative", false, "log time relative to start")
	delta := flag.Bool("delta", false, "log time relative to last line")
	format := flag.String("format", "Mon Jan 2 15:04:05", "date format (see https://golang.org/pkg/time/#Time.Format)")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Printf("got error %+v\n", err)
			}
			return
		}
		fmt.Printf("%s: %s\n", getTimeStr(*relative, *delta, *format), line)
	}
}
