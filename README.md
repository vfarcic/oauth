Explanation
===========

* After successful login, **authName**, **authAvatarURL** and **authID** are set and ready for your usage.
* **/auth/google/login** will redirect to the Google authentication page, receive response, store usar data to the database and, finally, redirect the user to the URL specified with the environment variable **REDIRECT_URL** with **authID** added as a query parameter.
* GET request to the **/auth/api/v1/user/[AUTH_ID] returns JSON with user's data. **AUTH_ID** should be replaced with the value from the **authID query or cookie**. This request can be used to obtain more information about the user or to validate it's identity.
* Who Am I Web component can be imported and displayed on a page.

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

* DOMAIN: Domain in which the application is running. Defaults to **localhost**.
* PORT: Port in which the application is running. Defaults to **8080**.
* SEC_KEY: Security key. This is mandatory variable.
* REDIRECT_URL: URL where users will be redirected to after authenticating. This is mandatory variable.
* GOOGLE_CLIENT_ID: Google client ID. This is mandatory variable if google authentication is used.
* GOOGLE_SECRET: Google client secret. This is mandatory variable if google authentication is used.
* GOOGLE_REDIRECT_URL: Google redirect URI. This is mandatory variable if google authentication is used.

Run
---

```bash
sudo docker run -d --name oauth \
	-p 8080:8080 \
	-v /etc/ssl/certs:/etc/ssl/certs \
	-v /data/oauth:/data/db \
	-e PORT=8080 \
	-e SEC_KEY="Bla" \
	-e REDIRECT_URL="http://localhost:8080/components/test.html" \
	-e GOOGLE_CLIENT_ID="472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com" \
	-e GOOGLE_SECRET="OnkptU4BTdE45mi-b3hACdAY" \
	-e GOOGLE_REDIRECT_URL="http://localhost:8080/auth/google/callback" \
	vfarcic/oauth
```

After run, wait until MongoDB is initialized.

Embed "Who Am I" Web Component
==============================

```html
<html>
<head>
	<!--Import Required Polymer Components-->
    <link rel="import" href="polymer/polymer.html">
    <link rel="import" href="iron-ajax/iron-ajax.html">
    <link rel="import" href="iron-image/iron-image.html">
    <!--Import "Who Am I" Component-->
	<link rel="import" href="http://localhost:8080/components/oauth/whoami.html">
</head>
<body>
	<!--Display "Who Am I" Component-->
	<who-am-i></who-am-i>
</body>
```

Backup
======

All MongoDB data is stored in the /data/db directory.


Display Help
============

```bash
sudo docker run --rm vfarcic/oauth oauth -help
```