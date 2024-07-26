path "be/data/*" {
  capabilities = ["create", "read", "update"]
}

path "be/metadata/*" {
  capabilities = ["list"]
}