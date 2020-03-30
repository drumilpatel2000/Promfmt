## Output of the promfmt

```
groups:
  - name: my-group-name
    interval: 30s # defaults to global interval
    rules:
      - alert: HighErrors
        expr: "sum without(instance) (rate(errors_total[5m]))\n/ \nsum without(instance)
            (rate(requests_total[5m]))\n"
        for: 5m
        labels:
            severity: critical
        annotations:
            description: "stuff's happening with {{ $.labels.service }}"
      - # Multiline comments
        # is working
        record: "new_metric"
        expr: "sum without(instance) (rate(errors_total[5m]))\n/ \nsum without(instance)
            (rate(requests_total[5m]))\n"
        labels:
            abc: edf
            uvw: xyz
``` 

## Parsing Yaml File To Get Promql Expressions
```
groups


name
my-group-name
interval
30s
Line # defaults to global interval
rules


alert
HighErrors
##################################################################
My Expressions
sum without(instance) (rate(errors_total[5m]))
/ 
sum without(instance) (rate(requests_total[5m]))

for
5m
labels

severity
critical
annotations

description
stuff's happening with {{ $.labels.service }}

record
Head # Multiline comments
# is working
new_metric
##################################################################
My Expressions
sum without(instance) (rate(errors_total[5m])) / sum without(instance) (rate(requests_total[5m]))
Line # Internal Comments 
labels

abc
edf
uvw
xyz
```

As you can see comments are preserved when using **yaml.v3**