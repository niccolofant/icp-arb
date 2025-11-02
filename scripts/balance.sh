#!/bin/bash
set -e

NETWORK="ic"
DEFAULT_OWNER="2w4lu-h3ht5-x3wbu-udqvf-wqczf-srnyt-stycn-6sci5-5lz77-x3ows-kqe"

usage() {
  echo "Usage:"
  echo "  Interactive: ./balance.sh"
  echo "  Non-interactive: ./balance.sh -t <token_canister_id> [-o <owner_principal>]"
  echo
  echo "Example:"
  echo "  ./balance.sh -t ryjl3-tyaaa-aaaaa-aaaba-cai -o abcde-fghij-klmno-pqrst-uvwxy-cai"
  exit 1
}

while getopts "t:o:h" opt; do
  case $opt in
    t) TOKEN_CANISTER="$OPTARG" ;;
    o) OWNER_PRINCIPAL="$OPTARG" ;;
    h) usage ;;
    *) usage ;;
  esac
done

if [ -z "$TOKEN_CANISTER" ]; then
  read -p "Enter the token canister ID: " TOKEN_CANISTER
fi
if [ -z "$OWNER_PRINCIPAL" ]; then
  read -p "Enter the owner principal (press Enter for default): " OWNER_PRINCIPAL
  OWNER_PRINCIPAL=${OWNER_PRINCIPAL:-$DEFAULT_OWNER}
fi

echo
echo "Fetching balance for $OWNER_PRINCIPAL on token $TOKEN_CANISTER ..."
echo

dfx canister call --network $NETWORK $TOKEN_CANISTER icrc1_balance_of "(
  record {
    owner = principal \"$OWNER_PRINCIPAL\";
    subaccount = null;
  }
)"
