package http.authz

# default allow = false

yyy {
    input.method = "GET"
    input.path = ["salary", user]
    input.user = user
}

xxxx {
    input.method = "GET"
    input.path = ["salary", user]
    input.user = user
}