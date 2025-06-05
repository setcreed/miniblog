# miniblog 开发中常用命令

## Dump 表结构

```bash
mysqldump -h127.0.0.1 -uminiblog --databases miniblog -p'miniblog1234' --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > /tmp/miniblog.sql
```