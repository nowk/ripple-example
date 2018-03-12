![ripple-expample][splash]

> Example for [ripple](https://github.com/nowk/ripple)


## Setup

Checkout

    $ go get github.com/nowk/ripple-example


Use `dep` to get our dependencies

    $ dep ensure


Start the server

    $ make            (installs dumb-init, so we can exit out of docker)
    $ make server


---

    $ curl <yourip>:4000/api/v1/items

    => [{"foo": "bar"}]


[splash]: https://s3.amazonaws.com/assets.github.com/splash-ripple-example.svg

