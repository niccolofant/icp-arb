#!/bin/bash
set -e

NETWORK="ic"

usage() {
  echo "Usage:"
  echo "  Interactive: ./approve.sh"
  echo "  Non-interactive: ./approve.sh -t <token_canister_id> -s <spender_principal> -a <amount>"
  echo
  echo "Example:"
  echo "  ./approve.sh -t ryjl3-tyaaa-aaaaa-aaaba-cai -s ig5kc-piaaa-aaaak-qosrq-cai -a 1000000000"
  exit 1
}

while getopts "t:s:a:h" opt; do
  case $opt in
    t) TOKEN_CANISTER="$OPTARG" ;;
    s) SPENDER="$OPTARG" ;;
    a) AMOUNT="$OPTARG" ;;
    h) usage ;;
    *) usage ;;
  esac
done

if [ -z "$TOKEN_CANISTER" ]; then
  read -p "Enter the token canister ID: " TOKEN_CANISTER
fi
if [ -z "$SPENDER" ]; then
  read -p "Enter the spender principal: " SPENDER
fi
if [ -z "$AMOUNT" ]; then
  read -p "Enter the amount (in smallest units): " AMOUNT
fi

echo
echo "Approving $AMOUNT for spender $SPENDER on token $TOKEN_CANISTER ..."
echo

dfx canister call --network $NETWORK $TOKEN_CANISTER icrc2_approve "(
  record {
    fee = null;
    memo = null;
    from_subaccount = null;
    created_at_time = null;
    amount = $AMOUNT : nat;
    expected_allowance = null;
    expires_at = null;
    spender = record {
      owner = principal \"$SPENDER\";
      subaccount = null;
    };
  }
)"
