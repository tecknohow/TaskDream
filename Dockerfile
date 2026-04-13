# Stage 1: Build Go backend
FROM golang:1.22-alpine AS backend-builder
RUN apk add --no-cache gcc musl-dev sqlite-dev
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o taskdream .

# Stage 2: Build Vue frontend
FROM node:20-alpine AS frontend-builder
RUN corepack enable
WORKDIR /build
COPY frontend/package.json frontend/pnpm-lock.yaml* ./
RUN pnpm install --frozen-lockfile || pnpm install
COPY frontend/ .
RUN pnpm build

# Stage 3: Runtime
FROM alpine:3.19
RUN apk add --no-cache ca-certificates sqlite-libs tzdata
WORKDIR /app
COPY --from=backend-builder /build/taskdream .
COPY --from=frontend-builder /build/dist ./frontend/dist
COPY config.example.yml ./config.example.yml
EXPOSE 3456
VOLUME /app/data
ENV TASKDREAM_DATABASE_TYPE=sqlite
ENV TASKDREAM_DATABASE_PATH=/app/data/taskdream.db
ENTRYPOINT ["./taskdream"]
CMD ["web"]
