{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, nixpkgs, nixpkgs-unstable }: let
    system = "x86_64-darwin";
    pkgs = import nixpkgs { inherit system; };
    unstable-pkgs = import nixpkgs-unstable {
      inherit system;
    };
  in {
    devShells.${system}.default = pkgs.mkShell {
      name = "yog-sothoth";
      packages = [
        unstable-pkgs.pulumictl
        unstable-pkgs.go
        unstable-pkgs.gopls
        unstable-pkgs.gotools
        unstable-pkgs.nodejs_20
      ];
    };
  };
}
