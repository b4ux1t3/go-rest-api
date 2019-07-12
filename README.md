# REST in Go

This is a simple project that I'm using as a demonstartion of implementing RESTful APIs in Go. This will be a toy implementation, but a fully "standards-compliant" implementation.

Right now, the only thing this program can do is handle requests for the `/collatz/` API endpoint and spit out a compiled template with the results of whatever is passed into the path after `/collatz/`. 

## Dependancies

This project depends on my own [`libgollatz`](https://github.com/b4ux1t3/libgollatz) library, which simply runs an algorithm based on the [Collatz conjecture](https://en.wikipedia.org/wiki/Collatz_conjecture).

## Why the Collatz Conjecture?

Why not? It's a reasonably simple mathematical conjecture which is easy to implement and can spit out an arbitraily large number of individual data points.

It also gives us the potential to examine a few different endpoints, such as `steps` or `values`, if we want to only check the number of steps it takes to get to 1, or which individual values we hit on our way to 1. As such, we can see how we might operate on the same data set with multiple API methods.

We also get the functionality of being able to generate different datasets based on different HTTP methods. For example, if I `GET` a value for `1337`, we can just do the algorithm and return the result down to the client. However, if we `PUSH` to the same endpoint, we can generate a cached copy of the algorithm's output.

So, if we request `GET /collatz/1337`, we'll receive a JSON object containing our values and the number of steps:

    
    {
        values: [4012, 2006, 1003, 3010, 1505,...8, 4, 2, 1],
        steps: 44
    }

However, if we `PUSH /collatz/1337`, the server would store the result of the algorithm to a database, and the client would receive a pointer to that entry.

## But it doesn't do any of that right now!

Right. And as such, the above section describes the ideal next steps. Check [TODO.md](TODO.md) for more information on what I have planned.