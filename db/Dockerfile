FROM migrate/migrate

# Copy all db files
COPY ./migrations /migrations

ARG DB_PASS
ARG DB_NAME
ARG DB_PORT

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]
CMD ["migrate", "-path", "/migrations", "-database", "mysql://root:${DB_PASS}@tcp(meal-planner.db:${DB_PORT}):/${DB_NAME}"]
