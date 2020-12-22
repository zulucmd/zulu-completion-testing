#!sh

#############################################################################
# This file holds all tests that apply to all shells (bash, zsh, fish)
# It cannot use any if statements since their syntax is not portable
# between shells.
#
# For tests that are specific to a shell, use the proper specific file.
#############################################################################

# Basic first level commands (static completion)
_completionTests_verifyCompletion "testprog comp " "completion"
_completionTests_verifyCompletion "testprog help comp" "completion"
_completionTests_verifyCompletion "testprog completion " "bash fish zsh"


# Test logging using $BASH_COMP_DEBUG_FILE

# Test completion from middle of line

# Test --flag= form

# Test ShellCompDirectiveDefault => File completion when no other completions
# Test ShellCompDirectiveNoSpace => No space even when there is a single completion
# Test ShellCompDirectiveNoFileComp => No file completion even when there are no other completions
# Test ShellCompDirectiveFilterFileExt => Not supported for fish, file completion instead
# Test ShellCompDirectiveFilterDirs => Not supported for fish, file completion instead
# Test ShellCompDirectiveError => File completion only

# Test when completions are not filtered by prefix.  NoSpace and FileComp should still work.

