api_addr = "http://127.0.0.1:5200"
cluster_addr = "https://127.0.0.1:5201"
ui = true
disable_mlock = true
disable_sealwrap = true

storage "file" {
  path    = "./vault/data"
  node_id = "node1"
}

listener "tcp" {
  address     = "0.0.0.0:5200"
  tls_disable = "true"
}

path "kv/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
path "kv/my-secret" {
  capabilities = ["read"]
}