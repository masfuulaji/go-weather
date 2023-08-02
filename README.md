
## Overview

Weather Check CLI is a command-line tool built in Go that allows you to check the current weather for a given city using data from the OpenWeatherMap API.

## Prerequisites

- Go (Golang) installed on your machine.
- An API key from OpenWeatherMap. You can sign up for a free API key at https://home.openweathermap.org/users/sign_up.

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/weather-check-cli.git
   ```

2. Navigate to the project directory:

   ```bash
   cd weather-check-cli
   ```

3. Duplicate `env.go.example` then rename it to `env.go` and Replace `"API_KEY"` in the `env.go` file with your actual OpenWeatherMap API key.

4. Build the executable:

   ```bash
   go build -o weather-check
   ```

5. Move the executable to a location in your PATH:

   ```bash
   sudo mv weather-check /usr/local/bin/
   ```

## Usage

The CLI tool can be used as follows:

```bash
weather-check --city "New York"
```

Replace `"New York"` with the desired city name for which you want to check the weather. If you do not provide the `--city` flag, it will default to "Malang".


## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to create an issue or submit a pull request.

## Contact

For questions or support, you can reach out to me at masfuulaji@gmail.com.
