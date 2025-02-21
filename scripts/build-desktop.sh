#!/usr/bin/env bash

# Usage: ./extract_methods.sh "Repository" /path/to/directory [--lowercase]

if [ -z "$1" ] || [ -z "$2" ]; then
  echo "Usage: $0 <StructName> <Directory> [--lowercase]"
  exit 1
fi

mkdir -p "$HOME/.local/share/applications"
cat > "$DESKTOP_FILE" <<EOF
[Desktop Entry]
Name=${app_name}
Comment=Infrastructure as Code Manager
Exec=$EXEC_PATH
Icon=amadla
Terminal=true
Type=Application
Categories=Development;Utility;
MimeType=application/x-hery;
EOF
echo "✅ Created .desktop file at $DESKTOP_FILE"
