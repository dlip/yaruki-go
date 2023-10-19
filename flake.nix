{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    templ-input.url = "github:a-h/templ";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    templ-input,
  }:
    flake-utils.lib.eachDefaultSystem
    (system: let
      pkgs = import nixpkgs {
        inherit system;
        config.allowUnfree = true;
        overlays = [
          (final: prev: {
            templ = templ-input.packages.${final.system}.default;
          })
        ];
      };
    in rec {
      inherit pkgs;
      devShell = pkgs.mkShell {
        buildInputs = with pkgs; [
          go
          air
          templ
          goreman
        ];
      };
    });
}
