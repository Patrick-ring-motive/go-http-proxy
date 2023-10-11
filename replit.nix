{ pkgs, legacyPolygott }: {
    deps = [
        pkgs.go_1_19
        pkgs.gopls
    ] ++ legacyPolygott;
}