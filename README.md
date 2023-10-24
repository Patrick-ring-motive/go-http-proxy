# go-http-proxy

This is my first ever project in Go! 
I find creating a reverse proxy, using only the standard library, to be the best way to learn a new language after you have a few under your belt already.
This proxy is pointed at the golang official docs.
You can view the results here:

[https://go.patrickring.net](https://go.patrickring.net)

I require some sort of modification to the original site and to include routing for all of the subdomains linked from the home page.
I also tried to emulate the async/await promise structure using goroutines. I did acheive it to some extent but I am definitely not done iterating over those structures.

## Lessons Learned
I live Go! It is amazing that I could go from writing my first line of go, to a working reverse proxy in under 48 hours.
This is probably going to become my prefered backend webserver as I get more comfortable with it. I did this in conjuction with doing tfe sane project in Python. I have more experience in Python but it took me 4 times as long to build.
Concurrency is quite different as threads are cheap and I can easily spawn them with litlke consequence as long as I don't get carried away.
Go has a very interesting type system that feels strict at first but is actually only as strict as you make it. I'm still learning some of the ins and outs but I was able to create a promise struct to effectively emulate async/await. It works completely differently under the hood since their is no event loop. I may actually try to implement an event loop at done point after I've played around with goroutines some more.

