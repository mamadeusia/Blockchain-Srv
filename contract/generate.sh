#!/usr/bin/sh

solc \
    --output-dir . \
    --combined-json abi,bin \
    @openzeppelin=strikeDAO/node_modules/@openzeppelin \
    ./strikeDAO/src/contracts/strikeDao/strikeGovernor.sol \

abigen \
    --combined-json combined.json \
    --pkg contract \
    --out generated.go

rm combined.json