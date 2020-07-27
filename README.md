# 🎎 Cameo

<p style="font-size: 18px">WORK IN PROGRESS 🚧</p>

## TODO

1. Tests,
2. Clean and sound documentation,
3. Add community guidelines,
4. Upload docker image to DockerHub.

🦄 Your very own GPG 🔐 encrypted contact form

## ⚙️ Configuration

1. Generate key using `$ gpg --generate-key` without password,
2. Export public key and prepare to share it with a container,
3. Configure SMTP server details.

Default configuration:

```yaml
version: 0.0.1
port: 4000
form_title: A Message To Unicorn
views: /etc/cameo/views

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
    3. Copy views folder and public key to required locations,
    4. Run server

## CLI options

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

## systemd example
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
