# Bambulab Exporter

A prometheus exporter for Bambu Lab 3D printers, pulling metrics directly from the printer using MQTT. This works for both "cloud" and "LAN Only" modes. Tested with the P1S and one AMS, but this likely works for others configurations as well.

A sample of the provided metrics can be found in [metrics-sample.txt](./metrics-sample.txt). (note: this sample is not guaranteed to be 100% up to date at all times)

### Setup

Things you need:

- The printer IP address/host. It's recommended to configure a static IP for your printer
- The printer [serial number](https://wiki.bambulab.com/en/general/find-sn)
- The printer Access Code. On P1S this can be found on the printer in the WLAN menu, below the current IP address

### Environment variables

Configuration is done via environment variables (or a `.env` file when running the binary directly). All variables use the `BAMBULAB_EXPORTER_` prefix.

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `BAMBULAB_EXPORTER_PRINTER_HOST` | Yes | — | Printer IP address or hostname. Use a static IP for reliable connectivity. |
| `BAMBULAB_EXPORTER_PRINTER_SERIAL` | Yes | — | Printer serial number. See [Bambu Wiki](https://wiki.bambulab.com/en/general/find-sn) for how to find it. |
| `BAMBULAB_EXPORTER_PRINTER_ACCESS_CODE` | Yes | — | Printer access code. On P1S: WLAN menu on the printer, shown below the current IP address. |
| `BAMBULAB_EXPORTER_PRINTER_NAME` | No | (serial number) | Friendly name for the printer. Used as a label in metrics; useful when running multiple exporters. |
| `BAMBULAB_EXPORTER_PORT` | No | `8080` | Port the HTTP server binds to for the `/metrics` endpoint. Must be between 1 and 65535. |

### Port

The exporter serves Prometheus metrics on the `/metrics` path. The listen port is controlled by `BAMBULAB_EXPORTER_PORT` (default: `8080`). To use a different port:

- **Docker**: set `BAMBULAB_EXPORTER_PORT` in `environment` and map it in `ports`, e.g. `"9090:9090"` with `BAMBULAB_EXPORTER_PORT: "9090"`.
- **Binary**: set the env var before running, e.g. `BAMBULAB_EXPORTER_PORT=9090 ./bambulab-exporter`.

### Running with Docker

1. Copy the sample compose file: `cp docker-compose.sample.yml docker-compose.yml`
2. Edit `docker-compose.yml` and set the required environment variables (see table above).
3. Optionally set `BAMBULAB_EXPORTER_PORT` and adjust the `ports` mapping if you don’t want to use 8080.
4. Start the service: `docker compose up -d`

### Running from source

Advanced users who do not wish to use Docker can build the binary and run it with the same environment variables, or use a `.env` file in the current working directory (loaded automatically via [godotenv](https://github.com/joho/godotenv)).

```bash
go build -o bambulab-exporter .
./bambulab-exporter
```
