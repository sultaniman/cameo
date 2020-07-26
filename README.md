# 🎎 Cameo

PGP encrypted contact form.

## Configuration

1. Generate key using `$ gpg --generate-key` without password,
2. Export public key and prepare to share it with container,
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

## How to build and run

1. Build docker image and run `make image`,
2. Compile `$ go build -o /cameo/cameo`.

## Using with systemd
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
