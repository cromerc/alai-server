FROM node:16-bullseye AS builder
WORKDIR /app
COPY . .
RUN yarn install
RUN yarn build

FROM node:16-bullseye
RUN yarn global add serve
WORKDIR /app
COPY --from=builder /app/dist .
CMD ["serve", "-p", "3000", "-s", "."]
