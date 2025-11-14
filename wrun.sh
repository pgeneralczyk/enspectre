#!/bin/bash

SRC="main.cr"
BIN="enspectre"

# Sprawdzamy, czy plik binarny istnieje i czy jest starszy niż źródło
if [ ! -f "$BIN" ] || [ "$SRC" -nt "$BIN" ]; then
    echo "Compiling $SRC..."
    crystal build "$SRC" -o "$BIN"
fi

# Uruchamiamy binarkę
./"$BIN"

