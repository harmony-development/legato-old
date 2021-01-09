{ pkgs }:
with pkgs;
buildGoModule rec {
  pname = "legato";
  version = "main";

  src = ../.;

  buildInputs = [ vips ];

  # This needs to be updated every time go dependencies change
  vendorSha256 = "sha256-ocWjxuDqKeDeHQLmNL6g5ecvQ8SYfWjWSlDGRc7nzIM=";

  subPackages = [ "." ];

  meta = with stdenv.lib; {
    description = "Reference server implementation for the Harmony protocol.";
    homepage = "https://github.com/harmony-development/legato";
    license = licenses.agpl3;
    maintainers = [ maintainers.yusdacra ];
    platforms = platforms.linux;
  };
}