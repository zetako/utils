#!/usr/bin/env sh

set -eu

if [ "$#" -ne 1 ]; then
	echo "usage: $0 <commit-message-file-or-message>" >&2
	exit 2
fi

input=$1

if [ -f "$input" ]; then
	message=$(sed '/^#/d' "$input" | sed -n '/./{p;q;}')
else
	message=$input
fi

if [ -z "$message" ]; then
	echo "commit message is empty" >&2
	exit 1
fi

if ! perl -CSDA -e 'exit(($ARGV[0] =~ /^(?:[\x{2600}-\x{27BF}]|[\x{1F300}-\x{1FAFF}])(?:\x{FE0F})?\s+\S.*$/u) ? 0 : 1)' "$message"; then
	echo "commit message must start with a gitmoji in unicode mode, followed by a space and message text" >&2
	echo "example: ✨ add ContextValue helper" >&2
	exit 1
fi