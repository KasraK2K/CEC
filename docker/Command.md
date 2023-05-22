# Commands

## Backup PostgreSQL Database

```bash
docker exec -t your-db-container pg_dumpall -c -U postgres | gzip > dump_$(date +%Y-%m-%d_%H_%M_%S).gz
```
&nbsp;

## Restore PostgreSQL Database

```bash
gunzip < your_dump.sql.gz | docker exec -i your-db-container psql -U postgres
```
&nbsp;

## Run Docker

```bash
make up_build
```