FROM golang:1.24
ARG APP_PORT

ARG ENV=prod
ENV ENV=${ENV}

WORKDIR /usr/src/app

# Install Ari live reloading
RUN if [ "$ENV" = "dev" ]; then \
      go install github.com/air-verse/air@latest ; \
    fi

COPY go.mod go.sum ./
RUN go mod tidy # Cleaning dependence
RUN go mod download

COPY . ./

# Prod building
RUN if [ "$ENV" = "prod" ]; then \
      go build -v -o /usr/local/bin/app ./cmd; \
    fi

EXPOSE ${APP_PORT}

# Run application depending on the environment
CMD if [ "$ENV" = "dev" ]; then \
      air -c .air.toml; \
    else \
      /usr/local/bin/app; \
    fi