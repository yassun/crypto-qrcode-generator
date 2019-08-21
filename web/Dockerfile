FROM node:11.1
RUN npm install -g @vue/cli

WORKDIR /app
COPY . .
RUN yarn install --modules-folder /node_modules

EXPOSE 8080

CMD ["yarn", "serve"]
