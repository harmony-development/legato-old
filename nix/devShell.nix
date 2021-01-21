{ pkgs }:
with pkgs;
devshell.mkShell {
  packages = [ go protobuf vips ];
  env = {
    GOROOT = "${go}/share/go";
    PROTOC = "${protobuf}/bin/protoc";
    PROTOC_INCLUDE = "${protobuf}/include";
    NIX_CONFIG = ''
      substituters = https://cache.nixos.org https://harmony.cachix.org
      trusted-public-keys = cache.nixos.org-1:6NCHdD59X431o0gWypbMrAURkbJ16ZPMQFGspcDShjY= harmony.cachix.org-1:yv78QZHgS0UHkrMW56rccNghWHRz18fFRl8mWQ63M6E=
    '';
  };
}
