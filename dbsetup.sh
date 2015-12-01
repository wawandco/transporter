docker-compose run lib psql -c "DROP DATABASE IF EXISTS ttest;" -h postgres.local -U postgres
docker-compose run lib psql -c "CREATE DATABASE ttest;" -h postgres.local -U postgres
