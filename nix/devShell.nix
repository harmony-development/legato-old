{ pkgs }:
with pkgs;
mkDevShell {
  packages = [ go protobuf vips ];
  env = {
    GOROOT = "${go}/share/go";
    PROTOC = "${protobuf}/bin/protoc";
    PROTOC_INCLUDE = "${protobuf}/include";
    NIX_CONFIG = ''
      substituters = https://cache.nixos.org https://legato.cachix.org
      trusted-public-keys = cache.nixos.org-1:6NCHdD59X431o0gWypbMrAURkbJ16ZPMQFGspcDShjY= legato.cachix.org-1:nz70bLTILsKu/HLxiEbk33jjfYxzeFgSvKzfwejwBTQ=
    '';
  };
}
