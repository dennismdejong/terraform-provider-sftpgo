# Configuration-based authentication
provider "sftpgo" {
  host     = "http://localhost:8080"
  username = "admin"
  password = "password"
}

# Alternative configuration with API key and disabled TLS verification (for dev/self-signed certs)
# provider "sftpgo" {
#   host             = "https://sftpgo.example.com"
#   api_key          = "..."
#   tls_verification = false
# }