package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() (err error) {
	godotenv.Load()
	return loadConfig()
}

var (
	PrinterHost       = ""
	PrinterSerial     = ""
	PrinterAccessCode = ""
	PrinterName       = ""
	Port              = 8080
)

func loadConfig() error {
	PrinterHost = os.Getenv("BAMBULAB_EXPORTER_PRINTER_HOST")
	PrinterSerial = os.Getenv("BAMBULAB_EXPORTER_PRINTER_SERIAL")
	PrinterAccessCode = os.Getenv("BAMBULAB_EXPORTER_PRINTER_ACCESS_CODE")
	PrinterName = os.Getenv("BAMBULAB_EXPORTER_PRINTER_NAME")

	if PrinterHost == "" {
		return fmt.Errorf("BAMBULAB_EXPORTER_PRINTER_HOST is not set")
	}
	if PrinterSerial == "" {
		return fmt.Errorf("BAMBULAB_EXPORTER_PRINTER_SERIAL is not set")
	}
	if PrinterAccessCode == "" {
		return fmt.Errorf("BAMBULAB_EXPORTER_PRINTER_ACCESS_CODE is not set")
	}
	if PrinterName == "" {
		PrinterName = PrinterSerial
	}

	if portStr := os.Getenv("BAMBULAB_EXPORTER_PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil || port < 1 || port > 65535 {
			return fmt.Errorf("BAMBULAB_EXPORTER_PORT must be a valid port number (1-65535), got %q", portStr)
		}
		Port = port
	}
	return nil
}
