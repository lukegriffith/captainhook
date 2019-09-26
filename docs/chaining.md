# Chaining


Executing the standalone server with the below config has hook1, forward its 
result via calling localhost, to the rules hook2 endpoint. 

```
/usr/lib/go/bin/go build -o /tmp/___cmd_standalone /go/src/github.com/lukemgriffith/captainhook/cmd/standalone/main.go #gosetup
/tmp/___cmd_standalone #gosetup
CaptainHook 2019/09/25 20:27:55 Starting Application Server.
2019/09/25 20:27:55 loading config.yml
2019/09/25 20:27:55 endpoints:
```

In this config, rule hook1 templates out a json object and forwards it to the 
rule hook2. 

```yaml
  - name: hook1
    secret: test
    rules:
      - destination: http://localhost:8081/webhook/hooks
        template: |
          { "test" : "{{ .test }}" }
  - name: hook2
    secret: supersecret
    rules:
      - destination: http://localhost:8082
        template: |
          { "hook_executed": "{{ .test }}" }
```

The standalone server reads the config in and loads it into memory. 

```
2019/09/25 20:27:55 {[{test test [{http://localhost:8081/webhook/hooks { "test" : "{{ .test }}" }
}] []} {hooks supersecret [{http://localhost:8082 { "hook_executed": "{{ .test }}" }
}] []}] }
```

Running a curl command, it can trigger the test rule to start the chain. 

```bash 
$ curl http://localhost:8081/webhook/test -X POST --data '{"test":"payload123"}'
```

The hooks are executed templating out hook1, then forwarding and moving to hook2.

As port 8082 is not the application, it stops the chain.

```

CaptainHook 2019/09/25 20:28:00 processing webhook
CaptainHook 2019/09/25 20:28:00 Rendered template:  { "test" : "payload123" }

CaptainHook 2019/09/25 20:28:00 Forwarding to http://localhost:8081/webhook/hooks
CaptainHook 2019/09/25 20:28:00 processing webhook
CaptainHook 2019/09/25 20:28:00 Rendered template:  { "hook_executed": "payload123" }

CaptainHook 2019/09/25 20:28:00 Forwarding to http://localhost:8082

```

Be careful not to execute a recursive loop as the thread will crash maxing out file descriptors.
