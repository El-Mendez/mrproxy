package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"net"
	"net/url"
	"os"
	"slices"
)

const thePORT string = ":3333"

func main() {
	var target string
	var port string
	if len(os.Args) > 1 {
		target = os.Args[1]
	} else {
		target = "http://localhost:3000"
	}

	if len(os.Args) > 2 {
		port = fmt.Sprintf(":%s", os.Args[2])
	} else {
		port = thePORT
	}

	p := tea.NewProgram(initialModel(GetAddresses(port)), tea.WithAltScreen())
	proxyUrl, err := url.Parse(target)
	if err != nil {
		fmt.Println("URL Error:", err)
		os.Exit(1)
	}
	go setupProxy(port, p, proxyUrl)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func GetAddresses(port string) []string {
	addresses := make([]string, 0)
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPAddr:
				ip = v.IP
				if !ip.IsLoopback() && ip.To4() != nil {
					addresses = append(addresses, fmt.Sprintf("http://%s%s", ip.To4().String(), port))
				}
			case *net.IPNet:
				ip = v.IP
				if !ip.IsLoopback() && ip.To4() != nil {
					addresses = append(addresses, fmt.Sprintf("http://%s%s", ip.To4().String(), port))
				}
			}
		}
	}
	if len(addresses) > 0 {
		addresses = slices.Insert(addresses, 0, fmt.Sprintf("http://localhost%s", port))
	} else {
		addresses = append(addresses, fmt.Sprintf("http://localhost%s", port))
	}
	return addresses
}
