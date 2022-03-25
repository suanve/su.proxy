package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"su.proxy/socks5"
)

func init() {
	pflag.String("host", "0.0.0.0", "proxy host")
	pflag.Int("port", 8000, "proxy port")
	pflag.String("user", "admin", "proxy username")
	pflag.String("pass", "123", "proxy password")
	pflag.String("prefix", "", "proxy prefix")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	fmt.Printf("[*] Proxy Listening on %s:%d\n", viper.Get("host"), viper.GetInt("port"))
	if viper.GetString("user") != "" || viper.GetString("pass") != "" {
		fmt.Printf("[*] Auth  %s:%s\n", viper.GetString("user"), viper.GetString("pass"))
	}
	if viper.GetString("prefix") != "" {
		fmt.Printf("[*] Prefix is %s\n", viper.GetString("prefix"))
	}
}

func main() {

	// 代理认证信息
	cred := socks5.StaticCredentials{
		viper.GetString("user"): viper.GetString("pass"),
	}
	cator := socks5.UserPassAuthenticator{Credentials: cred}
	conf := &socks5.Config{AuthMethods: []socks5.Authenticator{cator}}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// 创建一个监听器 监听本地8000端口
	if err := server.ListenAndServe("tcp", "0.0.0.0:8000"); err != nil {
		panic(err)
	}
}
