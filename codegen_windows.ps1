sqlc generate
Write-Host "Sqlc generate complete."

templ generate
Write-Host "Templ generate complete."

# Find all .fbs files in the current directory and subdirectories
Get-ChildItem -Recurse -Filter *.fbs | ForEach-Object {
    # Get the directory of the .fbs file
    $schema_dir = $_.Directory.FullName

    # Get the folder name (last part of the directory path)
    $folder_name = Split-Path $schema_dir -Leaf

    # Compile the .fbs file
    & .\flatc.exe -g -o $schema_dir $_.FullName --gen-onefile --go-namespace $folder_name
    & .\flatc.exe --ts -o web/schemas $_.FullName

    Write-Host "Compiled: $($_.FullName) to $schema_dir with namespace $folder_name"
}

Write-Host "`nCompilation completed successfully."
