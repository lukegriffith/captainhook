# CaptainHook

Cloud based awk, transform and glue webhooks and api's together.


```bash
.\captainhook help
serve: Start the CaptainHook application server.
  -configPath string
        YAML file to configure the service with endpoints. (default "config.yml")
  -passphrase string
        Passphrase for encrypted YAML blob.
  -port string
        TCP port for server to run, default is ':8081' (default ":8081")
  -secretPath string
        Encrypted YAML blob containing string map of secrets.

encrypt: Perform encryption operations on a yaml file.
  -decrypt
        should the file be decrypted
  -filepath string
        File to perform encryption operation.
  -passphrase string
        Passphrase for encrypted YAML blob.
```

## Docs
 
- [Chaining](docs/chaining.md)
- [DataMap & Secrets](docs/DataMap.md)
