#!/bin/bash

sqlc generate
echo "Sqlc generate complete."

templ generate
echo "Templ generate complete."

# Find all .fbs files in the current directory and subdirectories
find . -name "*.fbs" | while read fbs_file; do
    # Get the directory of the .fbs file
    schema_dir=$(dirname "$fbs_file")

    # Get the folder name (last part of the directory path)
    folder_name=$(basename "$schema_dir")

    # Compile the .fbs file
    ./flatc -g -o "$schema_dir" "$fbs_file" --gen-onefile --go-namespace "$folder_name"
    ./flatc --ts -o web/schemas "$fbs_file"

    echo "Compiled: $fbs_file -> $schema_dir using namespace $folder_name"
done

echo -e "\nCompilation completed successfully."