**WARNING: This is not production ready!**

This is a Discord implementation of https://github.com/l3lackShark/binance-api-listener

# Docs

* .env file in the root folder of the executable is used to load configuration.
Tokens:
  - `BOT_TOKEN` = Your Discord bot token
  - `BASE_URL` = Base http URL of the running instance of https://github.com/l3lackShark/binance-api-listener (default: `http://localhost:24080"`)
  


# Usage

* Add a bot to your server and call `!btc` to get avg BTC price in the last 5 minutes