# //TODO:

Items here may or may not have a GitHub issue associated with them. That's simply because I'm working on this on my own right now. I guess maybe I'll replace this with issues as time moves on.

## Return JSON instead of a web page.
Currently, we're returning a templated web page to the client when they send us a `GET`. That's not a REST API. Incidentally, this will resolve #1, since we will no longer be delivering a web page to the client. But we still want a front end, which brings us to:

## Make a separate front-end.
Yep. It will likely be its own project, but in the short term, we can use JavaScript to render our front-end by making requests to the API. Obviously we should gracefully degrade, so that if a client doesn't support JavaScript (read: has JavaScript disabled, or blocks its execution) we can fallback to a pure-HTML version. It may take being user-agent aware, by delivering a rendered page (like we have now) to browser user-agents, and the JSON results for non-browsers.

This is kind of secondary to the purpose of this project, though, hence it potentially spinning off into its own project.
