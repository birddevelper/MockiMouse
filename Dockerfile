FROM golang:1.19

# The exposed port should be changed to port number you specified in MockiMouse config.yml
EXPOSE 800
WORKDIR /MockiMouse

# Copy source code to working directory
COPY . .

# Build source code to mockimouse binary
RUN go build .

# Start MockiMouse server
CMD ./mockimouse