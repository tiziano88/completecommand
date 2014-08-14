function _completecommand {
  CMD=${words[1]}

  ARGS=(${words[2,-1]})
  # Remove currently edited word, since it may be an incomplete and therefore invalid flag.
  ARGS[(($CURRENT - 1))]=

  # TODO: Suppress output.
  flags=("${(@f)$($CMD $ARGS -__complete__)}")

  [[ $? == 0 ]] || return

  _arguments "${flags[@]}"
}

compdef _completecommand -P "*"
