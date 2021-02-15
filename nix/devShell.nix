{ pkgs }:
with pkgs;
devshell.mkShell {
  packages = [ protobuf vips ];
  commands = [
    { package = go; }
  ];
  env =
    let
      nv = lib.nameValuePair;
    in
    [
      (nv "GOROOT" "${go}/share/go")
      (nv "PROTOC" "${protobuf}/bin/protoc")
      (nv "PROTOC_INCLUDE" "${protobuf}/include")
      (
        nv "NIX_CONFIG" ''
          substituters = https://cache.nixos.org https://harmony.cachix.org
          trusted-public-keys = cache.nixos.org-1:6NCHdD59X431o0gWypbMrAURkbJ16ZPMQFGspcDShjY= harmony.cachix.org-1:yv78QZHgS0UHkrMW56rccNghWHRz18fFRl8mWQ63M6E=
        ''
      )
    ];
}
