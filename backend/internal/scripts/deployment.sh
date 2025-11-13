#!/bin/bash

# -------------------------------
# Configuration
# -------------------------------
PROJECTS_DIR="/mnt/z/projects"
TARGET_DIR="$PROJECTS_DIR/Gole-Certo"
REPO_URL="https://github.com/Mateus-MS/Gole-Certo.git"
BRANCH_NAME="refactor"
TEMP_DIR="/home/$(whoami)/gole-certo-temp"

echo "Starting deployment of Gole-Certo..."

# -------------------------------
# Step 1: Clean up temporary folder in WSL
# -------------------------------
if [ -d "$TEMP_DIR" ]; then
    echo "Removing existing temp folder: $TEMP_DIR"
    rm -rf "$TEMP_DIR"
fi
mkdir -p "$TEMP_DIR"

# -------------------------------
# Step 2: Clone repository locally in WSL
# -------------------------------
echo "Cloning branch '$BRANCH_NAME' from $REPO_URL into $TEMP_DIR..."
git clone --branch "$BRANCH_NAME" "$REPO_URL" "$TEMP_DIR"

if [ $? -ne 0 ]; then
    echo "Git clone failed!"
    exit 1
fi
echo "Clone completed successfully."

# -------------------------------
# Step 3: Delete old Gole-Certo folder on network share
# -------------------------------
if [ -d "$TARGET_DIR" ]; then
    echo "Deleting old Gole-Certo folder: $TARGET_DIR"
    rm -rf "$TARGET_DIR"
fi

# -------------------------------
# Step 4: Move the new clone to the network share
# -------------------------------
echo "Moving new clone to: $TARGET_DIR"
mv "$TEMP_DIR" "$TARGET_DIR"

if [ $? -eq 0 ]; then
    echo "Deployment completed successfully!"
else
    echo "Failed to move the new clone to $TARGET_DIR"
    exit 1
fi

# -------------------------------
# Done
# -------------------------------
echo "Deployment finished."
