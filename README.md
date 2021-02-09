# Notification Server

## About
Simple notification server written in GoLang.

## Usage
- **Email Connection Parameters:**
Email notifications sent using SMTP server based on parameters defined in `resources/connection.properties`.
This file is not versioned as it contains senitive information. Therfore, add it per your own discretion.
- **Non-Sensitive Parameters:** Any non sensitive parameter can be placed under `resources/env.properties`. 
Alternatively it is possible to set-up parameter reading from system environment variables.