#!/bin/bash


# Fetch the latest tags from the remote repository
git fetch --tags

# Get the current Git tag
VERSION=$(git describe --tags --abbrev=0)

if [ -z "$VERSION" ]; then
    echo "No version found. Exiting..."
    exit 1
fi


# Remove the "v" prefix (if it exists) and trim any leading/trailing spaces
VERSION=$(echo "$VERSION" | sed 's/^[[:space:]]*V//' | sed 's/^[[:space:]]*//')  # Strip spaces and "v"
VERSION=$(echo "$VERSION" | sed 's/^[.]//')  # Remove any leading dot if it exists

echo "Cleaned version: $VERSION"  # Debug output

# Split the version string into components (major, minor, patch)
IFS='.' read -r -a VERSION_PARTS <<< "$VERSION"

# Check if we have exactly three parts (major, minor, patch)
if [ ${#VERSION_PARTS[@]} -ne 3 ]; then
    echo "Invalid version format. Expected major.minor.patch format, got: $VERSION"
    exit 1
fi

# Increment the patch version (or change this logic to increment minor or major version)
PATCH=${VERSION_PARTS[2]}
PATCH=$((PATCH + 1))
VERSION_PARTS[2]=$PATCH

# Rebuild the version string
NEW_VERSION="${VERSION_PARTS[0]}.${VERSION_PARTS[1]}.${VERSION_PARTS[2]}"

# Update the version in the Cargo.toml files
sed -i.bak "s/^version = .*/version = \"$NEW_VERSION\"/" Cargo.toml
sed -i.bak "s/^version = .*/version = \"$NEW_VERSION\"/" go_generation_derive/Cargo.toml
sed -i.bak "s/^version = .*/version = \"$NEW_VERSION\"/" parser/Cargo.toml
sed -i.bak "s/^version = .*/version = \"$NEW_VERSION\"/" parser_unsafe/Cargo.toml
sed -i.bak "s/^version = .*/version = \"$NEW_VERSION\"/" siri_parser_python/Cargo.toml

echo "Updated Cargo.toml files to version $NEW_VERSION."




