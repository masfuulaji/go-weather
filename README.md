
## Overview

Weather Check CLI is a command-line tool built in Go that allows you to check the current weather for a given city using data from the OpenWeatherMap API.

## Prerequisites

- An API key from WeatherAPI. You can sign up for a free API key at https://www.weatherapi.com/signup.aspx.

## Installation

1. Run the App in terminal:

   ```bash 
   ./weather-check 
   ```

2. Set API key with the following command:

   ```bash
    ./weather-check set-key
   ```

5. Move the executable to a location in your PATH:

   ```bash
   sudo mv weather-check /usr/local/bin/
   ```

## Usage

The CLI tool can be used as follows:

## Help for see available commands
```bash
weather-check --help
```

## Examples of usage
```bash
weather-check weather --city "New York"
```
Replace `"New York"` with the desired city name for which you want to check the weather. If you do not provide the `--city` flag, it will default to "Malang".

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to create an issue or submit a pull request.

## Contact

For questions or support, you can reach out to me at masfuulaji@gmail.com.
