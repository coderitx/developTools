package tools


import (
    "fmt"
    "net"
)

// GetIp 获取ip地址
func GetIp() (string,error) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return "",err
    }
    for _, addr := range addrs {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
          return ipnet.IP.String(),nil 
        }
    }
}
