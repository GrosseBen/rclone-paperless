# Justfile for building and testing the rclone plugin

# Build rclone binary
build-rclone:
    mkdir -p bin
    cd ../../rclone/rclone && go build -o ../../GrosseBen/rclone-paperless/bin/rclone

# Setup replace directive for local development
setup-plugin:
    cd plugins/example && \
    go mod edit -replace github.com/rclone/rclone=../../../../rclone/rclone && \
    go mod tidy

# Build the plugin using local rclone
build: setup-plugin
    mkdir -p bin
    cd plugins/example && go build -buildmode=plugin -o ../../bin/librcloneplugin_example.so

# Test the plugin build (checks if .so file exists)
test:
    test -f bin/librcloneplugin_example.so && echo "Plugin built successfully"

# Test loading the plugin with rclone via RCLONE_PLUGIN_PATH
test-load: build-rclone
    RCLONE_PLUGIN_PATH=bin ./bin/rclone help backends | grep -i example

# Clean the build artifacts
clean:
    rm -f bin/librcloneplugin_example.so
    rm -f plugins/example/go.sum
    rm -rf bin/
