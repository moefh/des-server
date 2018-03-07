# des-server

The goal of this project is to reverse-engineer the protocol of Demon's Souls server
and maybe eventually write a replacement server.

## What is done so far

This is a Go project in *very* early stages. This only works for the NA (Norht America)
version of Demon's Souls.  This is currently just [Asapin's proof-of-concept server](https://github.com/Asapin/ds-server-PoC),
rewritten in Go and adapted to the NA Damon's Souls client. This means we're able to start
an online session that is closed very shortly by the game client because the server doesn't
reply to any requests beyond the initial login.

## How to setup/use this server

These are step-by-step instructions to reproduce my setup (it's a Windows machine, but since
everything I'm using is in Go, it should also work on Mac, Linux and FreeBSD). If you already
have experience with GitHub and/or Go, or just want to do things differently, feel free to
ignore any of this.

### Preparation

You need a PC that is reachable from the PS3 (e.g., you're able to stream movies or music
from the PC to the PS3).

You'll also need to [install Go](https://golang.org/doc/install) if you haven't already.

And finally, you'll have to get the files of this project either with `git clone` or by
downloading a ZIP from GitHub. Either way, I'll assume you've put the project files in
`cd c:\Users\YOUR_USERNAME\go\des-server`, so replace this path in the following instructions.
Just remember that since it's a Go project, it will have to be in a directory inside your Go
home.

### Redirect Demon's Souls Server DNS

You'll have to find a way to redirect the Demon's Souls server domain name (`demons-souls.com`)
to the IP address of your PC. I'm doing this by running the [CoreDNS](https://github.com/coredns/coredns)
DNS server. If you have another way of doing this, just skip this step.

    cd c:\Users\YOUR_USERNAME\go
    go get github.com/coredns/coredns

This should (after a while) produce a `coredns.exe` file in `c:\Users\YOUR_USERNAME\go\bin\`.

Now edit the file `db.demons-souls.com` found in the `dns` directory of this project and change
the 3 occurrences of `192.168.15.10` to the IP address of your computer, save the file and run
    
    cd c:\Users\YOUR_USERNAME\go\des-server\dns
    c:\Users\YOUR_USERNAME\go\bin\coredns.exe

Don't close this terminal, it has to be left open while CoreDNS is running. This is useful because
CoreDNS will log every server name it resolves, so you can see if it's working.

Finally, just change the DNS server of your PS3 to the IP address of your computer. After that, do a
"Network Test" on your PS3, you should see CoreDNS resolving some `playstation.net` host names in its
terminal window.

### Run the Demon's Souls Server

Open another terminal and run

    cd c:\Users\YOUR_USERNAME\go\des-server
    go build
    des-server.exe

This will show "`Starting server...`", meaning the server is running. Now open Demon's Souls in your
PS3 and load or start a new game. If everything goes right, you should see this server's announcement
message and be able to start the game online (but it will soon tell you that the connection was lost).

Notice that the server logs all requests made by the client (even though it doesn't respond to most of them).
