{ pkgs }:
with pkgs;
devshell.mkShell {
  packages = [ protobuf vips ];
  commands = [
    {
      package = go;
    }
  ];
  env = [
    {
      name = "GOROOT";
      value = "${go}/share/go";
    }
    { name = "PROTOC"; value = "${protobuf}/bin/protoc"; }
    { name = "PROTOC_INCLUDE"; value = "${protobuf}/include"; }
    {
      name = "NIX_CONFIG";
      value = ''
        substituters = https://cache.nixos.org https://harmony.cachix.org
        trusted-public-keys = cache.nixos.org-1:6NCHdD59X431o0gWypbMrAURkbJ16ZPMQFGspcDShjY= harmony.cachix.org-1:yv78QZHgS0UHkrMW56rccNghWHRz18fFRl8mWQ63M6E=
      '';
    }
  ];
}
