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

As you can see comments are preserved when using **yaml.v3**