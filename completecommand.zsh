function _completecommand {
  CMD=${words[1]}

  ARGS=(${words[2,-1]})
  # Remove currently edited word, since it may be an incomplete and therefore invalid flag.
  CANDIDATE=${ARGS[(($CURRENT - 1))]}
  ARGS[(($CURRENT - 1))]=

  out=$($CMD -__complete__=zsh $ARGS) 2> /dev/null
  [[ $? == 0 ]] || return
  [[ "$out" != "" ]] || return

  flags=("${(@f)out}")

  [[ $? == 0 ]] || return

  _arguments "${flags[@]}"
  _arguments "${flags[@]}"
}

compdef _completecommand -P "*"
