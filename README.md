Compile
=======

```bash
sudo docker run --rm \
	-v $PWD:/usr/src/oauth \
	-v $GOPATH:/go \
	-w /usr/src/oauth \
	golang:1.4 \
	go get -d -v && go build -v
```

Build Docker Container
======================

```bash
sudo docker build -t vfarcic/oauth .
```

Run
===

Google authentication data can be created in [Google Developers Console](https://console.developers.google.com).

Following environment variables can be used:

* HOST: Host where the application is running. Defaults to ":8080"
* SEC_KEY: Security key. This is mandatory variable.
* REDIRECT_URL: URL where users will be redirected to after authenticating. This is mandatory variable.
* GOOGLE_CLIENT_ID: Google client ID. This is mandatory variable if google authentication is used.
* GOOGLE_SECRET: Google client secret. This is mandatory variable if google authentication is used.
* GOOGLE_REDIRECT_URL: Google redirect URI. This is mandatory variable if google authentication is used.

Example Run
-----------

```bash
sudo docker run -d --name oauth \
	-p 8080:8080 \
	--link oauth-db:db \
	-v /etc/ssl/certs:/etc/ssl/certs \
	-v /data/oauth:/data/db \
	-e SEC_KEY="Bla" \
	-e REDIRECT_URL="http://www.wikipedia.org" \
	-e GOOGLE_CLIENT_ID="472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com" \
	-e GOOGLE_SECRET="OnkptU4BTdE45mi-b3hACdAY" \
	-e GOOGLE_REDIRECT_URL="http://localhost:8080/auth/google/callback" \
	vfarcic/oauth
```

After run, wait until MongoDB is initialized.

Backup
======

All MongoDB data is stored in the /data/db directory.

Display Help
============

```bash
sudo docker run --rm vfarcic/oauth oauth -help
```