{
  description = "Flake for legato, reference server implementation for Harmony protocol";

  inputs = {
    flakeUtils.url = "github:numtide/flake-utils";
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
  };

  outputs = inputs: with inputs; with flakeUtils.lib;
    eachSystem [ "x86_64-linux" "i686-linux" ] (system:
      let
        pkgs = import nixpkgs { inherit system; };

        packages = {
          legato = import ./nix/build.nix { inherit pkgs; };
        };
        apps = builtins.mapAttrs (n: v: mkApp { name = n; drv = v; }) packages;
      in
      {
        inherit packages;
        defaultPackage = packages.legato;

        inherit apps;
        defaultApp = apps.legato;

        devShell = import ./nix/devShell.nix { inherit pkgs; };
      });
}
