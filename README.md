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

Run With Docker
---------------

Following environment variables can be used. 

* DOMAIN: Domain in which the application is running. Defaults to **localhost**.
* PORT: Port in which the application is running. Defaults to **8080**.
* SEC_KEY: Security key. This is mandatory variable.
* REDIRECT_URL: URL where users will be redirected to after authenticating. This is mandatory variable.
* GOOGLE_CLIENT_ID: Google client ID. This is mandatory variable if google authentication is used.
* GOOGLE_SECRET: Google client secret. This is mandatory variable if google authentication is used.
* GOOGLE_REDIRECT_URL: Google redirect URI. This is mandatory variable if google authentication is used.

```bash
sudo docker run -d --name oauth \
	-p 8080:8080 \
	-v /etc/ssl/certs:/etc/ssl/certs \
	-v /data/oauth:/data/db \
	-e PORT=8080 \
	-e SEC_KEY="Bla" \
	-e REDIRECT_URL="http://localhost:8080/components/oauth/demo/index.html" \
	-e GOOGLE_CLIENT_ID="472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com" \
	-e GOOGLE_SECRET="OnkptU4BTdE45mi-b3hACdAY" \
	-e GOOGLE_REDIRECT_URL="http://localhost:8080/auth/google/callback" \
	vfarcic/oauth
```

After run, wait until MongoDB is initialized.

Run Executable
--------------

Run `./oauth -help` to get the list of arguments.

```bash
sudo docker run --name mongo -d \
	-p 27017:27017 \
	mongo

./oauth \
	-sec-key="Bla" \
	-redirect-url="http://localhost:8080/components/oauth/demo/index.html" \
	-google-client-id="472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com" \
	-google-secret="OnkptU4BTdE45mi-b3hACdAY" \
	-google-redirect-url="http://localhost:8080/auth/google/callback"
```

Wait until MongoDB is initialized.

Open [http://localhost:8080/auth/google/login](http://localhost:8080/auth/google/login).

Embed "Who Am I" Web Component
==============================

```html
<html>
<head>
	<!--Import Required Polymer Components-->
    <script src="../bower_components/webcomponentsjs/webcomponents-lite.min.js"></script>
    <link rel="import" href="../bower_components/polymer/polymer.html">
    <link rel="import" href="../bower_components/paper-styles/classes/global.html">
    <link rel="import" href="../bower_components/iron-ajax/iron-ajax.html">
    <link rel="import" href="../bower_components/iron-image/iron-image.html">
    <link rel="import" href="../bower_components/paper-button/paper-button.html">
    <link rel="import" href="../bower_components/paper-item/paper-item.html">
    <link rel="import" href="../bower_components/paper-item/paper-item-body.html">
    <!--Import OAuth Components-->
	<link rel="import" href="http://localhost:8080/components/oauth/oauth-who-am-i.html">
	<link rel="import" href="http://localhost:8080/components/oauth/oauth-providers.html">
</head>
<body>
	<oauth-providers></oauth-providers>
	<oauth-who-am-i></oauth-who-am-i>
</body>
```

Following properties can be used with **oauth-who-am-i**:

* backendHost
* hide-avatar
* hide-full-name
* hide-email
* hide-log-out
* auth-id

Following properties can be used with **oauth-providers**:

* backendHost

Backup
======

All MongoDB data is stored in the /data/db directory.


Display Help
============

```bash
sudo docker run --rm vfarcic/oauth oauth -help
```

Testing
=======

Prequisites
-----------

```bash
sudo npm install
```

Running
-------

```bash
go test -cover

gulp test:local
```