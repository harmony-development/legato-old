{ pkgs }:
with pkgs;
mkShell {
  buildInputs = [ protobuf vips ];
  nativeBuildInputs = [ go ];
  shellHook = ''
    export GOROOT="${go}/share/go"
    export PROTOC="${protobuf}/bin/protoc"
    export PROTOC_INCLUDE="${protobuf}/include"
  '';
}
