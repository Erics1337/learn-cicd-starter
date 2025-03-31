FROM debian:stable-slim

# Try updating and installing keyring first, check date, then update again
RUN apt-get update && apt-get install -y ca-certificates

ADD notely /usr/bin/notely

CMD ["notely"]
