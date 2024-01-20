# simple-go-template

To be used as a template project to bootstrap other projects

## Run Worker
```bash
go run worker/main.go
```

## Start Workflow
```bash
temporal workflow start --type SimpleGo --task-queue simple-task-queue --input '{"Val":"foo"}'
```
