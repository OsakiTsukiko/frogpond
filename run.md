# HOW TO RUN (LINUX) (DEVEOPMENT)
- Prepare a PostgreSQL instance (or docker container)
- generate jwt secret key (run `tools/generate_jwt_secret_key.py`)
- export all required environment variables
```bash
# Server
FP_PORT="1234"
FP_JWT_SECRET_KEY="secret-key"
FP_DOMAIN="localhost" # or domain (mainly for cookies)

# DataBase
FP_DB_HOST="localhost" # mainly for running inside docker
FP_DB_PORT="5432" # default postgres is 5432
FP_DB_USERNAME="postgres" # default postgress is postgress
FP_DB_PASSWORD="postgres" # default postgress is postgress
FP_DB_DATABASE="database_name" # your database name
```
look into `tools/export_env.sh` for an easy way to export a whole .env file at once.
- go `run .`
