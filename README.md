# mTLS

A simple example showing mutually-authenticated TLS (mTLS) between a client and server over HTTPS in Go.

### Setup

1. Edit your hosts file:
```
$ sudo nano /etc/hosts
```
Add the line:
```
127.0.0.1   the-server
```
Example:
```
##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1   the-server
127.0.0.1   localhost
...
```
2. Install minica:

Go here for installation instructions - https://github.com/jsha/minica

3. Clone this repo:
```
$ git clone github.com/umahmood/mtls
```
4. Generate certificates:
```
$ cd mtls/ca
$ minica -ca-cert minica.pem -ca-key minica-key.pem -domains the-client
$ minica -domains the-server
```
5. Copy certs to client and server directories:
```
$ cd mtls/ca
$ cp the-client/* ../client
$ cp the-server/* ../server
```

#### Running the example:

In a terminal run the server:
```
$ cd mtls/server
$ go run main.go
```
In another terminal run the client:
```
$ cd mtls/client
$ go run main.go
```
You should get the output:
```
Status: 200 OK  Body: Hello mTLS World!
```

**Tip**: Inside the client directory, run:
```
$ curl --trace trace.log --cacert ../ca/minica.pem \
    --cert cert.pem --key key.pem https://the-server:8080
```
Open trace.log and you can see the TLS handshake and lots of other information.

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
