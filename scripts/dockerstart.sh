until mysql -h $MARIA_HOST -p$MARIA_PASSWORD; do
    >&2 echo "mariadb is unavailable - sleeping"
    sleep 5
done
>&2 echo "mariadb is up - executing commands"

./engine