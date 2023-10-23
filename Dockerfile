# Use the official PostgreSQL image as the base image
FROM postgres:13-alpine

# Set the environment variables
ENV POSTGRES_DB=assessment
ENV POSTGRES_USER=devuser
ENV POSTGRES_PASSWORD=mypassword

# Specify the volume for data storage
VOLUME /var/lib/postgresql/data

EXPOSE 80