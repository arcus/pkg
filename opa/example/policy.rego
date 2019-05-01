package arcus.pkg.opa.example

default allow = false

allow {
	input.user = "joe"
}

allow {
	input.role = "admin"
}

get_perms = perms {
	input.user = "joe"
	perms = {"x", "y", "z"}
}

get_perms = perms {
	input.role = "admin"
	perms = {"a", "b"}
}
