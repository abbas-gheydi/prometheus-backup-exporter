# prometheus-backup-exporter
## About The Project
prometheus backup exporter is designed to monitor backups.   
it gets backup sizes and stores them in the prometheus.   
 
### How to build
```bash
go get ./..
go build -o backup-exporter ./
```
### How to Use
edit config.yml:   
```yml
Server:
        Port: 18083
DailyBackups:
        - "/home/ubuntu/backup/db/backend-db.sql"
        - "/home/ubuntu/tar/backup.tar"
        - "/home/ubuntu/backup/constant.json"
```

here are some prometheus alert rules example:   
```yaml
  rules.yml: |
   groups:
   - name: example
     rules:


     - alert: backups size is wrong
       expr: (backup_size == 0) or (backup_size {backup!="_home_ubuntu_backup_constant_json"} <= backup_size offset 1d)
       for: 1m
       labels:
        severity: critical

     - alert: backups maybe failed
       expr: backup_date <= backup_date offset 1d
       for: 1m
       labels:
        severity: critical
```
