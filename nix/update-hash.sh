#!/usr/bin/env bash
# This script only works with a flakes enabled version of Nix

sed -i -e 's/vendorSha256 = ".*";/vendorSha256 = "";/g' nix/build.nix
new_hash="$(nix build 2>&1 >/dev/null |  grep "got:" | sed -e 's/got://g' | xargs)"
echo "$new_hash"
sed -i -e "s|vendorSha256 = \"\";|vendorSha256 = \"$new_hash\";|g" nix/build.nix
nix build -L --show-trace