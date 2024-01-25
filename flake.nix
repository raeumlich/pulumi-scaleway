{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, nixpkgs-unstable, flake-utils }:
  let
  in
  flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs { inherit system; };
      unstable-pkgs = import nixpkgs-unstable { inherit system; };
    in {
      devShells.default = pkgs.mkShell
      {
        name = "pulumi-scaleway";
        packages = [
          unstable-pkgs.cosign
          unstable-pkgs.dotnet-sdk_8
          unstable-pkgs.go
          unstable-pkgs.golangci-lint
          unstable-pkgs.gopls
          unstable-pkgs.goreleaser
          unstable-pkgs.gotools
          unstable-pkgs.nodejs_20
          unstable-pkgs.pulumi-bin
          unstable-pkgs.pulumictl
          unstable-pkgs.python3
          unstable-pkgs.python311Packages.setuptools
          unstable-pkgs.syft
          unstable-pkgs.yarn
        ];
      };
    }
  );
}
