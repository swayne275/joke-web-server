# joke-web-server
This is a production-ready web service that returns a random Chuck Norris joke

# Nerdy stuff
The server works by combining two existing web services - a Chuck Norris jokes API anda random name generator API. It uses `icndb` to get the base form of a joke, and the `randomuser.me` name generator API.

The web server is concurrent to handle multiple requests (essentially) efficiently. Because the jokes API call relies on the results of the names API call we cannot run the calls simultaneously (unless we build some local cache of names ready to feed into the jokes api).

I broke the app up into various packages for modularity, which came in handy since the names API I had originally intended to use seems somewhat flaky (`uinames`). This made it easy to write two names packages that share the same API, so they can easily be swapped out in the calling code.

Some logic could further be shared in a "common" package (such as the log functions),
but since the app is (currently) so small it seemed like a coin toss on whether this was useful or a case of over-engineering. I condensed http get calls into a very minimal package to avoid some logic duplication and simplify some calls. On one hand it would be nice for the package to have direct access to HTTP, but on the other hand it can be nice to have a very simple call with a "yes/no" to the "did it work" question.

I used the `gjson` package for more deeply nested JSON responses (from the `randomuser.me` API), and used structs to unmarshal less nested JSON.

I added the ability to specify gender in the names package because some of the `icndb` jokes were inherently gendered (since they are about Chuck Norris). The `icndb` api did not detail how to specify gender. I tried adding a gender query parameter into
`icndb` calls and it did not work. I kept the joke topic nerdy because... why not?

I chose javascript encoding so that the results would look nice via curl/Postman, as that's a common way to interact with a server. Note that `icndb` uses HTML encoding by default for easy insertion into a web page.

Errors returned via APIs do not reveal specifically what went wrong (for separation of concerns), but those more-specific errors are logged by the individual packages. Status code 500 is returned to the user if anything goes wrong.

I used Docker for portability and easy deployment, and implemented a multi-stage build to generate a better production image (smaller, less attack surface, etc.). I also used Docker Compose to simplify running the app.

# Usage
Make sure you have the newest version of Docker installed (tested using Docker Engine 20.10.13). If you are using Docker Compose, ensure you're running the latest version (tested on 1.29.2)

`cd` to the top level directory of this project (the one containing the `Dockerfile`)

## Running the app
If using `Docker`, run the following commands (using port `5001` on your local machine):
Build the image: `docker build -t joke-web-server .`  
Run the image (interactively): `docker run -p 5001:5000 joke-web-server`  

If you'd rather run the image in the background you can swap the run command with:  
`docker run -d -p 5001:5000 joke-web-server`  

If using `Docker Compose`, run: `docker-compose up`  
If you'd like it to run in the background instead: `docker-compose up -d`  

## Getting a joke
If running and testing locally, you can use `curl`: `curl localhost:5001`  

Note: You can also interact with the app using a web browser (or API testing tool such as `Postman`) at `localhost:5001`

# Notes
A few of the jokes in `icndb` seem to be subtly broken, appending a `'ss` to the name instead of a `'s` to show ownership. For example, `icndb` returned this:
`Marcus Patel'ss beard can type 140 wpm.` when given this name: `Marcus Patel`. I validated to ensure this app wasn't mishandling anything related to that.