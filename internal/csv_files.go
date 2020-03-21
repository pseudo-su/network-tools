package internal

import (
	"bufio"
	"encoding/csv"
	"io"
	"net"
	"strings"
)

func ReadNetworks(reader io.Reader) ([]*net.IPNet, error) {
	res := []*net.IPNet{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if txt == "" {
			continue
		}
		_, pn, err := net.ParseCIDR(txt)
		if err != nil {
			return res, err
		}
		res = append(res, pn)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func WriteCSV(writer io.Writer, outCSV [][]string) error {
	w := csv.NewWriter(writer)
	w.WriteAll(outCSV)

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
