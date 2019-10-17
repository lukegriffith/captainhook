# DataMap and Secrets

DataMap is a structure that contains data related to the endpoint. Input parameters from incoming requests and access to
configured secrets. In an endpoints configuration, the secrets property can be provided as a list of secret names that are 
configured in the hook engine, these are then passed into the _secrets key on the DataMap accessible in rules.  


*Config.yml*

```yml
  - name: hooks
    secrets:
      - hookSecret
    rules:
      - type: template
        destination: http://localhost:8082
        arguments:
          template: |
            { "hook_executed": "{{ .test }}", "hook_secret": "{{ ._secrets.hookSecret }}" }
```


*Secrets.yml*

```yml
hookSecret: Secret Material
test123: test secret
```

In the above configuration, the hooks endpoint only has the hookSecret secret available in the DataMap, as its been
configured on the structure.