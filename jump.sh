function jumper() {
    local JUMP
    JUMP="$GOPATH/bin/jumper"

    # using this as a check for the env var
    $JUMP
    cd "$JUMP_TO"
}
