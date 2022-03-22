{ pkgs ? import <nixpkgs> { } }:
with pkgs;
mkShell {
  name = "go-shell";
  buildInputs = [
    go_1_18
  ];
}

