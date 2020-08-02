<p align="center">
    <img width="150" height="150" alt="Cameo" src="https://raw.githubusercontent.com/imanhodjaev/cameo/master/assets/unicorn.png"/>
</p>

<h1 align="center" style="text-align: center">🎎 Cameo</h1>
<h1 align="center" style="text-align: center">🚧 WORK IN PROGRESS 🚧</h1>

[![Build Status](https://travis-ci.com/imanhodjaev/cameo.svg?branch=master)](https://travis-ci.com/imanhodjaev/cameo)

## 💊 TODO

1. Tests,
2. Clean and sound documentation,
3. Upload docker image to DockerHub.

🦄 Your very own GPG 🔐 encrypted contact form

## 🧐 Why?

If you are an investigative journalist or cyber security professional and often need to have secure contact form
then this project aims to fill this gap with easy-to-use and deploy approach, also it might be ideal for people
who want to avoid sharing personal emails and only share a website with contact form.
Contact form uses GPG and encrypts messages before sending them, just generate your key and put
its location in configuration file.

## ⚙️ Configuration

1. Generate key using `$ gpg --generate-key` without password,
2. Export public key and prepare to share it with a container,
3. Configure SMTP server details.

Default configuration:

```yaml
version: 0.0.1
port: 4000
form_title: A Message To Unicorn

mailer:
  host: smtp.google.com
  port: 587
  user: user
  pass: pass
  retries: 3
  send_to: secret@secret.mail
  from_email: contact@myform.com

domains:
  - example.com
  - second-example.com

logs:
  level: INFO

gpg:
  pub_key: /path/to/pub.key
```

## 🔧 Build and run

1. Build docker image and run `make image`,
2. Compile `$ go build -o /cameo/cameo`
    1. Create and export GPG public key, 
    2. Update configuration,
    3. Run server

## 📺 CLI options

```sh
Generate your GPG key and deploy your contact form with Cameo

Usage:
   [flags]

Flags:
      --config string    Config file (default "/etc/cameo/config.yml")
  -h, --help             help for this command
      --port int         Server port (default 4000)
      --pub-key string   GPG public key
```

## 🤔 systemd example
Create `/etc/systemd/system/cameo.service` with the following contents

```unit file (systemd)
[Unit]
Description=PGP encrypted contact form
After=network.target

[Service]
User=root
Group=www-data
ExecStart=/path/to/cameo
WorkingDirectory=/path/to/workdir
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
```

## Screenshot
<img alt="Cameo" src="https://raw.githubusercontent.com/imanhodjaev/cameo/master/assets/screenshot.png"/>


## 💄 Assets

Favicon is taken from https://icons8.com/icon/104324/unicorn
