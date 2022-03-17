package http.authz

default allow = false

allow {
    input.method = "GET"
    input.path = ["salary", user]
    input.user = user
}

xxxxy {
    input.method = "GET"
    input.path = ["salary", user]
    input.user = user
}