#!/usr/bin/env bash

# Usage: ./extract_methods.sh "Repository" /path/to/directory [--lowercase]

if [ -z "$1" ] || [ -z "$2" ]; then
  echo "Usage: $0 <StructName> <Directory> [--lowercase]"
  exit 1
fi

STRUCT_NAME="$1"
DIRECTORY="$2"
LOWERCASE=false

if [ "$3" == "--lowercase" ]; then
  LOWERCASE=true
fi

# Find all .go files in the specified directory and process them
find "$DIRECTORY" -name "*.go" | while read -r file; do
  if [ "$LOWERCASE" = true ]; then
    # Extract lines containing the desired struct's methods starting with a lowercase letter
    grep -E "^\s*func \(\w+ \*${STRUCT_NAME}\) [a-z]" "$file" | \
      sed -E "s/^\s*func \(\w+ \*${STRUCT_NAME}\)\s+//" | \
      sed -E "s/\s*\{//"
  else
    # Extract lines containing the desired struct's methods starting with an uppercase letter
    grep -E "^\s*func \(\w+ \*${STRUCT_NAME}\) [A-Z]" "$file" | \
      sed -E "s/^\s*func \(\w+ \*${STRUCT_NAME}\)\s+//" | \
      sed -E "s/\s*\{//"
  fi
done
