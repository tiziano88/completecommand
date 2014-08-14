function _completecommand {
  CMD=${words[1]}
  flags=("${(@f)$($CMD -__complete__)}")

  _arguments "${flags[@]}"
}

compdef _completecommand -P "*"
