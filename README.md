# CI Bully

CI bully is a go program that will warn and close github PRs that are long living. You can define the actions in the configuration.
Ideally we want to developers to avoid using long living branches.

It is better to run PR bully from a CI daily. Keep in mind CI bully does not keep state or check if it posted before.

## Config
```yaml
---
## Define the token in yaml or environment variable "GITHUB_TOKEN"
#token: "XxxXXXXxxxx"

## Only count the workdays. Defaults to false
#only_workdays: true 

actions:
    # if you specify last anything greater than last will be enforced
  - day: 14
    last: true
    action: close
    message: |
              Hi _USER_ this PR exceeded _SINCE_ days in open state.
              We are trying to encourage developers to integrate with master quicker ideally daily.
              I will **close** this PR now.
              Please open a new PR if this branch is still needed.

  - day: 12
    action: warn
    message: |
              Hi _USER_ this PR exceeded _SINCE_ days in open state.
              We are trying to encourage developers to integrate with master quicker ideally daily.
              This is the last warning I will close this PR in _TILL_ days.
  - day: 7
    action: warn
    message: |
              Hi _USER_ this PR exceeded _SINCE_ days in open state.
              We are trying to encourage developers to integrate with master quicker ideally daily.
              I will close this PR in _TILL_ days.

  - day: 1
    action: warn
    message: |
              Hi _USER_ we are trying to encourage developers to integrate with master quicker ideally daily.

repos:
 - "ahelal/avm"
 - "ahelal/t-template"
 - "ahelal/ansible-concourse"
```

## Running CI-bully

CI bully does not keep state so it is important to run only once a day always at the same time, if you don't want to spam people.

