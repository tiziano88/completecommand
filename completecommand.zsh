function _completecommand {
  CMD=${words[1]}

  ARGS=(${words[2,-1]})
  # Remove currently edited word, since it may be an incomplete and therefore invalid flag.
  ARGS[(($CURRENT - 1))]=

  # TODO: Suppress output.
  out=$($CMD $ARGS -__complete__) 2> /dev/null
  [[ $? == 0 ]] || return
  [[ "$out" != "" ]] || return

  flags=("${(@f)out}")

  [[ $? == 0 ]] || return

  _arguments "${flags[@]}"
}

compdef _completecommand -P "*"
