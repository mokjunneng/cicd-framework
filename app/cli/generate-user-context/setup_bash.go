package main

import (
	"fmt"
	"os"
)

const BASH = `#! /bin/bash

: ${PROG:=$(basename ${BASH_SOURCE})}

_cli_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete $PROG
unset PROG`

func setupBash() error {
	for _, v := range BINARY_NAME {
		err := write("/etc/bash_completion.d/"+v, BASH, false)
		if err != nil {
			return err
		}
	}
	fmt.Println("Please either restart your shell or run:")
	fmt.Println("\tsource ~/.bashrc")
	return nil
}

func tearDownBash() error {
	for _, v := range BINARY_NAME {
		err := os.Remove("/etc/bash_completion.d/" + v)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
