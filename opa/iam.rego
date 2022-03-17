package aws

default allow = false

# FirstStatement
allow {
    principals_match
    input.action == "iam:ChangePassword"
}

# SecondStatement
allow {
    principals_match
    input.action == "s3:ListAllMyBuckets"
}

# ThirdStatement
#  Use helpers to handle implicit OR in the AWS policy.
#  Below all of the 'principals_match', 'actions_match' and 'resources_match' must be true.
allow {
    principals_match
    actions_match
    resources_match
}

# principals_match is true if input.principal matches
principals_match {
    input.principal == "alice"
}

# actions_match is true if input.action matches one in the list
actions_match {
    # iterate over the actions in the list
    actions := ["s3:List.*","s3:Get.*"]
    action := actions[_]
    # check if input.action matches an action
    regex.globs_match(input.action, action)
}

# resources_match is true if input.resource matches one in the list
resources_match {
    # iterate over the resources in the list
    resources := ["arn:aws:s3:::confidential-data","arn:aws:s3:::confidential-data/.*"]
    resource := resources[_]
    # check if input.resource matches a resource
    regex.globs_match(input.resource, resource)
}