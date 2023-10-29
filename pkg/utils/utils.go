package utils

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"regexp"
	"strings"
)

func ListSrvAddr(port int) string {
	return NewAddress("0.0.0.0", port)
}

func NewAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func CloseConn(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatalf("error closing the connection (%s)", conn.Target())
	}
}

func FormatErrors(errs []string) error {
	if errs == nil || len(errs) == 0 {
		return errors.New("{  }")
	}
	return fmt.Errorf("{\n\t%s\n}", strings.Join(errs, "\n\t"))
}

func FormatSliceStrings(slc []string) string {
	return fmt.Sprintf("[ %s ]", strings.Join(slc, ", "))
}

func MatchRegex(regex, s string) bool {
	r, err := regexp.Compile(regex)
	if err != nil {
		return false
	}
	matched := r.MatchString(s)
	return matched
}
