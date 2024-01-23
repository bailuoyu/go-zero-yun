Global:
    Namespace: Development
    EnvName: dev

Server:
    App:yourapp
    Rest:
        Name: {{.serviceName}}
        Host: {{.host}}
        Port: {{.port}}
