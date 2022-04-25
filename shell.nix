{ pkgs ? import <nixpkgs> { } }:
with pkgs;
let
  inherit (pkgs) stdenv;
  apple_sdk = darwin.apple_sdk.frameworks;
in mkShell {
  name = "doko";
  PGSQL_HOME = builtins.getEnv "PGSQL_HOME";
  PGSQL_DATADIR = builtins.getEnv "PGSQL_DATADIR";
  buildInputs = [
    buildkit
    docker
    docker-compose_2
    go_1_18
    go-tools
    go-migrate
    nodejs-17_x
    postgresql_14
    libiconv
    git
    git-secret
    gitAndTools.delta
    binutils-unwrapped
    hostname
    inetutils
    openssh
    pkg-config
    rsync
    fasd
    fzf
    zoxide
    gnupg
    jq
    less
    lesspipe
    lsof
    lzma
    jetbrains-mono
    yamllint
  ] ++ lib.optionals stdenv.isDarwin [
    darwin.apple_sdk.libs.utmp
    apple_sdk.ApplicationServices
    apple_sdk.CoreServices
    apple_sdk.OpenGL
    apple_sdk.QTKit
    apple_sdk.Security
    apple_sdk.SystemConfiguration
    xcodebuild
  ] ++ lib.optionals (!stdenv.isDarwin) [
    dbeaver
    ungoogled-chromium
    psmisc
    x11vnc
    xclip
    xvfb-run
  ];
  shellHook = ''
    export NIXPKGS_ALLOW_UNFREE=1
    export GPG_TTY=$(tty)
    export GO_BASE_PATH="$(readlink -e $(type -p go) | sed  -e 's/\/bin\/go//g')"
    export GOPATH="''${PWD}/golibs"
    export NODE_HOME="${pkgs.nodejs-17_x}"

    export PGSQL_BASEDIR=${pkgs.postgresql_14}
    export PGSQL_HOME="''${PGSQL_HOME:-$PWD/pgsql}"
    export PGSQL_DATADIR="''${PGSQL_DATADIR:-$PGSQL_HOME/data}"
    export PGSQL_PID_FILE=$PGSQL_HOME/pgsql.pid
    export PGSQL_TCP_PORT=5439

    export PATH=$GO_BASE_PATH/bin:$GOPATH/bin:$NODE_HOME/bin:$PGSQL_BASEDIR/bin:$PATH
    export GOPATH=$GOPATH:$PWD/goprojects

    echo "###################################################################################################################"
    echo "                                                                                "
    echo "##   !! DOKO NIX DEVELOPMENT ENVIRONMENT ;) !!      "
    echo "##   GO_BASE_PATH: $GO_BASE_PATH                    "
    echo "##   NODE_HOME: $NODE_HOME                          "
    echo "##   PGSQL_BASEDIR: $PGSQL_BASEDIR                  "
    echo "##   PGSQL_HOME: $PGSQL_HOME                        "
    echo "##   PGSQL_DATADIR: $PGSQL_DATADIR                  "
    echo "                                                                                "
    echo "###################################################################################################################"
    mkdir -p $PWD/golibs/{bin,src,pkg}
    mkdir -p $PWD/goprojects/{bin,src,pkg}
    mkdir -p $PGSQL_HOME
    mkdir -p $PGSQL_DATADIR

    export PGSQLD_PID=$(ps -ax | grep -v " grep " | grep $PGSQL_BASEDIR/bin/postgres | awk '{ print $1 }')
    if [[ -z "$PGSQLD_PID" ]]; then
      [ ! "$(ls -A ''${PGSQL_DATADIR})" ] && rm -rf $PGSQL_HOME/logfile && $PGSQL_BASEDIR/bin/pg_ctl -U postgres -D $PGSQL_DATADIR initdb
      $PGSQL_BASEDIR/bin/pg_ctl -U postgres -D $PGSQL_DATADIR -o "-p ''${PGSQL_TCP_PORT}" -l $PGSQL_HOME/logfile start
      export PGSQLD_PID=$!

      createuser -h localhost -p $PGSQL_TCP_PORT -d -r -S -e doko

      createdb -h localhost -p $PGSQL_TCP_PORT -U doko doko_development
      createdb -h localhost -p $PGSQL_TCP_PORT -U doko doko_test
      createdb -h localhost -p $PGSQL_TCP_PORT -U doko doko_e2e_test
    fi

    export GPG_TTY=$(tty)
    export NIX_SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt
    if [[ "$OSTYPE" == "darwin"* ]]; then
       export NIX_SSL_CERT_FILE=/etc/ssl/cert.pem
    fi

    export GPG_TTY='(tty)'

    cleanup()
    {
      echo -e "\nBYE!!! EXITING doko NIX DEVELOPMENT ENVIRONMENT."
      rm -f $PGSQL_HOME/init_doko_db.sql
      export PGSQLD_PID=$(ps -ax | grep -v " grep " | grep $PGSQL_BASEDIR/bin/postgres | awk '{ print $1 }')
      if [[ ! -z "$PGSQLD_PID" ]]; then
        echo -e "PostgreSQL Server still running on Port: $PGSQL_TCP_PORT, at PID: $PGSQLD_PID"
        echo -e "You may choose to SHUTDOWN PostgreSQL Server by issuing 'pg_ctl -D $PGSQL_DATADIR stop && rm -f $PGSQL_HOME/logfile'\n"
      fi
    }
    trap cleanup EXIT
  '';
}
