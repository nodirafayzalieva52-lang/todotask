package config

// import (
// 	"bufio"
// 	"os"
// 	"strings"
// )

// type Config struct {
// 	DBHost     string
// 	DBPort     string
// 	DBUser     string
// 	DBPassword string
// 	DBName     string
// }

// func New(path string) (*Config, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	cfg := &Config{}
// 	scanner := bufio.NewScanner(file)
	
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.HasPrefix(line, "#") || line == "" {
// 			continue
// 		}
// 		parts := strings.SplitN(line, "=", 2)
// 		if len(parts) != 2 {
// 			continue
// 		}
// 		key := strings.TrimSpace(parts[0])
// 		val := strings.TrimSpace(parts[1])

// 		switch key {
// 		case "DB_HOST": cfg.DBHost = val
// 		case "DB_PORT": cfg.DBPort = val
// 		case "DB_USER": cfg.DBUser = val
// 		case "DB_PASSWORD": cfg.DBPassword = val
// 		case "DB_NAME": cfg.DBName = val
// 		}
// 	}
// 	return cfg, nil
// }