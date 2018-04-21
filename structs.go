package main

import "time"

type User struct {
	ID        int    `storm:"id"`
	Username  string `storm:"unique"`
	Password  string
	CreatedAt time.Time `storm:"index"`
}

type Config struct {
	// Telegram Bot API token
	Token string `required:"true"`
	// SOCKS5 address, uses to show it user and for applying proxy settings
	Addr string `required:"true"`
	// SOCKS5 port
	Port int `default:"1080"`
	// Private mode
	Private bool `default:"true"`
	// Max limit of accounts
	UsersLimit int `default:"100"`
	// Max limit of connections
	ConnsLimit int `default:"300"`
	// Admin ID to get special features access
	AdminID int
	// Proxy for Bot API
	Proxy struct {
		Addr     string
		Port     int `default:"1080"`
		Username string
		Password string
	}
}
