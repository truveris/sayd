# sayd - OS X say(1) as a service

Produce streamable spoken sentences via a simple HTTP service.

We tried to tweak espeak on Linux to sound natural but couldn't get close
enough to the Mac OS X `say(1)` command. Instead of replacing all of our
speaking raspberry pis with mac minis we thought it was simpler to use mplayer
on these machines to reach out to a simple service running on an old macbook.

Both GET and POST return an AIFF (Content-type: audio/aiff) that can be played
via a browser, mplayer or VLC.


## GET /[voice].[format]?[sentence]
That's the easiest way to test with your browser, the following URL gets Alex
to say hello:

    http://localhost:9999/alex?hello

If your browser does not support AIFF and you have ffmpeg installed on the sayd
server, you can also try:

    http://localhost:9999/alex.mp3?hello

## POST /[voice].[format]
Use this from curl or your app, the body of the POST contains the sentence or
words you want to get. The advantage is that you don't need to worry about URL
encoding.


## Installation
Copy sayd.conf to /etc and change the port if needed, start it with "sayd" or
wrap it with supervisor.


## Requirements
 - Go 1.2+
 - A Mac.
 - github.com/jessevdk/go-flags

