FROM node:20

# Run as non root 'node' user for increased security
WORKDIR /usr/src/app
RUN chown node:node ./
USER node

# Default to production mode for optimized performance
ARG NODE_ENV=production
ENV NODE_ENV $NODE_ENV

# Install dependencies
COPY package*.json ./
RUN npm ci && npm cache clean --force

# Copy code
COPY index.js ./

# Run app
CMD ["node", "index.js"]