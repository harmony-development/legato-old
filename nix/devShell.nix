{ pkgs }:
with pkgs;
mkShell {
  buildInputs = [ protobuf vips ];
  nativeBuildInputs = [ go ];
  shellHook = ''
    export GOROOT="${go}/share/go"
    export PROTOC="${protobuf}/bin/protoc"
    export PROTOC_INCLUDE="${protobuf}/include"

    export NIX_CONFIG="
      substituters = https://cache.nixos.org https://legato.cachix.org
      trusted-public-keys = cache.nixos.org-1:6NCHdD59X431o0gWypbMrAURkbJ16ZPMQFGspcDShjY= legato.cachix.org-1:nz70bLTILsKu/HLxiEbk33jjfYxzeFgSvKzfwejwBTQ=
    "    
  '';
}
